package query

import (
	"errors"
	//"fmt"
	"strconv"

	"github.com/inazo1115/toydb/lib/file"
	"github.com/inazo1115/toydb/lib/table"
)

type Evaluator struct {
	runtime *Runtime
}

func NewEvaluator(runtime *Runtime) *Evaluator {
	return &Evaluator{runtime}
}

func (e *Evaluator) Eval(ast *ASTNode, ctx *Context) ([]*table.Record, error) {
	switch ast.Token.ID() {
	case TokenCREATE:
		return make([]*table.Record, 0), e.evalCreateTable(ast, ctx)
	case TokenINSERT:
		return make([]*table.Record, 0), e.evalInsert(ast, ctx)
	case TokenSELECT:
		return e.evalSelect(ast, ctx)
	default:
		return nil, errors.New("error: Eval")
	}
}

func (e *Evaluator) evalCreateTable(ast *ASTNode, ctx *Context) error {

	tableName := ast.Children[0].Token.Val()
	ctx.SetCurrentTable(tableName) // tmp

	cols := make([]*table.Column, 0)

	for i := 1; i < len(ast.Children); i += 2 {

		col := ast.Children[i]
		typ := ast.Children[i+1]

		var newCol *table.Column
		switch typ.Token.ID() {
		case TokenSTRING:
			intv, err := strconv.Atoi(typ.Children[0].Token.Val())
			if err != nil {
				return err
			}
			size := int32(intv)
			newCol = table.NewColumnString(col.Token.Val(), size)
		case TokenINT:
			newCol = table.NewColumnInt64(col.Token.Val())
		}

		cols = append(cols, newCol)
	}

	s := table.NewSchema(cols)
	f := file.NewHeapFile(e.runtime.BufferManager(), s)

	ctx.SetSchema(tableName, s)
	ctx.SetFile(tableName, f)

	return nil
}

func (e *Evaluator) evalInsert(ast *ASTNode, ctx *Context) error {

	tableName := ast.Children[0].Token.Val()
	ctx.SetCurrentTable(tableName) // tmp

	vals := make([]*table.Value, 0)

	for i := 1; i < len(ast.Children); i++ {

		col := ast.Children[i]
		val := ast.Children[i].Children[0]

		var newVal *table.Value
		t, err := ctx.CurrentSchema().Type(col.Token.Val())
		if err != nil {
			return err
		}
		switch t {
		case table.STRING:
			newVal = table.NewValueString(string(val.Token.Val()))
		case table.INT64:
			intv, err := strconv.Atoi(val.Token.Val())
			if err != nil {
				return err
			}
			newVal = table.NewValueInt64(int64(intv))
		}
		vals = append(vals, newVal)
	}

	record := table.NewRecord(vals)
	if err := ctx.CurrentFile().Insert(record); err != nil {
		return err
	}

	return nil
}

func (e *Evaluator) evalSelect(ast *ASTNode, ctx *Context) ([]*table.Record, error) {
	tableName := ast.Children[0].Token.Val()
	ctx.SetCurrentTable(tableName) // tmp
	return ctx.CurrentFile().Scan()
}
