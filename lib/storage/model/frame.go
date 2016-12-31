package model

type Frame struct {
	page     *Page
	pinCount int
	dirty    bool
}

func NewFrame(page *Page) *Frame {
	return &Frame{page, 0, false}
}

func (frame *Frame) Page() *Page {
	return frame.page
}

func (frame *Frame) SetPage(page *Page) {
	frame.page = page
}

func (frame *Frame) PinCount() int {
	return frame.pinCount
}

func (frame *Frame) IncPinCount() {
	frame.pinCount += 1
}

func (frame *Frame) DecPinCount() {
	frame.pinCount -= 1
}

func (frame *Frame) Dirty() bool {
	return frame.dirty
}

func (frame *Frame) TurnOnDirty() {
	frame.dirty = true
}

func (frame *Frame) TurnOffDirty() {
	frame.dirty = false
}
