package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"

	"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/util"
)

// BufferPoolSize is the size of BufferPool.
const BufferPoolSize = 3

// BufferManager manages resources of the main memory and behaves as cache of
// disk storage. If the page data is in buffer, BufferManager treats it on
// memory. If not, BufferManager will fetche the byte data from the disk
// storage. Where pid is the 'page id' and fidx is the 'frame index'.
type BufferManager struct {
	bufferpool [BufferPoolSize]*Frame
	dict       map[int64]int64 // pid -> fidx
	dm         *DiskManager
}

// NewBufferManager creates a BufferManager object with preparing frames.
func NewBufferManager() *BufferManager {

	var bufferpool [BufferPoolSize]*Frame
	for i := 0; i < BufferPoolSize; i++ {
		bufferpool[i] = NewFrame(nil)
	}

	return &BufferManager{bufferpool, make(map[int64]int64, BufferPoolSize),
		NewDiskManager()}
}

// Read reads a page from buffer or disk. The return value is set in given page
// pointer.
func (bm *BufferManager) Read(pid int64, p *page.DataPage) error {

	// The required page is on the cache.
	if fidx, ok := bm.hitPage(pid); ok {
		*p = *bm.bufferpool[fidx].Page()
		return nil
	}

	// Choose the frame which will be set the page.
	fidx, ok := bm.getFreeBuffer()
	if !ok {
		fidx = bm.chooseVictim()
		bm.WriteBack(fidx)
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
	bm.bufferpool[fidx].SetPage(p)
	bm.dict[pid] = fidx

	return nil
}

// Update updates the page which has the page id with the given page.
func (bm *BufferManager) Update(pid int64, p *page.DataPage) error {

	// Choose the frame which will be set the page.
	fidx, ok := bm.hitPage(pid)
	if !ok {
		fidx = bm.chooseVictim()
		bm.WriteBack(fidx)
	}

	// Set the page while considering about the transaction.
	if bm.bufferpool[fidx].PinCount() == 0 {
		bm.bufferpool[fidx].SetPage(p)
		bm.bufferpool[fidx].TurnOnDirty()
		return nil
	} else {
		// TODO: wait
		return errors.New("can't update")
	}

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
		fidx = bm.chooseVictim()
		bm.WriteBack(fidx)
	}

	// Set the new page to the frame.
	p.SetPid(pid)
	bm.bufferpool[fidx].SetPage(p)
	bm.bufferpool[fidx].TurnOnDirty()
	bm.dict[pid] = fidx

	return pid, nil
}

// WriteBack is the expiring process of the page. The dirty page must be written
// it's contents back to the disk. The non dirty page doesn't need to do so.
func (bm *BufferManager) WriteBack(fidx int64) error {

	// Get the target page.
	p := bm.bufferpool[fidx].Page()

	// Clean the buffer.
	bm.bufferpool[fidx].DeletePage()
	delete(bm.dict, p.Pid())

	// Non dirty page doesn't need to be written back.
	if !bm.bufferpool[fidx].Dirty() {
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
	bm.bufferpool[fidx].TurnOffDirty()

	return nil
}

// WriteBackAll executes WriteBack process for all of pages on the cache.
func (bm *BufferManager) WriteBackAll() error {
	for _, fidx := range bm.dict {
		if err := bm.WriteBack(fidx); err != nil {
			return err
		}
	}
	return nil
}

func (bm *BufferManager) pin(fidx int64) {
	bm.bufferpool[fidx].IncPinCount()
}

func (bm *BufferManager) unpin(fidx int64) {
	bm.bufferpool[fidx].DecPinCount()
}

func (bm *BufferManager) hitPage(pid int64) (int64, bool) {
	fidx, ok := bm.dict[pid]
	return fidx, ok
}

func (bm *BufferManager) getFreeBuffer() (int64, bool) {
	for i := 0; i < BufferPoolSize; i++ {
		if bm.bufferpool[i].Page() == nil {
			return int64(i), true
		}
	}
	return -1, false
}

func (bm *BufferManager) chooseVictim() int64 {
	// TODO: use some algorithm (i.g. LRU, FIFO, ...)
	return rand.Int63n(BufferPoolSize)
}

// Dump prints the inner information. It's for debug.
func (bm *BufferManager) Dump() {
	for i := 0; i < len(bm.bufferpool); i++ {
		fmt.Println(bm.bufferpool[i])
	}
	fmt.Println(bm.dict)
}
