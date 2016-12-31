package storage

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/pkg"
	"github.com/inazo1115/toydb/lib/storage/model"
)

const BufferPoolSize = 3

type BufferManager struct {
	bufferpool [BufferPoolSize]model.Frame
	dict       map[int]int // pid -> frame_idx
	dm         DiskManager
}

func NewBufferManager() *BufferManager {

	var bufferpool [BufferPoolSize]model.Frame
	for i := 0; i < BufferPoolSize; i++ {
		bufferpool[i] = model.NewFrame(nil)
	}

	dict := make(map[int]int, BufferPoolSize)

	dm := NewDiskManager()

	return &BufferManager{bufferpool, dict, dm}
}

func (bm *BufferManager) Fetch(pid int) (*model.Page, error) {

	if frame_idx, ok := bm.hitPage(pid); ok {

		fmt.Println("hit")

		bm.pin(frame_idx)
		defer bm.unpin(frame_idx)

		return bm.bufferpool[frame_idx].Page(), nil

	} else {

		buffer := make([]byte, pkg.BlockSize)

		if frame_idx, ok := bm.getFreeBuffer(); ok {

			fmt.Println("not hit free")

			err := bm.dm.Fetch(pid, buffer)
			if err != nil {
				return nil, err
			}

			page := model.NewPage(pid, buffer)
			bm.bufferpool[frame_idx] = model.NewFrame(page)
			bm.dict[pid] = frame_idx

			bm.pin(frame_idx)
			defer bm.unpin(frame_idx)

			return page, nil

		} else {

			fmt.Println("not hit not free")

			frame_idx, _ := bm.chooseVictim()
			bm.flushPage(frame_idx)

			err := bm.dm.Fetch(pid, buffer)
			if err != nil {
				return nil, err
			}

			page := model.NewPage(pid, buffer)
			bm.bufferpool[frame_idx] = model.NewFrame(page)
			bm.dict[pid] = frame_idx

			bm.pin(frame_idx)
			defer bm.unpin(frame_idx)

			return page, nil
		}
	}
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
		_, err := bm.Fetch(pid)
		if err != nil {
			return nil
		}
		return bm.Update(pid, page)
	}
}

func (bm *BufferManager) Insert(page *model.Page) error {

	pid, err := bm.dm.GetMaximumPageID()
	if err != nil {
		return err
	}

	return bm.Update(pid, page)
}

func (bm *BufferManager) pin(frame_idx int) {
	bm.bufferpool[frame_idx].IncPinCount()
}

func (bm *BufferManager) unpin(frame_idx int) {
	bm.bufferpool[frame_idx].DecPinCount()
}

func (bm BufferManager) hitPage(pid int) (int, bool) {
	frame_idx, ok := bm.dict[pid]
	return frame_idx, ok
}

func (bm BufferManager) getFreeBuffer() (int, bool) {
	for i := 0; i < BufferPoolSize; i++ {
		if frame_idx, ok := bm.dict[i]; !ok {
			return frame_idx, true
		}
	}
	return -1, false
}

func (bm BufferManager) chooseVictim() (int, bool) {
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
	//data := page.marshalBinary()
	//dm.WriteBackPage(page.pid, data)
	//frame = nil
	delete(bm.dict, page.Pid())
}

func (bm *BufferManager) Dump() {
	fmt.Printf("bufferpool: %s\n", bm.bufferpool)
	fmt.Printf("dict: %s\n", bm.dict)
}
