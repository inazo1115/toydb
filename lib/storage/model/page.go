package file

/*
** packed implements
 */

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/toydb/record"
)

const RecordsSize = 10

type Page struct {
	pid        int
	records    [RecordsSize]Record
	num_record int
}

type PageBinary struct {
	pid        int
	data       []byte
	num_record int
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

func (p *Page) marshalBinary() PageBinary {
	// TODO: impl
}

func (pb *PageBinary) unmarshalBinary() Page {
	// TODO: impl
}
