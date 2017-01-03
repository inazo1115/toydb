package storage

import (
	"math"
)

// LRUStrategy is the cache eviction strategy which uses pseudo LRU algorithm.
type LRUStrategy struct {
}

// NewLRUStrategy return the pointer of LRUStrategy.
func NewLRUStrategy() *LRUStrategy {
	return &LRUStrategy{}
}

// TouchPage is the process which is happened when page use.
func (s *LRUStrategy) TouchPage(bm *BufferManager, pid int64) {

	// Increment ages.
	for _, frame := range bm.bufferPool {
		frame.IncHitCount()
	}

	// The page use resets count.
	bm.bufferPool[bm.dict[pid]].SetHitCount(0)
}

// ChooseVictim selects the eviction target.
func (s *LRUStrategy) ChooseVictim(bm *BufferManager) int64 {

	// Choose the frame which stores the oldest page.
	max := int64(-1)
	min := int64(math.MaxInt64)
	target := int64(-1)
	for fidx, frame := range bm.bufferPool {
		if max < frame.HitCount() {
			max = frame.HitCount()
			target = int64(fidx)
		} else if min > frame.HitCount() {
			min = frame.HitCount()
		}
	}

	// Align minimum count to zero.
	for _, frame := range bm.bufferPool {
		frame.SetHitCount(frame.HitCount() - min)
	}

	return target
}
