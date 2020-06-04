package util

import (
	"testing"
)

func TestCopyStrToByte(t *testing.T) {

	bufSize := 10
	buf := make([]byte, bufSize)
	s := "test"
	expected := []byte{116, 101, 115, 116, 0, 0, 0, 0, 0, 0}

	if err := CopyStrToByte(s, buf); err != nil {
		t.Errorf("CopyStrToByte failed.")
	}

	for i, b := range expected {
		if buf[i] != b {
			t.Errorf("actual:%s doesn't equal expected:%s.", string(buf), s)
		}
	}
}
