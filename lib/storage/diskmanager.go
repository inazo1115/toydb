package storage

import (
	"errors"
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
func (dm *DiskManager) Read(pid int, buffer []byte) error {

	file, err := dm.getFile()
	if err != nil {
		return err
	}

	_, err = file.ReadAt(buffer, int64(pid*pkg.BlockSize))
	if err != nil {
		return err
	}

	return nil
}

// Write writes given data to a file.
func (dm *DiskManager) Write(pid int, data []byte) error {

	file, err := dm.getFile()
	if err != nil {
		return err
	}

	_, err = file.WriteAt(data, int64(pid*pkg.BlockSize))
	if err != nil {
		return err
	}

	return nil
}

// GetFreePageID returns the available free page id.
func (dm *DiskManager) GetFreePageID(used []int) (int, error) {

	file, err := dm.getFile()
	if err != nil {
		return -1, err
	}

	info, err := file.Stat()
	if err != nil {
		return -1, err
	}

	// TODO: refine selection logic.
	min := int(math.Ceil(float64(info.Size()) / pkg.BlockSize))

	if len(used) == 0 {
		return min, nil
	}

	fmt.Println("min")
	fmt.Println(min)

	for pid := min; pid < (math.MaxInt64 / pkg.BlockSize) ; pid++ {
		for _, v := range used {
			if pid != v {
				return pid, nil
			}
		}
	}

	return -1, errors.New("there is no free page id")
}

// Dump dumps bytes of specified block. It's for debug.
func (dm *DiskManager) Dump(pid int) {

	file, err := dm.getFile()
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, pkg.BlockSize)
	_, err = file.ReadAt(buffer, int64(pid*pkg.BlockSize))
	if err != nil {
		panic(err)
	}

	fmt.Println(buffer)
}

// getFile returns pointer to file.
func (dm *DiskManager) getFile() (*os.File, error) {

	file, err := os.OpenFile(pkg.DataFile, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		return nil, err
	}

	return file, nil
}
