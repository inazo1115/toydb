package file

import (
	"errors"
	"fmt"

	"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/storage"
)

// HeapFile is the representation of the file and access methods. HeapFile is
// the linked list of pages.
type HeapFile struct {
	bm *storage.BufferManager
}

// NewHeapFile creates HeapFile struct and returns it's pointer.
func NewHeapFile() *HeapFile {
	return &HeapFile{storage.NewBufferManager()}
}

// Scan scans all records. The traversing begins from given page id.
func (f *HeapFile) Scan(pid int64) ([]string, error) {

	// Prepare the result variable.
	ret := make([]string, 0)

	// Traverse the linked list of pages.
	p := &page.DataPage{}
	next := pid
	for next != -1 {

		// Read the page.
		err := f.bm.Read(next, p)
		if err != nil {
			return nil, err
		}

		// Read records.
		for i := 0; i < int(p.NumRecords()); i++ {
			rec := deserializeRecord(p.ReadRecord(i))
			ret = append(ret, rec)
		}

		next = p.Next()
	}

	return ret, nil
}

// Insert inserts a record into the page.
func (f *HeapFile) Insert(pid int64, record string) error {

	p := &page.DataPage{}

	err := f.bm.Read(pid, p)
	if err != nil {
		return err
	}

	// Insert the record into this page.
	if p.HasFreeSpace() {
		p.AddRecord(serializeRecord(record))
		f.bm.Update(p.Pid(), p)
		return nil
	}

	// Follow the link to the next page.
	if p.Next() != -1 {
		return f.Insert(p.Next(), record)
	}

	// Create the new page and insert the record into it.
	newPage := page.NewDataPage(-1, p.Pid(), -1)
	newPage.AddRecord(serializeRecord(record))
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

	return nil
}

// Dump is a debug function.
func (f *HeapFile) Dump(pid int64) {

	p := &page.DataPage{}

	var err error
	err = f.bm.Read(pid, p)
	if err != nil {
		panic(err)
	}

	fmt.Println("dump")
	fmt.Println(pid)
	fmt.Println(p)
}

func SearchEq(rootPid int64) error {
	return errors.New("not implemented")
}

func SearchRange(rootPid int64) error {
	return errors.New("not implemented")
}

func Delete(rootPid int64) error {
	return errors.New("not implemented")
}

func (f *HeapFile) WriteBackAll() error {
	err := f.bm.WriteBackAll()
	if err != nil {
		return err
	}
	return nil
}

// tmp
func deserializeRecord(b []byte) string {
	return string(b)
}

// tmp
func serializeRecord(s string) []byte {
	return []byte(s)
}
