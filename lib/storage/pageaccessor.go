package storage

import (
	"errors"

	"github.com/inazo1115/toydb/lib/storage/model"
)

// PageAccessor is a facade of storage. When you want to get page or modify
// page or do something else for page, use this module. PageAccessor provides
// simple CRUD methods and some utilities.
type PageAccessor struct {
	bm *BufferManager
}

// NewPageAccessor returns a pointer of PageAccessor.
func NewPageAccessor() *PageAccessor {
	return &PageAccessor{NewBufferManager()}
}

// Create creates a new page which has contents given and returns a new page id.
func (pa *PageAccessor) Create(page *model.Page) (int, error) {
	return pa.bm.Create(page)
}

// Read returns a page which has given page id.
func (pa *PageAccessor) Read(pid int) (*model.Page, error) {
	return pa.bm.Read(pid)
}

// Update updates a page which has given page id to a new page.
func (pa *PageAccessor) Update(pid int, page *model.Page) error {
	return pa.bm.Update(pid, page)
}

// Delete deletes a page which has page id given through an argument.
func (pa *PageAccessor) Delete(pid int) error {
	return errors.New("not implemented")
}

// WriteBackAll writes back all pages on main memory to disk.
func (pa *PageAccessor) WriteBackAll() error {
	return pa.bm.WriteBackAll()
}

// Dump shows the inner information. it's for debug.
func (pa *PageAccessor) Dump() {
	pa.bm.Dump()
}
