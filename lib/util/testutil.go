package util

import (
	"testing"
)

func Assert(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("actual: %v doesn't equals expected: %v.", actual, expected)
	}
}
