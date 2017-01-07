package query

import (
	"errors"
	//"fmt"
	"strconv"

	"github.com/inazo1115/toydb/lib/file"
	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/table"
)

type Evaluator struct {
	bm     *storage.BufferManager
	tables map[string]*file.HeapFile
}

func NewEvaluator(bm *storage.BufferManager) *Evaluator {
	return &Evaluator{bm, make(map[string]*file.HeapFile, 0)}
}

func (e *Evaluator) Schema(name string) *table.Schema {
	return e.tables[name].Schema()
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

	schema := table.NewSchema(cols)
	e.tables[tableName] = file.NewHeapFile(e.bm, schema)

	return nil
}

func (e *Evaluator) EvalInsert(ast *ASTNode) error {

	tableName := ast.Children[0].Token.Val

	vals := make([]*table.Value, 0)

	for i := 1; i < len(ast.Children); i++ {

		col := ast.Children[i]
		val := ast.Children[i].Children[0]

		var newVal *table.Value
		t, err := e.tables[tableName].Schema().Type(col.Token.Val)
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
	file := e.tables[tableName]
	err := file.Insert(record)
	if err != nil {
		return err
	}
	return nil
}

func (e *Evaluator) EvalSelect(ast *ASTNode) ([]*table.Record, error) {
	tableName := ast.Children[0].Token.Val
	file := e.tables[tableName]
	return file.Scan()
}
