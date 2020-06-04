package file

import (
	//"fmt"
	"os"
	"testing"

	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/table"
	"github.com/inazo1115/toydb/lib/util"
)

func TestInsertAndScan(t *testing.T) {

	// Setup.
	bm := storage.NewBufferManager()
	storage.DataFile = "heap_test_TestInsertAndScan.tmp"
	schema := table.NewSchema([]*table.Column{
		table.NewColumnString("name", 20),
		table.NewColumnInt64("age"),
	})
	hf := NewHeapFile(bm, schema)

	// Insert.
	for i := 0; i < 200; i++ {
		record := table.NewRecord([]*table.Value{
			table.NewValueString("test"),
			table.NewValueInt64(int64(i)),
		})
		err := hf.Insert(record)
		if err != nil {
			t.Errorf("Insert failed.")
		}
	}

	// Scan.
	actual, err := hf.Scan()
	if err != nil {
		t.Errorf("Scan failed.")
	}

	// Assert.
	util.Assert(t, len(actual), 200)
	util.Assert(t, actual[0].Values()[0].Type(), table.STRING)
	util.Assert(t, actual[0].Values()[0].String()[:4], "test")
	util.Assert(t, actual[0].Values()[1].Type(), table.INT64)
	util.Assert(t, actual[0].Values()[1].String()[:1], "0")
	util.Assert(t, actual[1].Values()[0].Type(), table.STRING)
	util.Assert(t, actual[1].Values()[0].String()[:4], "test")
	util.Assert(t, actual[1].Values()[1].Type(), table.INT64)
	util.Assert(t, actual[1].Values()[1].String()[:1], "1")

	// Teardown.
	os.Remove(storage.DataFile)
}
