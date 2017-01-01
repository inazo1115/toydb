// for test
package main

import (
	//"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	
	//"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/storage/model"
)

func log(msg interface{}) {
	fmt.Println(msg)
}

func main() {
	log("start")

	page := model.NewPage(100)

	log("*")
	log(binary.Size(page))

	//buf := new(bytes.Buffer)
	//buf := new(bytes.Buffer)
	//err := binary.Write(buf, binary.LittleEndian, page)
	buf, err := page.MarshalBinary()
	if err != nil {
		log(err)
	}
	fmt.Printf("% x", buf)

	log("")
	log("=============")

	page0 := &model.Page{}
	
	//reader := bytes.NewReader(buf.Bytes())

	//v := reflect.Indirect(reflect.ValueOf(page0))
	log("^^")
	//log(reflect.ValueOf(page0))
	//log(reflect.TypeOf(reflect.ValueOf(page0)))
	log(reflect.ValueOf(page0).Kind())
	log(reflect.ValueOf(page0).Elem())
	//log(intDestSize(page0))

	page0.UnmarshalBinary(buf)
	//err = binary.Read(reader, binary.LittleEndian, page0)
	//if err != nil {
	//log(err)
//}
	log(page0)
	log(page0.Pid())
}











/*func main() {
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
}*/


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
