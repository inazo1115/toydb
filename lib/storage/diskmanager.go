package storage

import (
	"fmt"
	"math"
	"os"

	"github.com/inazo1115/toydb/lib/pkg"
)

// DiskManager deals random access files provided by OS. DiskManager calls
// file read and file write API and returns results to BufferManager.
type DiskManager struct {
}

// NewDiskManager returns pointer to DiskManager.
func NewDiskManager() *DiskManager {
	return &DiskManager{}
}

// Read reads a byte block from a file and packs it into given buffer.
func (dm *DiskManager) Read(pid int64, buf []byte) error {

	file, err := dm.getFile()
	if err != nil {
		return err
	}

	_, err = file.ReadAt(buf, pid*pkg.BlockSize)
	if err != nil {
		return err
	}

	return nil
}

// Write writes given data to a file.
func (dm *DiskManager) Write(pid int64, data []byte) error {

	file, err := dm.getFile()
	if err != nil {
		return err
	}

	_, err = file.WriteAt(data, pid*pkg.BlockSize)
	if err != nil {
		return err
	}

	return nil
}

// GetFreePageID returns the available free page id.
func (dm *DiskManager) GetFreePageID(used []int64) (int64, error) {

	file, err := dm.getFile()
	if err != nil {
		return -1, err
	}

	stat, err := file.Stat()
	if err != nil {
		return -1, err
	}

	// TODO: refine selection logic.
	min := int64(math.Ceil(float64(stat.Size()) / pkg.BlockSize))

	if len(used) == 0 {
		return min, nil
	}

	max := int64(-1)
	for _, v := range used {
		if max < v {
			max = v
		}
	}

	return int64(math.Max(float64(min), float64(max+1))), nil
}

// GetBufferSize returns a buffer size which is needed to obtain page.
func (dm *DiskManager) GetBufferSize(pid int64) (int32, error) {

	file, err := dm.getFile()
	if err != nil {
		return -1, err
	}

	stat, err := file.Stat()
	if err != nil {
		return -1, err
	}

	if stat.Size() < pkg.BlockSize*(pid+1) {
		return int32(stat.Size() - (pkg.BlockSize * pid)), nil
	}

	return pkg.BlockSize, nil
}

// Dump dumps bytes of specified block. It's for debug.
func (dm *DiskManager) Dump(pid int64) {

	file, err := dm.getFile()
	if err != nil {
		panic(err)
	}

	buf := make([]byte, pkg.BlockSize)
	_, err = file.ReadAt(buf, int64(pid*pkg.BlockSize))
	if err != nil {
		panic(err)
	}

	fmt.Println(buf)
}

// getFile returns pointer to file.
func (dm *DiskManager) getFile() (*os.File, error) {

	file, err := os.OpenFile(pkg.DataFile, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		return nil, err
	}

	/*defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()*/

	return file, nil
}
