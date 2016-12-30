package buffer

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/toydb/disk"
	"github.com/inazo1115/toydb/lib/toydb/file"
)

const BufferPoolSize = 10

type BufferManager struct {
	bufferpool [BufferPoolSize]Frame
	dict       map[int]int // pid -> frame_idx
	dm         DiskManager
}

func NewBufferManager() BufferManager {
	dm := NewDiskManager()
	dm.Init()
	return BufferManager{[BufferPoolSize]Frame{}, map[int]int{}, dm}
}

func (bm *BufferManager) ReadPage(pid int) Page {
	if bm.hitPage(pid) {
		frame_idx := bm.dict[pid]
		bm.pin(frame_idx)
		return bm.bufferpool[frame_idx].GetPage()
	} else {
		if frame_idx, ok := bm.getFreeBuffer(); ok {
			page := bm.dm.fetchPage(pid).unmarshalBinary()
			bm.bufferpool[frame_idx] = NewFrame(page)
			return page
		} else {
			frame_idx := bm.chooseVictim()
			bm.flushPage(frame_idx)
			page := bm.dm.fetchPage(pid).unmarshalBinary()
			bm.bufferpool[frame_idx] = NewFrame(page)
			return page
		}
	}
}

func (bm *BufferManager) ReleasePage(pid int) {
	frame_idx := bm.dict[pid]
	bm.unpin(frame_idx)
}

func (bm *BufferManager) pin(frame_idx int) {
	bm.bufferpool[frame_idx].IncPinCount()
}

func (bm *BufferManager) unpin(frame_idx int) {
	bm.bufferpool[frame_idx].DecPinCount()
}

func (bm BufferManager) hitPage(pid int) bool {
	_, ok := bm.dict[pid]
	return ok
}

func (bm BufferManager) getFreeBuffer() (int, bool) {
	for i := 0; i < BufferPoolSize; i++ {
		if n, ok := bm.dict[pid]; !ok {
			return n, true
		}
	}
	return nil, false
}

func (bm BufferManager) chooseVictim() (int, bool) {
	// TODO: use some algorithm (i.g. LRU, FIFO, ...)
	for i := 0; i < BufferPoolSize; i++ {
		if bm.bufferpool[i].pin_count == 0 {
			return i, true
		}
	}
	return nil, false
}

func (bm *BufferManager) flushPage(frame_idx int) int {
	frame := bm.bufferpool[frame_idx]
	page := frame.getPage()
	data := page.marshalBinary()
	dm.WriteBackPage(page.pid, data)
	frame = nil
	delete(bm.dict, pid)
}
