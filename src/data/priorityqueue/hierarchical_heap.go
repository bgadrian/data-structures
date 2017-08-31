package priorityqueue

import "sync"

/*HierarchicalHeap An O(logN/k) if the buckets are chosen correctly, priority queue
Is similar with the Hierchical Queue, it removes it's limitation but adds some complexity.
Unlike HQ that has 1 bucket for each priority value, HHeap priorities are grouped (using a linear mapping formula) in K buckets.
Buckets are Implicit heaps (special binary tree) that stores all values for a specific range of priorities.
Example: pMin=0, pMax=100 (0-100 priorities), K=15 (buckets)
Enqueue("a", 21) will add "a" to bucket i, where
i = (2 − pmin / pmax − pmin) * K = 3
.*/
type HierarchicalHeap struct {
	autoLock bool
	sync.Mutex
	count int
	pMin  int
	pMax  int
	// b     []*ImplicitHeap
}

//NewHierarchicalHeap Generates a new HQ
func NewHierarchicalHeap(buckets int, autoMutexLock bool) *HierarchicalHeap {
	return &HierarchicalHeap{
		autoLock: autoMutexLock,
	}
}
