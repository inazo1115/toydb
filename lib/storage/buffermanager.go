package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	//"math/rand"

	"github.com/inazo1115/toydb/lib/pkg"
	"github.com/inazo1115/toydb/lib/storage/model"
	"github.com/inazo1115/toydb/lib/util/datautil"
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

	// Return the page from the cache.
	if frame_idx, ok := bm.hitPage(pid); ok {
		return bm.bufferpool[frame_idx].Page(), nil
	}

	// Choose the frame which will be set the page.
	frame_idx, ok := bm.getFreeBuffer()
	if !ok {
		frame_idx, _ = bm.chooseVictim()
		bm.flushPage(frame_idx)
	}

	// Fetch the byte data from the disk storage.
	buf := make([]byte, pkg.BlockSize)
	err := bm.dm.Read(pid, buf)
	if err != nil {
		return nil, err
	}

	// Deserialize from byte data to a page struct.
	dec := gob.NewDecoder(bytes.NewBuffer(buf))
	var page model.Page
	err = dec.Decode(&page)
	if err != nil {
		return nil, err
	}
	bm.setPage(frame_idx, pid, &page)

	return &page, nil
}

func (bm *BufferManager) Update(pid int, page *model.Page) error {

	// When the page isn't on the cache.
	if _, ok := bm.hitPage(pid); !ok {
		_, err := bm.Read(pid)
		if err != nil {
			return err
		}
	}

	frame_idx := bm.dict[pid]

	if bm.bufferpool[frame_idx].PinCount() == 0 {
		bm.bufferpool[frame_idx].SetPage(page)
		return nil
	} else {
		// TODO: wait
		return errors.New("can't update")
	}

	return nil
}

func (bm *BufferManager) Create(page *model.Page) (int, error) {

	fmt.Println("++~~")
	fmt.Println(datautil.Keys(bm.dict))
	bm.WriteBackAll()

	pid, err := bm.dm.GetFreePageID(datautil.Keys(bm.dict))
	if err != nil {
		return -1, err
	}

	var frame_idx int

	if frame_idx, ok := bm.getFreeBuffer(); !ok {
		frame_idx, _ = bm.chooseVictim()
		bm.flushPage(frame_idx)
	}

	page.SetPid(int64(pid))
	bm.setPage(frame_idx, pid, page)

	return pid, nil
}

func (bm *BufferManager) WriteBackAll() error {
	for _, frame_idx := range bm.dict {
		err := bm.flushPage(frame_idx)
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

func (bm *BufferManager) setPage(frame_idx int, pid int, page *model.Page) {
	bm.bufferpool[frame_idx].SetPage(page)
	bm.dict[pid] = frame_idx
}

func (bm *BufferManager) chooseVictim() (int, bool) {
	// TODO: use some algorithm (i.g. LRU, FIFO, ...)
	/*for i := 0; i < BufferPoolSize; i++ {
		if bm.bufferpool[i].PinCount() == 0 {
			return i, true
		}
	}
	return -1, false*/
	return 0, true
	//return rand.Intn(BufferPoolSize), true
}

func (bm *BufferManager) flushPage(frame_idx int) error {

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

func (bm *BufferManager) Dump() {
	fmt.Printf("bufferpool: %s\n", bm.bufferpool)
	fmt.Printf("dict: %s\n", bm.dict)
}

func log(msg interface{}) {
	fmt.Println(msg)
}
