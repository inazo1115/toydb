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
	pa := storage.NewPageAccessor()
	bm := storage.NewBufferManager()

	bm.Dump()

	data := []byte("datadatadatadata")
	page := model.NewPage(0, data)

	pa.Read(0)
	pa.Read(1)

	pid, err := pa.Create(page)
	if err != nil {
		panic(err)
	}
	bm.Dump()

	page, err = pa.Read(pid)
	page, _ = pa.Read(pid)

	pa.WriteBackAll()

	bm.Dump()
}


/*func main() {
	dm := storage.NewDiskManager()
	buffer := make([]byte, pkg.BlockSize)

	fmt.Println("++ fetch ++")
	dm.Fetch(0, buffer)
	fmt.Println(buffer)

	fmt.Println("++ update ++")
	data := []byte("datadatadatadata")
	dm.Update(1, data)

	fmt.Println("++ fetch ++")
	dm.Fetch(0, buffer)
	fmt.Println(buffer)
	fmt.Println("++ fetch ++")
	dm.Fetch(1, buffer)
	fmt.Println(buffer)

	fmt.Println("++ next ++")
	i, err := dm.NextPageID()
	fmt.Println(i)
	fmt.Println(err)

	//fmt.Println("++ insert ++")
	//data := []byte("hoge")
	//dm.Insert(data)

	fmt.Println("++ print ++")
	dm.PrintPage(0)
}*/
