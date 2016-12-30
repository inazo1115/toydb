package block

import (
	"fmt"
	"os"
)

const DBFileName = "toydb.strage"

const BlockSize = 4096

type DiskManager struct {}

func NewDiskManager() {
	return DiskManager{}
}

func (dm *DiskManager) FetchPage(pid int) ([]byte, error) {

	file, err := dm.getFile()
	if err != nil {
		return []byte(), err
	}

	buffer := make([]byte, BlockSize)
	file.ReadAt(buffer, pid * BlockSize)

	return buffer, nil
}

func (dm *DiskManager) WriteBackPage(pid int, data []byte) error {

	file, err := dm.getFile()
	if err != nil {
		return err
	}

	file.WriteAt(data, pid * BlockSize)

	return nil
}

func (dm *DiskManager) getFile() (*os.File, error) {

	file, err := os.OpenFile(FileName, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		return nil, err
	}

	return file, nil
}
