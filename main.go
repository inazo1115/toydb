// for test
package main

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/page"
	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/file"
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


	// Insert
	h := file.NewHeapFile()

	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("foofoofoo%d", i)
		err = h.Insert(int64(rootPid), s)
		if err != nil {
			panic(err)
		}
		//log("*********************************")
		//log(i)
		//h.Dump(0)
	}
	//h.Dump(1)

	// Scan
	//res, err := h.Scan(0)
	//if err != nil {
	//panic(err)
//}

	//log(res)

	ba.WriteBackAll()
}

/*func main() {
	log("start")

	ba := storage.NewBufferManager()

	p := page.NewDataPage(-1, -1, 1)
	p.AddRecord([]byte("0-0"))
	p.AddRecord([]byte("0-1"))
	p.AddRecord([]byte("0-2"))
	_, err := ba.Create(p)
	if err != nil {
		panic(err)
	}

	p = page.NewDataPage(-1, 0, -1)
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
	for i := 0; i < int(pp.NumRecords()); i++ {
		log(string(pp.ReadRecord(i)))
	}

	ba.Read(1, pp)
	for i := 0; i < int(pp.NumRecords()); i++ {
		log(string(pp.ReadRecord(i)))
	}

	h := file.NewHeapFile()
	res := h.Scan(0)
	log(res)

	ba.WriteBackAll()
}*/
