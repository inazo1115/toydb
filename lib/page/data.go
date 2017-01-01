package page

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"

	"github.com/inazo1115/toydb/lib/pkg"
)

// tmp
const RecordSize = 24 // 8 * 3
const StructMetaInfoSize = 16
const Int64Size = 8
const FreeSpaceSize = pkg.BlockSize - (StructMetaInfoSize + (Int64Size * 4))

type DataPage struct {
	pid        int64
	previous   int64
	next       int64
	numRecords int64
	data       []byte
}

func NewDataPage(pid int64, previous int64, next int64) *DataPage {
	return &DataPage{pid, previous, next, 0, make([]byte, getDataSize())}
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

func (p *DataPage) AddRecord(r []byte) error {

	if len(r) > RecordSize {
		return errors.New("record size is too big")
	}

	for i := 0; i < RecordSize; i++ {
		if i >= len(r) {
			p.data[int(p.numRecords)*RecordSize+i] = 0
		} else {
			p.data[int(p.numRecords)*RecordSize+i] = r[i]
		}
	}

	p.numRecords++

	return nil
}

func (p *DataPage) Data() []byte {
	return p.data
}

func (p *DataPage) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.LittleEndian, p.pid)
	_ = binary.Write(buf, binary.LittleEndian, p.previous)
	_ = binary.Write(buf, binary.LittleEndian, p.next)
	_ = binary.Write(buf, binary.LittleEndian, p.numRecords)
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

	tmp := make([]byte, getDataSize())
	if err := binary.Read(buf, binary.LittleEndian, &tmp); err != nil {
		return err
	}
	p.data = tmp

	return nil
}

func getDataSize() int {
	return int(math.Floor(float64(FreeSpaceSize/RecordSize))) * RecordSize
}
