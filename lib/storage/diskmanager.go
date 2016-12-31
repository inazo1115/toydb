package storage

import (
	"fmt"
	"math"
	"os"

	"github.com/inazo1115/toydb/lib/pkg"
)

type DiskManager struct {
}

func NewDiskManager() DiskManager {
	return DiskManager{}
}

func (dm *DiskManager) Fetch(pid int, buffer []byte) error {

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

func (dm *DiskManager) Update(pid int, data []byte) error {

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

func (dm *DiskManager) Insert(data []byte) error {

	pid, err := dm.GetMaximumPageID()

	if err != nil {
		return err
	}

	return dm.Update(pid+1, data)
}

func (dm *DiskManager) GetMaximumPageID() (int, error) {

	file, err := dm.getFile()

	if err != nil {
		return 0, err
	}

	info, err := file.Stat()

	if err != nil {
		return 0, err
	}

	pid := int(math.Floor(float64(info.Size() / pkg.BlockSize)))
	return pid, nil
}

func (dm *DiskManager) PrintPage(pid int) {

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

func (dm *DiskManager) getFile() (*os.File, error) {

	file, err := os.OpenFile(pkg.DataFile, os.O_CREATE|os.O_RDWR, 0660)

	if err != nil {
		return nil, err
	}

	return file, nil
}
