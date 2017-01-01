package model

import (
	"bytes"
	"encoding/binary"
	//"fmt"
	"errors"
	"math"

	"github.com/inazo1115/toydb/lib/pkg"
)

// tmp
const recordSize = 24 // 8 * 3
const StructMetaInfoSize = 16
const Int64Size = 8
const FreeSpaceSize = pkg.BlockSize - (StructMetaInfoSize + (Int64Size * 4))

type Page struct {
	pid        int64
	previous   int64
	next       int64
	numRecords int64
	// TODO: add schema etc.
	data []byte
}

func NewPage(pid int64, previous int64, next int64) *Page {
	maxNumRecords := int(math.Floor(float64(FreeSpaceSize / recordSize)))
	size := maxNumRecords * recordSize
	return &Page{pid, previous, next, 0, make([]byte, size)}
}

func (p *Page) Pid() int64 {
	return p.pid
}

func (p *Page) SetPid(pid int64) {
	p.pid = pid
}

func (p *Page) Previous() int64 {
	return p.previous
}

func (p *Page) SetPrevious(page int64) {
	p.previous = page
}

func (p *Page) Next() int64 {
	return p.next
}

func (p *Page) SetNext(page int64) {
	p.next = page
}

func (p *Page) NumRecords() int64 {
	return p.numRecords
}

func (p *Page) AddRecord(r []byte) error {

	if len(r) > recordSize {
		return errors.New("record size is too big")
	}

	for i := 0; i < recordSize; i++ {
		if i >= len(r) {
			p.data[int(p.numRecords)*recordSize+i] = 0
		} else {
			p.data[int(p.numRecords)*recordSize+i] = r[i]
		}
	}

	p.numRecords++

	return nil
}

func (p *Page) Data() []byte {
	return p.data
}

func (p *Page) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.LittleEndian, p.previous)
	_ = binary.Write(buf, binary.LittleEndian, p.next)
	_ = binary.Write(buf, binary.LittleEndian, p.numRecords)
	_ = binary.Write(buf, binary.LittleEndian, p.data)
	return buf.Bytes(), nil
}

func (p *Page) UnmarshalBinary(data []byte) error {

	buf := new(bytes.Buffer)
	buf.Write(data)

	if err := binary.Read(buf, binary.LittleEndian, &p.previous); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.next); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.numRecords); err != nil {
		return err
	}

	tmp := make([]byte, p.numRecords*recordSize)
	if err := binary.Read(buf, binary.LittleEndian, &tmp); err != nil {
		return err
	}
	p.data = tmp

	return nil
}
