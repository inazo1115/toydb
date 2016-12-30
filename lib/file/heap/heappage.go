package file

import (
	"fmt"
)

type Heap struct {
	page Page
	prev_pid int
	next_pid int
}

func NewHeap(page Page, prev_pid int, next_pid int) {
	return Heap{page, prev_pid, next_pid}
}
