package util

// Keys get keys from given map. Target type is int64.
func Keys(m map[int64]int64) []int64 {
	ret := make([]int64, len(m))
	i := 0
	for k := range m {
		ret[i] = k
		i++
	}
	return ret
}
