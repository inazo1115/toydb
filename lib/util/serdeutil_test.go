package util

import (
	"testing"
)

func TestSerdeInt64(t *testing.T) {
	expected := int64(1234)
	actual := DeserializeInt64(SerializeInt64(expected))
	if actual != expected {
		t.Errorf("actual: %d doesn't equals expected: %d.", actual, expected)
	}
}

func TestSerdeString(t *testing.T) {
	expected := "test"

	ser := SerializeString(expected, int64(len(expected)))
	de := DeserializeString(ser, int64(len(expected)))

	actual := de
	if actual != expected {
		t.Errorf("actual: %s doesn't equals expected: %s.", actual, expected)
	}
}
