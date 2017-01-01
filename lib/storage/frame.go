package storage

import (
	"github.com/inazo1115/toydb/lib/page"
)

type Frame struct {
	page     *page.DataPage
	pinCount int
	dirty    bool
}

func NewFrame(p *page.DataPage) *Frame {
	return &Frame{p, 0, false}
}

func (f *Frame) Page() *page.DataPage {
	return f.page
}

func (f *Frame) SetPage(p *page.DataPage) {
	f.page = p
}

func (f *Frame) DeletePage() {
	f.page = nil
}

func (f *Frame) PinCount() int {
	return f.pinCount
}

func (f *Frame) IncPinCount() {
	f.pinCount += 1
}

func (f *Frame) DecPinCount() {
	f.pinCount -= 1
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
