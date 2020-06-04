package file

import (
	"errors"
	//"fmt"

	"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/table"
)

// HeapFile is the representation of the file and access methods. HeapFile is
// the linked list of pages.
type HeapFile struct {
	rootPid int64
	bm      *storage.BufferManager
	schema  *table.Schema
}

// NewHeapFile creates HeapFile struct and returns it's pointer.
func NewHeapFile(bm *storage.BufferManager, schema *table.Schema) *HeapFile {

	rootPage := page.NewDataPage(-1, -1, -1, schema.RecordSize())
	rootPid, err := bm.Create(rootPage)
	if err != nil {
		panic(err)
	}

	return &HeapFile{rootPid, bm, schema}
}

// RootPid is a getter of the root page id.
func (f *HeapFile) RootPid() int64 {
	return f.rootPid
}

// RootPid is a getter of the buffer manager.
func (f *HeapFile) BufferManager() *storage.BufferManager {
	return f.bm
}

// Schema is a getter of the schema.
func (f *HeapFile) Schema() *table.Schema {
	return f.schema
}

// Scan scans all records. The traversing begins from given page id.
func (f *HeapFile) Scan() ([]*table.Record, error) {

	// Prepare the result variable.
	ret := make([]*table.Record, 0)

	// Traverse the linked list of pages.
	p := &page.DataPage{}
	next := f.rootPid

	for next != -1 {

		// Read the page.
		err := f.bm.Read(next, p)
		if err != nil {
			return nil, err
		}

		// Read records.
		for i := 0; i < int(p.NumRecords()); i++ {
			r := p.ReadRecord(int64(i))
			b, err := f.schema.DeserializeRecord(r)
			if err != nil {
				return nil, err
			}
			ret = append(ret, b)
		}

		next = p.Next()
	}

	return ret, nil
}

// Insert inserts a record into the page.
func (f *HeapFile) Insert(record *table.Record) error {

	serialized, err := f.schema.SerializeRecord(record)
	if err != nil {
		return err
	}

	next := f.rootPid
	for {

		p := &page.DataPage{}

		err = f.bm.Read(next, p)
		if err != nil {
			return err
		}

		// Insert the record into this page.
		if p.HasFreeSpace() {
			p.AddRecord(serialized)
			f.bm.Update(p.Pid(), p)
			return nil
		}

		// Follow the link to the next page.
		if p.Next() != -1 {
			next = p.Next()
			continue
		}

		// Create the new page and insert the record into it.
		newPage := page.NewDataPage(-1, p.Pid(), -1, f.schema.RecordSize())
		newPage.AddRecord(serialized)
		newPid, err := f.bm.Create(newPage)
		if err != nil {
			return err
		}
		// Set the link to the next page.
		if err = f.bm.Read(p.Pid(), p); err != nil {
			return err
		}
		p.SetNext(newPid)
		f.bm.Update(p.Pid(), p)
		break
	}

	return nil
}

// TODO: impl
func SearchEq(rootPid int64) error {
	return errors.New("not implemented")
}

// TODO: impl
func SearchRange(rootPid int64) error {
	return errors.New("not implemented")
}

// TODO: impl
func Delete(rootPid int64) error {
	return errors.New("not implemented")
}
