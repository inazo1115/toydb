package page

import (
	"bytes"
	"encoding/binary"
	"errors"
	//"fmt"

	"github.com/inazo1115/toydb/lib/pkg"
)

// tmp
const StructMetaInfoSize = 20

//const Int64Size = table.INT64.Size()
const Int64Size = 8
const FieldValueSize = Int64Size * 5
const FreeSpaceSize = pkg.BlockSize - (StructMetaInfoSize + FieldValueSize)

type DataPage struct {
	pid        int64
	previous   int64
	next       int64
	numRecords int64
	recordSize int64
	data       []byte
}

func NewDataPage(pid int64, previous int64, next int64, recordSize int64) *DataPage {
	data := make([]byte, FreeSpaceSize)
	return &DataPage{pid, previous, next, 0, recordSize, data}
}

func (p *DataPage) Pid() int64 {
	return p.pid
}

func (p *DataPage) SetPid(pid int64) {
	p.pid = pid
}

func (p *DataPage) Previous() int64 {
	return p.previous
}

func (p *DataPage) SetPrevious(page int64) {
	p.previous = page
}

func (p *DataPage) Next() int64 {
	return p.next
}

func (p *DataPage) SetNext(page int64) {
	p.next = page
}

func (p *DataPage) NumRecords() int64 {
	return p.numRecords
}

func (p *DataPage) HasFreeSpace() bool {
	return (FreeSpaceSize - p.numRecords*p.recordSize) >= p.recordSize
}

func (p *DataPage) Data() []byte {
	return p.data
}

func (p *DataPage) AddRecord(r []byte) error {

	if len(r) != int(p.recordSize) {
		return errors.New("record size is wrong")
	}

	for i := 0; i < int(p.recordSize); i++ {
		p.data[p.numRecords*p.recordSize+int64(i)] = r[i]
	}

	p.numRecords++

	return nil
}

func (p *DataPage) ReadRecord(ridx int64) []byte {
	off := ridx * p.recordSize
	return p.data[off:(off + p.recordSize)]
}

func (p *DataPage) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.LittleEndian, p.pid)
	_ = binary.Write(buf, binary.LittleEndian, p.previous)
	_ = binary.Write(buf, binary.LittleEndian, p.next)
	_ = binary.Write(buf, binary.LittleEndian, p.numRecords)
	_ = binary.Write(buf, binary.LittleEndian, p.recordSize)
	_ = binary.Write(buf, binary.LittleEndian, p.data)
	return buf.Bytes(), nil
}

func (p *DataPage) UnmarshalBinary(data []byte) error {
	buf := new(bytes.Buffer)
	buf.Write(data)

	if err := binary.Read(buf, binary.LittleEndian, &p.pid); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.previous); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.next); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.numRecords); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.recordSize); err != nil {
		return err
	}

	tmp := make([]byte, FreeSpaceSize)
	if err := binary.Read(buf, binary.LittleEndian, &tmp); err != nil {
		return err
	}
	p.data = tmp

	return nil
}
