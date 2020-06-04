package util

// Int64Arr is the sort interface of int64.
type Int64Arr []int64

func (a Int64Arr) Len() int           { return len(a) }
func (a Int64Arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64Arr) Less(i, j int) bool { return a[i] < a[j] }
