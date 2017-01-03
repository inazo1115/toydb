package util

import (
	"sort"
	"testing"
)

func TestKeys(t *testing.T) {

	m := map[int64]int64{
		0: 10,
		1: 11,
		2: 12,
	}

	actual := Keys(m)
	sort.Sort(Int64Arr(actual))
	expected := []int64{0, 1, 2}

	for i := 0; i < 3; i++ {
		if actual[i] != expected[i] {
			t.Errorf("actual:%v doesn't equal expected:%v.", actual, expected)
		}
	}
}
