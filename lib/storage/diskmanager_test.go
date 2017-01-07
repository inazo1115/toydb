package storage

import (
	"os"
	"testing"
)

// TestWriteAndRead tests that DiskManager can write the message to the file and
// read it.
func TestWriteAndRead(t *testing.T) {

	// Setup.
	dm := NewDiskManager()
	DataFile = "diskmanager_test_TestWriteAndRead.tmp"
	expected := "this is the test message."

	// Write.
	if err := dm.Write(0, []byte(expected)); err != nil {
		t.Errorf("Write faild.")
	}

	// Read.
	size, err := dm.GetBufferSize(0)
	if err != nil {
		t.Errorf("GetBufferSize faild.")
	}
	buf := make([]byte, size)
	if err = dm.Read(0, buf); err != nil {
		t.Errorf("Read faild.")
	}

	// Assert.
	actual := string(buf)
	if actual != expected {
		t.Errorf("actual: %s doesn't equals expected: %s", actual, expected)
	}

	// Teardown.
	os.Remove(DataFile)
}

// TestGetFreePageID_0 tests that DiskManager returns first page id(0) when
// there is no pages on the disk and the buffer.
func TestGetFreePageID_0(t *testing.T) {

	// Setup.
	dm := NewDiskManager()
	DataFile = "diskmanager_test_TestGetFreePageID_0.tmp"
	expected := int64(0)

	// Get the page id.
	pid, err := dm.GetFreePageID(make([]int64, 0))
	if err != nil {
		t.Errorf("GetFreePageID failed.")
	}

	// Assert.
	actual := pid
	if actual != expected {
		t.Errorf("actual: %d doesn't equals expected: %d.", actual, expected)
	}

	// Teardown.
	os.Remove(DataFile)
}

// TestGetFreePageID_1 tests that DiskManager returns next of the maximum page
// id when the page is on the cache.
func TestGetFreePageID_1(t *testing.T) {

	// Setup.
	dm := NewDiskManager()
	DataFile = "diskmanager_test_TestGetFreePageID_1.tmp"
	expected := int64(5)

	// Get the page id.
	pagesOnCache := []int64{0, 2, 4}
	pid, err := dm.GetFreePageID(pagesOnCache)
	if err != nil {
		t.Errorf("GetFreePageID failed.")
	}

	// Assert.
	actual := pid
	if actual != expected {
		t.Errorf("actual: %d doesn't equals expected: %d.", actual, expected)
	}

	// Teardown.
	os.Remove(DataFile)
}

// TestGetFreePageID_2 tests that DiskManager returns next of the maximum page
// id when the page is on the disk.
func TestGetFreePageID_2(t *testing.T) {

	// Setup.
	dm := NewDiskManager()
	DataFile = "diskmanager_test_TestGetFreePageID_2.tmp"
	expected := int64(1)

	// Write.
	if err := dm.Write(0, []byte("this is the test message.")); err != nil {
		t.Errorf("Write faild.")
	}

	// Get the page id.
	pid, err := dm.GetFreePageID(make([]int64, 0))
	if err != nil {
		t.Errorf("GetFreePageID failed.")
	}

	// Assert.
	actual := pid
	if actual != expected {
		t.Errorf("actual: %d doesn't equals expected: %d.", actual, expected)
	}

	// Teardown.
	os.Remove(DataFile)
}

// TestGetBufferSize_0 tests that DiskManager returns the buffer size. When the
// placement of the page is the last of file, the buffer size equals the page's
// contents size.
func TestGetBufferSize_0(t *testing.T) {

	// Setup.
	dm := NewDiskManager()
	DataFile = "diskmanager_test_TestGetBufferSize_0.tmp"
	message := "this is the test message."
	expected := int64(len(message))

	// Write.
	if err := dm.Write(0, []byte(message)); err != nil {
		t.Errorf("Write faild.")
	}

	// Get the buffer size.
	size, err := dm.GetBufferSize(0)
	if err != nil {
		t.Errorf("GetBufferSize failed.")
	}

	// Assert.
	actual := size
	if actual != expected {
		t.Errorf("actual: %d doesn't equals expected: %d.", actual, expected)
	}

	// Teardown.
	os.Remove(DataFile)
}

// TestGetBufferSize_1 tests that DiskManager returns the buffer size. When the
// placement of the page is not the last of file, the buffer size is a precise
// size which equals the block size.
func TestGetBufferSize_1(t *testing.T) {

	// Setup.
	dm := NewDiskManager()
	DataFile = "diskmanager_test_TestGetBufferSize_1.tmp"
	message := "this is the test message."
	expected := int64(4096) // It's the block size.

	// Write twice.
	if err := dm.Write(0, []byte(message)); err != nil {
		t.Errorf("Write faild.")
	}
	if err := dm.Write(1, []byte(message)); err != nil {
		t.Errorf("Write faild.")
	}

	// Get the buffer size.
	size, err := dm.GetBufferSize(0)
	if err != nil {
		t.Errorf("GetBufferSize failed.")
	}

	// Assert.
	actual := size
	if actual != expected {
		t.Errorf("actual: %d doesn't equals expected: %d.", actual, expected)
	}

	// Teardown.
	os.Remove(DataFile)
}
