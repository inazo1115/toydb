package main

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/client"
)

func main() {

	query0 := "create table table_name (name string(20), age int, tel int)"
	query1 := "insert into table_name (name, age, tel) values (%s, %d, %d)"
	query2 := "select * from table_name"

	c := client.NewClient()
	c.Query(query0)
	fmt.Println(query0)
	for i := 0; i < 5000; i++ {
		q := fmt.Sprintf(query1, "\"foo\"", i, 200)
		fmt.Println(q)
		c.Query(q)
	}
	c.Query(query2)
	fmt.Println(query2)
}
