// for test
package main

import (
	"fmt"
	"runtime/debug"

	"github.com/inazo1115/toydb/lib/file"
	"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/storage"
)

func log(msg interface{}) {
	fmt.Println(msg)
}

func main() {

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
	for i := 0; i < 3000; i++ {
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
}
