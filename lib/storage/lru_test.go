package storage

import (
	"os"
	"testing"

	"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/util"
)

// TestTouchPage_0 tests frame's hit count. The access pattern is create new
// pages.
func TestTouchPage_0(t *testing.T) {

	// Setup.
	bm := NewBufferManager()
	bufferPoolSize_ = 3
	DataFile = "lru_test_TestTouchPage_0.tmp"

	// Page accesses.
	p0 := page.NewDataPage(-1, -1, -1, 10)
	pid0, err := bm.Create(p0)
	if err != nil {
		t.Errorf("Create failed.")
	}

	p1 := page.NewDataPage(-1, -1, -1, 10)
	pid1, err := bm.Create(p1)
	if err != nil {
		t.Errorf("Create failed.")
	}

	p2 := page.NewDataPage(-1, -1, -1, 10)
	pid2, err := bm.Create(p2)
	if err != nil {
		t.Errorf("Create failed.")
	}

	// Assert.
	actual := bm.bufferPool[bm.dict[pid0]].HitCount()
	expected := int64(2)
	util.Assert(t, actual, expected)

	actual = bm.bufferPool[bm.dict[pid1]].HitCount()
	expected = int64(1)
	util.Assert(t, actual, expected)

	actual = bm.bufferPool[bm.dict[pid2]].HitCount()
	expected = int64(0)
	util.Assert(t, actual, expected)

	// Teardown.
	os.Remove(DataFile)
}

// TestTouchPage_0 tests frame's hit count. The access pattern is to read one
// page.
func TestTouchPage_1(t *testing.T) {

	// Setup.
	bm := NewBufferManager()
	bufferPoolSize_ = 3
	DataFile = "lru_test_TestTouchPage_1.tmp"

	// Page accesses.
	p0 := page.NewDataPage(-1, -1, -1, 10)
	pid, err := bm.Create(p0)
	if err != nil {
		t.Errorf("Create failed.")
	}

	p1 := &page.DataPage{}
	if err = bm.Read(pid, p1); err != nil {
		t.Errorf("Read failed.")
	}

	p2 := &page.DataPage{}
	if err = bm.Read(pid, p2); err != nil {
		t.Errorf("Read failed.")
	}

	// Assert.
	actual := bm.bufferPool[bm.dict[pid]].HitCount()
	expected := int64(0)
	util.Assert(t, actual, expected)

	// Teardown.
	os.Remove(DataFile)
}

// TestChooseVictim tests the logic that selects the eviction target.
func TestChooseVictim(t *testing.T) {

	// Setup.
	bm := NewBufferManager()
	bufferPoolSize_ = 3
	DataFile = "lru_test_TestChooseVictim.tmp"

	// Page accesses.
	p0 := page.NewDataPage(-1, -1, -1, 10)
	pid0, err := bm.Create(p0)
	if err != nil {
		t.Errorf("Create failed.")
	}

	p1 := page.NewDataPage(-1, -1, -1, 10)
	pid1, err := bm.Create(p1)
	if err != nil {
		t.Errorf("Create failed.")
	}

	p2 := page.NewDataPage(-1, -1, -1, 10)
	pid2, err := bm.Create(p2)
	if err != nil {
		t.Errorf("Create failed.")
	}

	// Assert.
	actual := bm.cacheStrat.ChooseVictim(bm)
	expected := int64(bm.dict[pid0])
	util.Assert(t, actual, expected)

	actual = bm.bufferPool[bm.dict[pid0]].HitCount()
	expected = int64(2)
	util.Assert(t, actual, expected)

	actual = bm.bufferPool[bm.dict[pid1]].HitCount()
	expected = int64(1)
	util.Assert(t, actual, expected)

	actual = bm.bufferPool[bm.dict[pid2]].HitCount()
	expected = int64(0)
	util.Assert(t, actual, expected)

	// Teardown.
	os.Remove(DataFile)
}
