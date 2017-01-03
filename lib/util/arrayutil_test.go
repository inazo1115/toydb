package util

import (
	"sort"
	"testing"
)

func TestInt64Arr(t *testing.T) {

	actual := []int64{5, 2, 3, 7, 9, 8, 1, 0, 4, 6}
	sort.Sort(Int64Arr(actual))
	expected := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < 10; i++ {
		if actual[i] != expected[i] {
			t.Errorf("actual:%v doesn't equal expected:%v.", actual, expected)
		}
	}
}
