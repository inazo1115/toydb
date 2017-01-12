package query

import (
	//"fmt"

	"github.com/inazo1115/toydb/lib/file"
	"github.com/inazo1115/toydb/lib/storage"
	"github.com/inazo1115/toydb/lib/table"
)

type Runtime struct {
	bm      *storage.BufferManager
	curTbl  string
	schemas map[string]*table.Schema  // tmp
	files   map[string]*file.HeapFile // tmp
}

func NewRuntime() *Runtime {
	bm := storage.NewBufferManager()
	return &Runtime{bm, "", make(map[string]*table.Schema),
		make(map[string]*file.HeapFile)}
}

func (r *Runtime) BufferManager() *storage.BufferManager {
	return r.bm
}

func (r *Runtime) CurrentTableName() string {
	return r.curTbl
}

func (r *Runtime) SetCurrentTableName(s string) {
	r.curTbl = s
}

func (r *Runtime) CurrentSchema() *table.Schema {
	return r.schemas[r.curTbl]
}

func (r *Runtime) SetSchema(name string, schema *table.Schema) {
	r.schemas[name] = schema
}

func (r *Runtime) CurrentFile() *file.HeapFile {
	return r.files[r.curTbl]
}

func (r *Runtime) SetFile(name string, file *file.HeapFile) {
	r.files[name] = file
}
