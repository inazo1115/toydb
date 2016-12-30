package buffer

import (
	"fmt"

	"github.com/inazo1115/toydb/lib/toydb/file"
)

type Frame struct {
	page Page
	pin_count int
	dirty bool
}

func NewFrame(page Page) Frame {
	return Frame(page, 0, false)
}

func (frame Frame) GetPage() Page {
	return frame.page
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
