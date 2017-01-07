package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/inazo1115/toydb/lib/client"
	"github.com/inazo1115/toydb/lib/query"
	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/table"
)

type Runtime struct {
	bm *storage.BufferManager
	ev *query.Evaluator
}

func NewRuntime() *Runtime {
	bm := storage.NewBufferManager()
	ev := query.NewEvaluator(bm)
	return &Runtime{bm, ev}
}

func CleanInput(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.TrimRight(s, ";")
	return s
}

func main() {
	fmt.Printf("version: %s\n", client.Version())

	reader := bufio.NewReader(os.Stdin)
	runtime := NewRuntime()

	for {
		fmt.Printf("toydb> ")
		q, err := reader.ReadString(';')
		if err != nil {
			// panic(err)
			fmt.Println("exit.")
			break
		}
		q = CleanInput(q)

		tokens, err := query.Lex(q)
		if err != nil {
			panic(err)
		}

		ast, err := query.Parse(tokens)
		if err != nil {
			panic(err)
		}

		res, err := runtime.ev.Eval(ast)
		if err != nil {
			panic(err)
		}

		if len(res) == 0 {
			fmt.Println("ok")
		} else {
			schema := runtime.ev.Schema("table_name")
			table.PrintResult(schema, res)
		}
	}
}
