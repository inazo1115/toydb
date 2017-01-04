// for test
package main

import (
	"fmt"
	"runtime/debug"

	"github.com/inazo1115/toydb/lib/file"
	//"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/table"
)

func log(msg interface{}) {
	fmt.Println(msg)
}

func main() {

	log("start")

	// Init
	cols := make([]*table.Column, 2)
	cols[0] = table.NewColumnString("name")
	cols[1] = table.NewColumnInt64("age")
	schame := table.NewSchema(cols)
	bm := storage.NewBufferManager()
	hf := file.NewHeapFile(bm, schame)
	rootPid := hf.RootPid()

	// Insert
	for i := 0; i < 500; i++ {
		vals := make([]*table.Value, 2)
		vals[0] = table.NewValueString(fmt.Sprintf("name%d", i))
		vals[1] = table.NewValueInt64(int64(i))
		record := table.NewRecord(vals)
		err := hf.Insert(int64(rootPid), record)
		if err != nil {
			debug.PrintStack()
			panic(err)
		}
	}

	hf.WriteBackAll()

	// Scan
	/*res, err := hf.Scan(int64(rootPid))
	if err != nil {
		panic(err)
	}
	log(res)*/

	log("end")
}

/*func main() {

	log("start")

	// Init
	ba := storage.NewBufferManager()
	p := page.NewDataPage(-1, -1, -1)
	rootPid, err := ba.Create(p)
	if err != nil {
		panic(err)
	}
	ba.WriteBackAll()

	h := file.NewHeapFile()

	// Insert
	for i := 0; i < 300; i++ {
		s := fmt.Sprintf("foofoofoo%d", i)
		err = h.Insert(int64(rootPid), s)
		if err != nil {
			debug.PrintStack()
			panic(err)
		}
	}

	h.WriteBackAll()

	// Scan
	res, err := h.Scan(0)
	if err != nil {
		panic(err)
	}
	log(res)

	log("end")
}*/
