package table

import (
	//"fmt"
	"errors"

	"github.com/inazo1115/toydb/lib/util"
)

// Schema

type Schema struct {
	cols []*Column
}

func NewSchema(cols []*Column) *Schema {
	return &Schema{cols}
}

func (s *Schema) Columns() []*Column {
	return s.cols
}

func (s *Schema) RecordSize() int64 {
	size := int64(0)
	for _, c := range s.cols {
		size += int64(c.colSize)
	}
	return size
}

func (s *Schema) SerializeRecord(r *Record) ([]byte, error) {

	if len(s.cols) != len(r.values) {
		return nil, errors.New("length unmatch")
	}

	ret := make([]byte, s.RecordSize())
	idx := 0

	for i, col := range s.cols {
		val := r.Values()[i]

		if col.Type() != val.Type() {
			return nil, errors.New("type unmatch")
		}

		b := col.SerializeValue(val)
		for j := 0; j < len(b); j++ {
			ret[idx] = b[j]
			idx++
		}
	}

	return ret, nil
}

func (s *Schema) DeserializeRecord(b []byte) (*Record, error) {

	values := make([]*Value, len(s.cols))
	from := 0
	to := 0

	for i, col := range s.cols {
		to += int(col.colSize)
		values[i] = col.DeserializeValue(b[from:to])
		from = to
	}

	return NewRecord(values), nil
}

// Column

type Column struct {
	name    string
	type_   ToyDBType
	colSize int32
}

func NewColumnInt64(name string) *Column {
	return &Column{name, INT64, INT64Size}
}

func NewColumnString(name string, colSize int32) *Column {
	return &Column{name, STRING, colSize}
}

func (c *Column) Name() string {
	return c.name
}

func (c *Column) Type() ToyDBType {
	return c.type_
}

func (c *Column) SerializeValue(v *Value) []byte {
	switch c.Type() {
	case INT64:
		return util.SerializeInt64(v.vInt64)
	case STRING:
		return util.SerializeString(v.vString, int64(c.colSize))
	default:
		panic("SerializeValue")
	}
}

func (c *Column) DeserializeValue(b []byte) *Value {
	switch c.Type() {
	case INT64:
		return NewValueInt64(util.DeserializeInt64(b))
	case STRING:
		return NewValueString(util.DeserializeString(b, int64(c.colSize)))
	default:
		panic("DeserializeValue")
	}
}
