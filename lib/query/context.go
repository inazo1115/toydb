package query

import (
	//"fmt"

	"github.com/inazo1115/toydb/lib/file"
	"github.com/inazo1115/toydb/lib/table"
)

type Context struct {
	currentTable string
	schemas      map[string]*table.Schema  // tmp
	files        map[string]*file.HeapFile // tmp
}

func NewContext() *Context {
	schemas := make(map[string]*table.Schema)
	files := make(map[string]*file.HeapFile)
	return &Context{"", schemas, files}
}

func (c *Context) CurrentTable() string {
	return c.currentTable
}

func (c *Context) SetCurrentTable(name string) {
	c.currentTable = name
}

func (c *Context) CurrentSchema() *table.Schema {
	return c.schemas[c.currentTable]
}

func (c *Context) SetSchema(name string, schema *table.Schema) {
	c.schemas[name] = schema
}

func (c *Context) CurrentFile() *file.HeapFile {
	return c.files[c.currentTable]
}

func (c *Context) SetFile(name string, file *file.HeapFile) {
	c.files[name] = file
}
