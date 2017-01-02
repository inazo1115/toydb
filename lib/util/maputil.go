package util

// TODO: use generics??
func Keys(m map[int64]int64) []int64 {
	ret := make([]int64, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}
