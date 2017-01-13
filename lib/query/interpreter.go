package query

import (
	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/table"
)

type Interpreter struct {
	runtime *Runtime
	ctx     *Context
}

func NewInterpreter() *Interpreter {
	bufman := storage.NewBufferManager()
	diskman := storage.NewDiskManager()
	runtime := NewRuntime(bufman, diskman)
	return &Interpreter{runtime, NewContext()}
}

func (i *Interpreter) Interpret(q string) ([]*table.Record, error) {

	tokens, err := NewLexer(q).Lex()
	if err != nil {
		return nil, err
	}

	ast, err := NewParser().Parse(tokens)
	if err != nil {
		return nil, err
	}

	result, err := NewEvaluator(i.runtime).Eval(ast, i.ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i *Interpreter) PrintResult(result []*table.Record) {
	table.PrintResult(i.ctx.CurrentSchema(), result)
}
