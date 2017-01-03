package storage

import (
	"testing"

	"github.com/inazo1115/toydb/lib/page"
)

func TestFramePage(t *testing.T) {
	p0 := page.NewDataPage(-1, -1, -1)
	f := NewFrame(p0)
	if f.Page() != p0 {
		t.Errorf("NewFrame failed")
	}
	p1 := page.NewDataPage(-1, -1, -1)
	f.SetPage(p1)
	if f.Page() != p1 {
		t.Errorf("SetFrame failed")
	}
	f.DeletePage()
	if f.Page() != nil {
		t.Errorf("DeleteFrame failed")
	}
}

func TestFramePin(t *testing.T) {
	f := NewFrame(nil)
	if f.PinCount() != 0 {
		t.Errorf("The initial pin counts should be 0.")
	}
	f.Pin()
	if f.PinCount() != 1 {
		t.Errorf("This should be 1.")
	}
	f.UnPin()
	if f.PinCount() != 0 {
		t.Errorf("This should be 0.")
	}
}

func TestFrameDirty(t *testing.T) {
	f := NewFrame(nil)
	if f.Dirty() != false {
		t.Errorf("The initial dirty should be false.")
	}
	f.TurnOnDirty()
	if f.Dirty() != true {
		t.Errorf("This should be true.")
	}
	f.TurnOffDirty()
	if f.Dirty() != false {
		t.Errorf("This should be false.")
	}
}

func TestFrameHitCount(t *testing.T) {
	f := NewFrame(nil)
	if f.HitCount() != 0 {
		t.Errorf("The initial hit counts should be 0.")
	}
	f.IncHitCount()
	if f.HitCount() != 1 {
		t.Errorf("This should be 1.")
	}
	f.SetHitCount(0)
	if f.HitCount() != 0 {
		t.Errorf("This should be 0.")
	}
}
