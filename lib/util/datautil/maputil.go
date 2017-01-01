package datautil

// TODO: use generics??
func Keys(m map[int]int) []int {
	ret := make([]int, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}
