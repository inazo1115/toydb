package query

import (
	"github.com/inazo1115/toydb/lib/storage"
)

type Runtime struct {
	bufman  *storage.BufferManager
	diskman *storage.DiskManager
}

func NewRuntime(bufman *storage.BufferManager, diskman *storage.DiskManager) *Runtime {
	return &Runtime{bufman, diskman}
}

func (r *Runtime) BufferManager() *storage.BufferManager {
	return r.bufman
}

func (r *Runtime) DiskManager() *storage.DiskManager {
	return r.diskman
}
