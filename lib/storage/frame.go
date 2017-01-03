package storage

import (
	"github.com/inazo1115/toydb/lib/page"
)

// Frame is the model of buffer which is on the memory. Frame stores a page and
// some information of the cache control.
type Frame struct {
	page     *page.DataPage
	pinCount int64
	dirty    bool
	hitCount int64 // hitCount is used by cache control strategy.
}

// NewFrame creates a new frame and returns it's pointer.
func NewFrame(page *page.DataPage) *Frame {
	return &Frame{page, 0, false, 0}
}

// Page is the getter of the page.
func (f *Frame) Page() *page.DataPage {
	return f.page
}

// SetPage is the setter of the page.
func (f *Frame) SetPage(page *page.DataPage) {
	f.page = page
}

// DeletePage deletes the reference of the page which had been stored.
func (f *Frame) DeletePage() {
	f.page = nil
}

// PinCount is the setter of the pinCount.
func (f *Frame) PinCount() int64 {
	return f.pinCount
}

// Pin increments pin counts.
func (f *Frame) Pin() {
	f.pinCount++
}

// Pin decrements pin counts.
func (f *Frame) UnPin() {
	f.pinCount--
}

// Dirty is the getter of the dirty. If this parameter is true, the page has
// been modified since load from disk.
func (f *Frame) Dirty() bool {
	return f.dirty
}

// TurnOnDirty sets the dirty true.
func (f *Frame) TurnOnDirty() {
	f.dirty = true
}

// TurnOnDirty sets the dirty false.
func (f *Frame) TurnOffDirty() {
	f.dirty = false
}

// HitCount is the getter of the hitCount.
func (f *Frame) HitCount() int64 {
	return f.hitCount
}

// IncHitCount increments the hit counts. This represents the cache hit.
func (f *Frame) IncHitCount() {
	f.hitCount++
}

// SetHitCount is the setter of the hitCount.
func (f *Frame) SetHitCount(hitCount int64) {
	f.hitCount = hitCount
}
