package model

import (
	"bytes"
	"fmt"
	"encoding/binary"
)

const RecordsSize = 10

type Page struct {
	pid int64
	//	records   [RecordsSize]*Record
	records   int64
	numRecord int64
}

/*func NewPage(pid int, data []byte) *Page {
	return &Page{pid, data, 0}
}*/
func NewPage(pid int64) *Page {
	return &Page{pid, 33, 44}
}

func (p *Page) Pid() int64 {
	return p.pid
}

/*func (p *Page) Records() [RecordsSize]*Record {
	return p.records
}*/

func (p *Page) NumRecord() int64 {
	return p.numRecord
}

func (p *Page) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.LittleEndian, p.pid)
	_ = binary.Write(buf, binary.LittleEndian, p.records)
	_ = binary.Write(buf, binary.LittleEndian, p.numRecord)
	return buf.Bytes(), nil
}

func (p *Page) UnmarshalBinary(data []byte) error {

	fmt.Println("****")

	buf := new(bytes.Buffer)
	buf.Write(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.pid); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.records); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.numRecord); err != nil {
		return err
	}
	return nil
}

/*
func NewPage(pid int, records []Record) {
	page := Page{pid, [RecordsSize]Record, 0}
	for i := 0; i < len(records); i++ {
		page.records[i] = records[i]
		page.num_record += 1
	}
}
*/

/*func (p *Page) marshalBinary() PageBinary {
	// TODO: impl
}

func (p *Page) unmarshalBinary() Page {
	// TODO: impl
}*/
