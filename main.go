// for test
package main

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/storage"
)

func log(msg interface{}) {
	fmt.Println(msg)
}

func main() {
	log("start")

	ba := storage.NewBufferManager()

	p := page.NewDataPage(-1, -1, -1)
	p.AddRecord([]byte("0-0"))
	p.AddRecord([]byte("0-1"))
	p.AddRecord([]byte("0-2"))
	_, err := ba.Create(p)
	if err != nil {
		panic(err)
	}

	p = page.NewDataPage(-1, -1, -1)
	p.AddRecord([]byte("1-0"))
	p.AddRecord([]byte("1-1"))
	p.AddRecord([]byte("1-2"))
	_, err = ba.Create(p)
	if err != nil {
		panic(err)
	}

	pp := &page.DataPage{}
	err = ba.Read(1, pp)
	if err != nil {
		panic(err)
	}
	err = ba.Read(0, pp)
	if err != nil {
		panic(err)
	}
	err = ba.Read(1, pp)
	if err != nil {
		panic(err)
	}
	err = ba.Read(0, pp)
	if err != nil {
		panic(err)
	}
	err = ba.Read(1, pp)
	if err != nil {
		panic(err)
	}

	ba.Read(0, pp)
	log(string(pp.Data()))

	ba.Read(1, pp)
	log(string(pp.Data()))

	ba.WriteBackAll()
}
