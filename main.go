// for test
package main

import (
	"fmt"
	
	"github.com/inazo1115/toydb/lib/storage"
	//"github.com/inazo1115/toydb/lib/pkg"
)

func main() {
	bm := storage.NewBufferManager()

	bm.Dump()

	fmt.Println("++ fetch ++")
	bm.Fetch(0)
	fmt.Println("++ fetch ++")
	bm.Fetch(1)

	bm.Dump()

	fmt.Println("++ fetch ++")
	bm.Fetch(2)

	bm.Dump()

	fmt.Println("++ fetch ++")
	bm.Fetch(0)
	fmt.Println("++ fetch ++")
	bm.Fetch(1)
	fmt.Println("++ fetch ++")
	bm.Fetch(3)

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
