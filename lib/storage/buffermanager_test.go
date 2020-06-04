package storage

import (
	"os"
	"testing"

	"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/util"
)

// TestCreateAndRead_0 tests that BufferManager can create a new page and read
// it.
func TestCreateAndRead_0(t *testing.T) {

	// Setup.
	bm := NewBufferManager()
	bufferPoolSize_ = 3
	DataFile = "lru_test_TestCreateAndRead_0.tmp"
	recordSize := 10
	data := make([]byte, recordSize)
	message := "test"
	util.CopyStrToByte(message, data)

	// Create a page.
	p0 := page.NewDataPage(-1, -1, -1, int64(recordSize))
	err := p0.AddRecord(data)
	if err != nil {
		t.Errorf("AddRecord failed.")
	}
	pid, err := bm.Create(p0)
	if err != nil {
		t.Errorf("Create failed.")
	}

	// Read a page.
	p1 := &page.DataPage{}
	if err = bm.Read(pid, p1); err != nil {
		t.Errorf("Read failed.")
	}

	// Assert.
	actual := string(p1.Data()[:len(message)]) // TODO: len is a bad hack. fix it.
	expected := message
	util.Assert(t, actual, expected)

	// Teardown.
	os.Remove(DataFile)
}

// TestCreateAndRead_1 tests that BufferManager can create new pages which of
// number their is over cache size and read them.
func TestCreateAndRead_1(t *testing.T) {

	// Setup.
	bm := NewBufferManager()
	bufferPoolSize_ = 3
	DataFile = "lru_test_TestCreateAndRead_1.tmp"
	recordSize := 10
	data := make([]byte, recordSize)
	message := "test"
	util.CopyStrToByte(message, data)

	// Create a page.
	p0 := page.NewDataPage(-1, -1, -1, int64(recordSize))
	p0.AddRecord(data)
	pid, err := bm.Create(p0)
	if err != nil {
		t.Errorf("Create failed.")
	}

	p1 := page.NewDataPage(-1, -1, -1, int64(recordSize))
	data = make([]byte, recordSize)
	util.CopyStrToByte("foo", data)
	p1.AddRecord(data)
	if _, err := bm.Create(p1); err != nil {
		t.Errorf("Create failed.")
	}

	p2 := page.NewDataPage(-1, -1, -1, int64(recordSize))
	data = make([]byte, recordSize)
	util.CopyStrToByte("foofoo", data)
	p2.AddRecord(data)
	if _, err := bm.Create(p2); err != nil {
		t.Errorf("Create failed.")
	}

	p3 := page.NewDataPage(-1, -1, -1, int64(recordSize))
	data = make([]byte, recordSize)
	util.CopyStrToByte("foofoofoo", data)
	p3.AddRecord(data)
	if _, err := bm.Create(p3); err != nil {
		t.Errorf("Create failed.")
	}

	// Read a page.
	p := &page.DataPage{}
	if err = bm.Read(pid, p); err != nil {
		t.Errorf("Read failed.")
	}

	// Assert.
	actual := string(p.Data()[:len(message)]) // TODO: len is a bad hack. fix it.
	expected := message
	util.Assert(t, actual, expected)

	// Teardown.
	os.Remove(DataFile)
}

// TestUpdate tests that BufferManager can update a page.
func TestUpdate(t *testing.T) {

	// Setup.
	bm := NewBufferManager()
	bufferPoolSize_ = 3
	DataFile = "lru_test_TestUpdate.tmp"
	recordSize := 10

	// Create a page.
	p0 := page.NewDataPage(-1, -1, -1, int64(recordSize))
	data := make([]byte, recordSize)
	util.CopyStrToByte("foo", data)
	p0.AddRecord(data)
	pid, err := bm.Create(p0)
	if err != nil {
		t.Errorf("Create failed.")
	}

	// Update a page.
	p1 := page.NewDataPage(-1, -1, -1, int64(recordSize))
	data = make([]byte, recordSize)
	message := "test"
	util.CopyStrToByte(message, data)
	p1.AddRecord(data)
	if err = bm.Update(pid, p1); err != nil {
		t.Errorf("Create failed.")
	}

	// Read a page.
	p2 := &page.DataPage{}
	if err = bm.Read(pid, p2); err != nil {
		t.Errorf("Read failed.")
	}

	// Assert.
	actual := string(p2.Data()[:len(message)]) // TODO: len is a bad hack. fix it.
	expected := message
	util.Assert(t, actual, expected)

	// Teardown.
	os.Remove(DataFile)
}

// TestWriteBackAll tests that BufferManager can write a page back to disk.
func TestWriteBackAll(t *testing.T) {

	// Setup.
	bm0 := NewBufferManager() // Do write process
	bm1 := NewBufferManager() // Do read process
	bufferPoolSize_ = 3
	DataFile = "lru_test_TestWriteBackAll.tmp"
	recordSize := 10
	data := make([]byte, recordSize)
	message := "test"
	util.CopyStrToByte(message, data)

	// Create a page.
	p0 := page.NewDataPage(-1, -1, -1, int64(recordSize))
	p0.AddRecord(data)
	pid, err := bm0.Create(p0)
	if err != nil {
		t.Errorf("Create failed.")
	}

	// At this time, the read process will fail.
	p1 := &page.DataPage{}
	if err = bm1.Read(pid, p1); err == nil {
		t.Errorf("Read would fail.")
	}

	// Write back.
	bm0.WriteBackAll()

	// Read a page.
	p2 := &page.DataPage{}
	if err = bm1.Read(pid, p2); err != nil {
		t.Errorf("Read failed.")
	}

	// Assert.
	actual := string(p2.Data()[:len(message)]) // TODO: len is a bad hack. fix it.
	expected := message
	util.Assert(t, actual, expected)

	// Teardown.
	os.Remove(DataFile)
}
