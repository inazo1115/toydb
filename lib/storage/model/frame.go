package model

type Frame struct {
	page     *Page
	pinCount int
	dirty    bool
}

func NewFrame(p *Page) *Frame {
	return &Frame{p, 0, false}
}

func (f *Frame) Page() *Page {
	return f.page
}

func (f *Frame) SetPage(p *Page) {
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
