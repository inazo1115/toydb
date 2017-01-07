// for test
package main

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/query"
	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/table"
)

func log(msg interface{}) {
	fmt.Println(msg)
}

func main() {

	query0 := "create table table_name (name string(20), age int, tel int)"
	query1 := "insert into table_name (name, age, tel) values (\"foofoo\", 100, 200)"
	//	query1 := insert into table_name (name, age, tel) values ("foofoo", 100, 200)
	query2 := "select * from table_name"

	token0, err := query.Lex(query0)
	if err != nil {
		panic(err)
	}
	token1, err := query.Lex(query1)
	if err != nil {
		panic(err)
	}
	token2, err := query.Lex(query2)
	if err != nil {
		panic(err)
	}

	ast0, err := query.Parse(token0)
	if err != nil {
		panic(err)
	}
	ast1, err := query.Parse(token1)
	if err != nil {
		panic(err)
	}
	ast2, err := query.Parse(token2)
	if err != nil {
		panic(err)
	}

	bm := storage.NewBufferManager()
	eva := query.NewEvaluator(bm)

	_, err = eva.Eval(ast0)
	if err != nil {
		panic(err)
	}
	_, err = eva.Eval(ast1)
	if err != nil {
		panic(err)
	}
	res, err := eva.Eval(ast2)
	if err != nil {
		panic(err)
	}

	schema := eva.Schema("table_name")
	table.PrintResult(schema, res)
}
