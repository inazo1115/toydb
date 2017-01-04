// for test
package main

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/file"
	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/table"
)

func log(msg interface{}) {
	fmt.Println(msg)
}

func main() {

	// Init
	cols := make([]*table.Column, 2)
	cols[0] = table.NewColumnString("name", 20)
	cols[1] = table.NewColumnInt64("age")
	schema := table.NewSchema(cols)
	bm := storage.NewBufferManager()
	hf := file.NewHeapFile(bm, schema)
	rootPid := hf.RootPid()

	// Insert
	for i := 0; i < 3000; i++ {
		vals := make([]*table.Value, 2)
		vals[0] = table.NewValueString(fmt.Sprintf("name%d", i))
		vals[1] = table.NewValueInt64(int64(i))
		record := table.NewRecord(vals)
		err := hf.Insert(int64(rootPid), record)
		if err != nil {
			panic(err)
		}
	}

	hf.WriteBackAll()

	// Scan
	res, err := hf.Scan(int64(rootPid))
	if err != nil {
		panic(err)
	}

	table.PrintResult(schema, res)
}
