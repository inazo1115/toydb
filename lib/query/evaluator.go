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

func NewEvaluator() *Evaluator {
	return &Evaluator{NewRuntime()}
}

func (e *Evaluator) Runtime() *Runtime {
	return e.runtime
}

func (e *Evaluator) Eval(ast *ASTNode) ([]*table.Record, error) {

	switch {
	case ast.Token.ID == TokenCREATE:
		return nil, e.EvalCreateTable(ast)

	case ast.Token.ID == TokenINSERT:
		return nil, e.EvalInsert(ast)

	case ast.Token.ID == TokenSELECT:
		return e.EvalSelect(ast)

	}
	return nil, errors.New("Eval")
}

func (e *Evaluator) EvalCreateTable(ast *ASTNode) error {

	tableName := ast.Children[0].Token.Val
	e.runtime.SetCurrentTableName(tableName) // tmp

	cols := make([]*table.Column, 0)

	for i := 1; i < len(ast.Children); i += 2 {

		col := ast.Children[i]
		typ := ast.Children[i+1]

		var newCol *table.Column
		switch typ.Token.ID {
		case TokenSTRING:
			intv, err := strconv.Atoi(typ.Children[0].Token.Val)
			if err != nil {
				return err
			}
			size := int32(intv)
			newCol = table.NewColumnString(col.Token.Val, size)
		case TokenINT:
			newCol = table.NewColumnInt64(col.Token.Val)
		}

		cols = append(cols, newCol)
	}

	s := table.NewSchema(cols)
	f := file.NewHeapFile(e.runtime.BufferManager(), s)

	e.runtime.SetSchema(tableName, s)
	e.runtime.SetFile(tableName, f)

	return nil
}

func (e *Evaluator) EvalInsert(ast *ASTNode) error {

	tableName := ast.Children[0].Token.Val
	e.runtime.SetCurrentTableName(tableName) // tmp

	vals := make([]*table.Value, 0)

	for i := 1; i < len(ast.Children); i++ {

		col := ast.Children[i]
		val := ast.Children[i].Children[0]

		var newVal *table.Value
		t, err := e.runtime.CurrentSchema().Type(col.Token.Val)
		if err != nil {
			return err
		}
		switch t {
		case table.STRING:
			newVal = table.NewValueString(string(val.Token.Val))
		case table.INT64:
			intv, err := strconv.Atoi(val.Token.Val)
			if err != nil {
				return err
			}
			newVal = table.NewValueInt64(int64(intv))
		}
		vals = append(vals, newVal)
	}

	record := table.NewRecord(vals)
	if err := e.runtime.CurrentFile().Insert(record); err != nil {
		return err
	}

	return nil
}

func (e *Evaluator) EvalSelect(ast *ASTNode) ([]*table.Record, error) {
	tableName := ast.Children[0].Token.Val
	e.runtime.SetCurrentTableName(tableName) // tmp
	return e.runtime.CurrentFile().Scan()
}
