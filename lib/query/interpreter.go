package query

import (
	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/table"
)

type Interpreter struct {
	lexer     *Lexer
	parser    *Parser
	evaluator *Evaluator
	ctx       *Context
}

func NewInterpreter() *Interpreter {
	bufman := storage.NewBufferManager()
	diskman := storage.NewDiskManager()
	runtime := NewRuntime(bufman, diskman)
	return &Interpreter{NewLexer(), NewParser(), NewEvaluator(runtime), NewContext()}
}

func (i *Interpreter) Interpret(q string) ([]*table.Record, error) {

	tokens, err := i.lexer.Lex(q)
	if err != nil {
		return nil, err
	}

	ast, err := i.parser.Parse(tokens)
	if err != nil {
		return nil, err
	}

	result, err := i.evaluator.Eval(ast, i.ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i *Interpreter) PrintResult(result []*table.Record) {
	table.PrintResult(i.ctx.CurrentSchema(), result)
}
