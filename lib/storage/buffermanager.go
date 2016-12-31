package storage

import (
	"errors"
	"fmt"

	"github.com/inazo1115/toydb/lib/pkg"
	"github.com/inazo1115/toydb/lib/storage/model"
)

// BufferPoolSize is the size of BufferPool.
const BufferPoolSize = 3

// BufferManager manages resources of the main memory and behaves as the cache
// of disk storage. If the page data is in buffer, BufferManager treats it
// on memory. If not, BufferManager will fetche the byte data from the disk
// storage.
type BufferManager struct {
	bufferpool [BufferPoolSize]*model.Frame
	dict       map[int]int // pid -> frame_idx
	dm         *DiskManager
}

func NewBufferManager() *BufferManager {

	var bufferpool [BufferPoolSize]*model.Frame
	for i := 0; i < BufferPoolSize; i++ {
		bufferpool[i] = model.NewFrame(nil)
	}

	return &BufferManager{bufferpool, make(map[int]int, BufferPoolSize),
		NewDiskManager()}
}

func (bm *BufferManager) Read(pid int) (*model.Page, error) {

	// cache hit
	if frame_idx, ok := bm.hitPage(pid); ok {
		log("hit")

		bm.pin(frame_idx)
		defer bm.unpin(frame_idx)

		return bm.bufferpool[frame_idx].Page(), nil

	}

	log("not hit")

	frame_idx, ok := bm.getFreeBuffer()
	if !ok {
		frame_idx, _ = bm.chooseVictim()
		bm.flushPage(frame_idx)
	}

	buffer := make([]byte, pkg.BlockSize)
	err := bm.dm.Read(pid, buffer)
	if err != nil {
		return nil, err
	}

	page := model.NewPage(pid, buffer) // deserialize
	bm.setPage(frame_idx, pid, page)

	bm.pin(frame_idx)
	defer bm.unpin(frame_idx)

	return page, nil
}

func (bm *BufferManager) Update(pid int, page *model.Page) error {

	if frame_idx, ok := bm.hitPage(pid); ok {
		if bm.bufferpool[frame_idx].PinCount() == 0 {
			bm.bufferpool[frame_idx].SetPage(page)
			return nil
		} else {
			// TODO: wait
			return errors.New("can't update")
		}
	} else {
		_, err := bm.Read(pid)
		if err != nil {
			return err
		}
		return bm.Update(pid, page)
	}
}

func (bm *BufferManager) Create(page *model.Page) (int, error) {

	pid, err := bm.dm.GetFreePageID()
	if err != nil {
		return -1, err
	}

	frame_idx, ok := bm.getFreeBuffer()
	if !ok {
		frame_idx, _ := bm.chooseVictim()
		bm.flushPage(frame_idx)
	}

	bm.setPage(frame_idx, pid, page)

	return pid, nil
}

func (bm *BufferManager) WriteBackAll() error {
	for _, frame_idx := range bm.dict {
		bm.flushPage(frame_idx)
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
		if frame_idx, ok := bm.dict[i]; !ok {
			return frame_idx, true
		}
	}
	return -1, false
}

func (bm *BufferManager) setPage(frame_idx int, pid int, page *model.Page) {
	bm.bufferpool[frame_idx].SetPage(page)
	bm.dict[pid] = frame_idx
}

func (bm *BufferManager) chooseVictim() (int, bool) {
	// TODO: use some algorithm (i.g. LRU, FIFO, ...)
	for i := 0; i < BufferPoolSize; i++ {
		if bm.bufferpool[i].PinCount() == 0 {
			return i, true
		}
	}
	return -1, false
}

func (bm *BufferManager) flushPage(frame_idx int) {
	frame := bm.bufferpool[frame_idx]
	page := frame.Page()
	bm.dm.Write(page.Pid(), page.Records())
	delete(bm.dict, page.Pid())
}

func (bm *BufferManager) Dump() {
	fmt.Printf("bufferpool: %s\n", bm.bufferpool)
	fmt.Printf("dict: %s\n", bm.dict)
}

func log(msg interface{}) {
	fmt.Println(msg)
}
