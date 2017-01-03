package storage

import (
	"os"
	"testing"

	"github.com/inazo1115/toydb/lib/page"
)

// TestCreateAndRead_0 tests that BufferManager can create a new page and read
// it.
func TestCreateAndRead_0(t *testing.T) {

	// Setup.
	bm := NewBufferManager()
	bufferPoolSize_ = 3
	dataFile = "lru_test_TestCreateAndRead_0.tmp"
	message := "test"

	// Create a page.
	p0 := page.NewDataPage(-1, -1, -1)
	p0.AddRecord([]byte(message))
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
	if actual != expected {
		t.Errorf("actual: %s doesn't equals expected: %s.", actual, expected)
	}

	// Teardown.
	os.Remove(dataFile)
}

// TestCreateAndRead_1 tests that BufferManager can create new pages which of
// number their is over cache size and read them.
func TestCreateAndRead_1(t *testing.T) {

	// Setup.
	bm := NewBufferManager()
	bufferPoolSize_ = 3
	dataFile = "lru_test_TestCreateAndRead_1.tmp"
	message := "test"

	// Create a page.
	p0 := page.NewDataPage(-1, -1, -1)
	p0.AddRecord([]byte(message))
	pid, err := bm.Create(p0)
	if err != nil {
		t.Errorf("Create failed.")
	}
	p1 := page.NewDataPage(-1, -1, -1)
	p1.AddRecord([]byte("foo"))
	if _, err := bm.Create(p1); err != nil {
		t.Errorf("Create failed.")
	}
	p2 := page.NewDataPage(-1, -1, -1)
	p2.AddRecord([]byte("foofoo"))
	if _, err := bm.Create(p2); err != nil {
		t.Errorf("Create failed.")
	}
	p3 := page.NewDataPage(-1, -1, -1)
	p3.AddRecord([]byte("foofoofoo"))
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
	if actual != expected {
		t.Errorf("actual: %s doesn't equals expected: %s.", actual, expected)
	}

	// Teardown.
	os.Remove(dataFile)
}

// TestUpdate tests that BufferManager can update a page.
func TestUpdate(t *testing.T) {

	// Setup.
	bm := NewBufferManager()
	bufferPoolSize_ = 3
	dataFile = "lru_test_TestUpdate.tmp"
	message := "test"

	// Create a page.
	p0 := page.NewDataPage(-1, -1, -1)
	p0.AddRecord([]byte("foo"))
	pid, err := bm.Create(p0)
	if err != nil {
		t.Errorf("Create failed.")
	}

	// Update a page.
	p1 := page.NewDataPage(-1, -1, -1)
	p1.AddRecord([]byte(message))
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
	if actual != expected {
		t.Errorf("actual: %s doesn't equals expected: %s.", actual, expected)
	}

	// Teardown.
	os.Remove(dataFile)
}

// TestWriteBackAll tests that BufferManager can write a page back to disk.
func TestWriteBackAll(t *testing.T) {

	// Setup.
	bm0 := NewBufferManager() // Do write process
	bm1 := NewBufferManager() // Do read process
	bufferPoolSize_ = 3
	dataFile = "lru_test_TestWriteBackAll.tmp"
	message := "test"

	// Create a page.
	p0 := page.NewDataPage(-1, -1, -1)
	p0.AddRecord([]byte(message))
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
	if actual != expected {
		t.Errorf("actual: %s doesn't equals expected: %s.", actual, expected)
	}

	// Teardown.
	os.Remove(dataFile)
}
