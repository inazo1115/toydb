package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"

	"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/util/datautil"
)

// BufferPoolSize is the size of BufferPool.
const BufferPoolSize = 1

// BufferManager manages resources of the main memory and behaves as the cache
// of disk storage. If the page data is in buffer, BufferManager treats it
// on memory. If not, BufferManager will fetche the byte data from the disk
// storage.
type BufferManager struct {
	bufferpool [BufferPoolSize]*Frame
	dict       map[int]int // pid -> frame_idx
	dm         *DiskManager
}

// NewBufferManager creates a BufferManager object with preparing frames.
func NewBufferManager() *BufferManager {

	var bufferpool [BufferPoolSize]*Frame
	for i := 0; i < BufferPoolSize; i++ {
		bufferpool[i] = NewFrame(nil)
	}

	return &BufferManager{bufferpool, make(map[int]int, BufferPoolSize),
		NewDiskManager()}
}

// Read reads a page from buffer or disk.
func (bm *BufferManager) Read(pid int, page *page.DataPage) error {

	// Return the page from the cache.
	if frame_idx, ok := bm.hitPage(pid); ok {
		page = bm.bufferpool[frame_idx].Page()
		return nil
	}

	// Choose the frame which will be set the page.
	frame_idx, ok := bm.getFreeBuffer()
	if !ok {
		frame_idx, _ = bm.chooseVictim()
		bm.WriteBack(frame_idx)
	}

	// Fetch the byte data from the disk storage.
	size, err := bm.dm.GetBufferSize(int64(pid))
	if err != nil {
		return err
	}
	buf := make([]byte, size)
	err = bm.dm.Read(pid, buf)
	if err != nil {
		return err
	}

	// Deserialize from byte data to a page struct.
	dec := gob.NewDecoder(bytes.NewBuffer(buf))
	err = dec.Decode(page)
	if err != nil {
		return err
	}

	// Set the page to the buffer.
	bm.setPage(frame_idx, pid, page)

	return nil
}

func (bm *BufferManager) Update(pid int, page *page.DataPage) error {

	frame_idx, ok := bm.hitPage(pid)
	if !ok {
		frame_idx, _ = bm.chooseVictim()
		bm.WriteBack(frame_idx)
	}

	if bm.bufferpool[frame_idx].PinCount() == 0 {
		bm.bufferpool[frame_idx].SetPage(page)
		return nil
	} else {
		// TODO: wait
		return errors.New("can't update")
	}

	return nil
}

func (bm *BufferManager) Create(page *page.DataPage) (int, error) {

	// Get an available free page id.
	pid, err := bm.dm.GetFreePageID(datautil.Keys(bm.dict))
	if err != nil {
		return -1, err
	}

	// Choose the frame which will be set the page.
	frame_idx, ok := bm.getFreeBuffer()
	if !ok {
		frame_idx, _ = bm.chooseVictim()
		bm.WriteBack(frame_idx)
	}

	// Set the new page to the frame.
	page.SetPid(int64(pid))
	bm.setPage(frame_idx, pid, page)

	return pid, nil
}

func (bm *BufferManager) WriteBack(frame_idx int) error {

	// Get and delete the target page.
	page := bm.bufferpool[frame_idx].Page()

	// Do seriarize.
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(page)
	if err != nil {
		return err
	}

	// Write back the page to disk storage.
	bm.dm.Write(int(page.Pid()), buf.Bytes())

	// Clean the buffer.
	bm.bufferpool[frame_idx].DeletePage()
	delete(bm.dict, int(page.Pid()))

	return nil
}

func (bm *BufferManager) WriteBackAll() error {
	for _, frame_idx := range bm.dict {
		err := bm.WriteBack(frame_idx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bm *BufferManager) pin(frame_idx int) {
	bm.bufferpool[frame_idx].IncPinCount()
}

func (bm *BufferManager) unpin(frame_idx int) {
	bm.bufferpool[frame_idx].DecPinCount()
}

func (bm *BufferManager) hitPage(pid int) (int, bool) {
	frame_idx, ok := bm.dict[pid]
	return frame_idx, ok
}

func (bm *BufferManager) getFreeBuffer() (int, bool) {
	for i := 0; i < BufferPoolSize; i++ {
		if bm.bufferpool[i].Page() == nil {
			return i, true
		}
	}
	return -1, false
}

func (bm *BufferManager) setPage(frame_idx int, pid int, page *page.DataPage) {
	bm.bufferpool[frame_idx].SetPage(page)
	bm.dict[pid] = frame_idx
}

func (bm *BufferManager) chooseVictim() (int, bool) {
	// TODO: use some algorithm (i.g. LRU, FIFO, ...)
	return rand.Intn(BufferPoolSize), true
}

func (bm *BufferManager) Dump() {
	fmt.Println(bm.bufferpool[0])
	fmt.Println(bm.dict)
}
