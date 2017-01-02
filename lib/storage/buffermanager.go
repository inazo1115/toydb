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

// Read reads a page from buffer or disk.
func (bm *BufferManager) Read(pid int64, page *page.DataPage) (*page.DataPage, error) {

	// Return the page from the cache.
	if fidx, ok := bm.hitPage(pid); ok {
		return bm.bufferpool[fidx].Page(), nil
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
		return nil, err
	}
	buf := make([]byte, size)
	err = bm.dm.Read(pid, buf)
	if err != nil {
		return nil, err
	}

	// Deserialize from byte data to a page struct.
	dec := gob.NewDecoder(bytes.NewBuffer(buf))
	err = dec.Decode(page)
	if err != nil {
		return nil, err
	}

	// Set the page to the buffer.
	bm.bufferpool[fidx].SetPage(page)
	bm.dict[pid] = fidx

	return page, nil
}

// Update updates the page which has the page id with the given page.
func (bm *BufferManager) Update(pid int64, page *page.DataPage) error {

	// Choose the frame which will be set the page.
	fidx, ok := bm.hitPage(pid)
	if !ok {
		fidx = bm.chooseVictim()
		bm.WriteBack(fidx)
	}

	// Set the page while considering about the transaction.
	if bm.bufferpool[fidx].PinCount() == 0 {
		bm.bufferpool[fidx].SetPage(page)
		bm.bufferpool[fidx].TurnOnDirty()
		return nil
	} else {
		// TODO: wait
		return errors.New("can't update")
	}

	return nil
}

// Create appends the page to the disk and returns the new page id.
func (bm *BufferManager) Create(page *page.DataPage) (int64, error) {

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
	page.SetPid(pid)
	bm.bufferpool[fidx].SetPage(page)
	bm.bufferpool[fidx].TurnOnDirty()
	bm.dict[pid] = fidx

	return pid, nil
}

// WriteBack is the expiring process of the page. The dirty page must be written
// it's contents back to the disk. The non dirty page doesn't need to do so.
func (bm *BufferManager) WriteBack(fidx int64) error {

	// Get the target page.
	page := bm.bufferpool[fidx].Page()

	// Clean the buffer.
	bm.bufferpool[fidx].DeletePage()
	delete(bm.dict, page.Pid())

	// Non dirty page doesn't need to be written back.
	if !bm.bufferpool[fidx].Dirty() {
		return nil
	}

	// Dirty page needs to be written back.

	// Do seriarize.
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(page)
	if err != nil {
		return err
	}

	// Write the page back to the disk storage.
	err = bm.dm.Write(page.Pid(), buf.Bytes())
	if err != nil {
		return err
	}

	// Reset the dirty flag.
	bm.bufferpool[fidx].TurnOffDirty()

	return nil
}

// WriteBackAll executes WriteBack process for all of pages on the cache.
func (bm *BufferManager) WriteBackAll() error {
	for _, fidx := range bm.dict {
		err := bm.WriteBack(fidx)
		if err != nil {
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
