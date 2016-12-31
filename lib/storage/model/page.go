package model

const RecordsSize = 10

type Page struct {
	pid int
	//records    [RecordsSize]Record
	records   []byte
	numRecord int
}

func NewPage(pid int, data []byte) *Page {
	return &Page{pid, data, 100}
}

func (p *Page) Pid() int {
	return p.pid
}

func (p *Page) Records() []byte {
	return p.records
}

func (p *Page) NumRecord() int {
	return p.numRecord
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
