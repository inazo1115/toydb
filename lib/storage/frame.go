package storage

import (
	"github.com/inazo1115/toydb/lib/page"
)

type Frame struct {
	page     *page.DataPage
	pinCount int32
	dirty    bool
}

func NewFrame(page *page.DataPage) *Frame {
	return &Frame{page, 0, false}
}

func (f *Frame) Page() *page.DataPage {
	return f.page
}

func (f *Frame) SetPage(page *page.DataPage) {
	f.page = page
}

func (f *Frame) DeletePage() {
	f.page = nil
}

func (f *Frame) PinCount() int32 {
	return f.pinCount
}

func (f *Frame) IncPinCount() {
	f.pinCount++
}

func (f *Frame) DecPinCount() {
	f.pinCount--
}

func (f *Frame) Dirty() bool {
	return f.dirty
}

func (f *Frame) TurnOnDirty() {
	f.dirty = true
}

func (f *Frame) TurnOffDirty() {
	f.dirty = false
}
