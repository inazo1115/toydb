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
	schema := table.NewSchema([]*table.Column{
		table.NewColumnString("name", 20),
		table.NewColumnInt64("age"),
	})

	bm := storage.NewBufferManager()
	hf := file.NewHeapFile(bm, schema)

	// Insert
	for i := 0; i < 2000; i++ {
		record := table.NewRecord([]*table.Value{
			table.NewValueString("foofoobarbar"),
			table.NewValueInt64(2018),
		})
		err := hf.Insert(record)
		if err != nil {
			panic(err)
		}
	}

	hf.WriteBackAll()

	// Scan
	res, err := hf.Scan()
	if err != nil {
		panic(err)
	}

	table.PrintResult(schema, res)
}
