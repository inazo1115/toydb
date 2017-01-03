package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/pkg"
	"github.com/inazo1115/toydb/lib/util"
)

// bufferPoolSize is the size of BufferPool.
var bufferPoolSize_ = int64(pkg.BufferPoolSize)

// BufferManager manages resources of the main memory and behaves as cache of
// disk storage. If the page data is in buffer, BufferManager treats it on
// memory. If not, BufferManager will fetche the byte data from the disk
// storage. Where pid is the 'page id' and fidx is the 'frame index'.
type BufferManager struct {
	bufferPool     []*Frame
	bufferPoolSize int64
	dict           map[int64]int64 // pid -> fidx
	dm             *DiskManager
	cacheStrat     *LRUStrategy
}

// NewBufferManager creates a BufferManager object with preparing frames.
func NewBufferManager() *BufferManager {

	bufferPool := make([]*Frame, bufferPoolSize_)
	for i := 0; i < int(bufferPoolSize_); i++ {
		bufferPool[i] = NewFrame(nil)
	}

	return &BufferManager{bufferPool, bufferPoolSize_,
		make(map[int64]int64, bufferPoolSize_), NewDiskManager(),
		NewLRUStrategy()}
}

// Read reads a page from buffer or disk. The return value is set in given page
// pointer.
func (bm *BufferManager) Read(pid int64, p *page.DataPage) error {

	// The required page is on the cache.
	if fidx, ok := bm.hitPage(pid); ok {
		*p = *bm.bufferPool[fidx].Page()
		return nil
	}

	// Choose the frame which will be set the page.
	fidx, ok := bm.getFreeBuffer()
	if !ok {
		fidx = bm.cacheStrat.ChooseVictim(bm)
		bm.writeBack(fidx)
	}

	// Fetch the byte data from the disk storage.
	size, err := bm.dm.GetBufferSize(pid)
	if err != nil {
		return err
	}
	buf := make([]byte, size)
	if err = bm.dm.Read(pid, buf); err != nil {
		return err
	}

	// Deserialize from byte data to a page struct.
	dec := gob.NewDecoder(bytes.NewBuffer(buf))
	if err = dec.Decode(p); err != nil {
		return err
	}

	// Set the page to the buffer.
	bm.bufferPool[fidx].SetPage(p)
	bm.dict[pid] = fidx

	// Tell a page access to the cache strategy.
	bm.cacheStrat.TouchPage(bm, pid)

	return nil
}

// Update updates the page which has the page id with the given page.
func (bm *BufferManager) Update(pid int64, p *page.DataPage) error {

	// Choose the frame which will be set the page.
	fidx, ok := bm.hitPage(pid)
	if !ok {
		fidx = bm.cacheStrat.ChooseVictim(bm)
		bm.writeBack(fidx)
	}

	// Set the page while considering about the transaction.
	if bm.bufferPool[fidx].PinCount() == 0 {
		bm.bufferPool[fidx].SetPage(p)
		bm.bufferPool[fidx].TurnOnDirty()
		return nil
	} else {
		// TODO: wait
		return errors.New("can't update")
	}

	// Tell a page access to the cache strategy.
	bm.cacheStrat.TouchPage(bm, pid)

	return nil
}

// Create appends the page to the disk and returns the new page id.
func (bm *BufferManager) Create(p *page.DataPage) (int64, error) {

	// Get an available free page id.
	pid, err := bm.dm.GetFreePageID(util.Keys(bm.dict))
	if err != nil {
		return -1, err
	}

	// Choose the frame which will be set the page.
	fidx, ok := bm.getFreeBuffer()
	if !ok {
		fidx = bm.cacheStrat.ChooseVictim(bm)
		bm.writeBack(fidx)
	}

	// Set the new page to the frame.
	p.SetPid(pid)
	bm.bufferPool[fidx].SetPage(p)
	bm.bufferPool[fidx].TurnOnDirty()
	bm.dict[pid] = fidx

	// Tell a page access to the cache strategy.
	bm.cacheStrat.TouchPage(bm, pid)

	return pid, nil
}

// WriteBackAll executes WriteBack process for all of pages on the cache.
func (bm *BufferManager) WriteBackAll() error {
	for _, fidx := range bm.dict {
		if err := bm.writeBack(fidx); err != nil {
			return err
		}
	}
	return nil
}

// hitPage checks given page id is on the cache.
func (bm *BufferManager) hitPage(pid int64) (int64, bool) {
	fidx, ok := bm.dict[pid]
	return fidx, ok
}

// getFreeBuffer looks for an empty frame.
func (bm *BufferManager) getFreeBuffer() (int64, bool) {
	for i := 0; i < int(bm.bufferPoolSize); i++ {
		if bm.bufferPool[i].Page() == nil {
			return int64(i), true
		}
	}
	return -1, false
}

// writeBack is the expiring process of the page. The dirty page must be written
// it's contents back to the disk. The non dirty page doesn't need to do so.
func (bm *BufferManager) writeBack(fidx int64) error {

	// Get the target page.
	p := bm.bufferPool[fidx].Page()

	// Clean the buffer.
	bm.bufferPool[fidx].DeletePage()
	delete(bm.dict, p.Pid())

	// Reset the cache hit count.
	bm.bufferPool[fidx].SetHitCount(0)

	// Non dirty page doesn't need to be written back.
	if !bm.bufferPool[fidx].Dirty() {
		return nil
	}

	// Dirty page needs to be written back.

	// Do seriarize.
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		return err
	}

	// Write the page back to the disk storage.
	if err := bm.dm.Write(p.Pid(), buf.Bytes()); err != nil {
		return err
	}

	// Reset the dirty flag.
	bm.bufferPool[fidx].TurnOffDirty()

	return nil
}

// Dump prints the inner information. It's for debug.
func (bm *BufferManager) Dump() {
	for i := 0; i < len(bm.bufferPool); i++ {
		fmt.Println(bm.bufferPool[i])
	}
	fmt.Println(bm.dict)
}
