package model

//import (
//	"fmt"
//)

type Frame struct {
	page *Page
	pin_count int
	dirty bool
}

func NewFrame(page *Page) Frame {
	return Frame{page, 0, false}
}

func (frame Frame) Page() *Page {
	return frame.page
}

func (frame Frame) SetPage(page *Page) {
	frame.page = page
}

func (frame Frame) PinCount() int {
	return frame.pin_count
}

func (frame Frame) Dirty() bool {
	return frame.dirty
}

func (frame *Frame) IncPinCount() {
	frame.pin_count += 1
}

func (frame *Frame) DecPinCount() {
	frame.pin_count -= 1
}

func (frame *Frame) TurnOnDirty() {
	frame.dirty = true
}

func (frame *Frame) TurnOffDirty() {
	frame.dirty = false
}
