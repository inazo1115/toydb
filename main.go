// for test
package main

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/storage/model"
)

func log(msg interface{}) {
	fmt.Println(msg)
}

func main() {
	log("start")

	pa := storage.NewPageAccessor()

	page := model.NewPage(-1, -1, -1)
	page.AddRecord([]byte("0-0"))
	page.AddRecord([]byte("0-1"))
	page.AddRecord([]byte("0-2"))
	_, err := pa.Create(page)
	if err != nil {
		panic(err)
	}

	page = model.NewPage(-1, -1, -1)
	page.AddRecord([]byte("1-0"))
	page.AddRecord([]byte("1-1"))
	page.AddRecord([]byte("1-2"))
	_, err = pa.Create(page)
	if err != nil {
		panic(err)
	}

	_, err = pa.Read(1)
	if err != nil {
		panic(err)
	}
	_, err = pa.Read(0)
	if err != nil {
		panic(err)
	}
	_, err = pa.Read(1)
	if err != nil {
		panic(err)
	}
	_, err = pa.Read(0)
	if err != nil {
		panic(err)
	}
	_, err = pa.Read(1)
	if err != nil {
		panic(err)
	}

	ppp, _ := pa.Read(0)
	log(string(ppp.Data()))

	ppp, _ = pa.Read(1)
	log(string(ppp.Data()))

	pa.WriteBackAll()
}
