package priorityqueue

import "sync"

import "github.com/BTooLs/data-structures/heap"
import "errors"

/*HierarchicalHeap It is a modification of the Hierarchical Queue structure, adding some complexity (O(log n/k)) but removing it's limitations.
Inspired by [Cris L. Luengo Hendriks paper](http://www.cb.uu.se/~cris/Documents/Luengo2010a_preprint.pdf)

Unlike HQ that has 1 bucket for each priority value, HHeap priorities are grouped (using a linear mapping formula) in K buckets.
Buckets are Implicit heaps (special binary tree) that stores all values for a specific range of priorities.
Example: pMin=0, pMax=100 (0-100 priorities), K=15 (buckets)
Enqueue("a", 21) will add "a" to bucket i, where
i = (21 − pmin) / (pmax − pmin) * K = 3

For the best performance benchmark the Enq/Deq functions with your own data, and tweak the buckets,minP,maxP parameters!
.*/
type HierarchicalHeap struct {
	autoLock bool
	sync.Mutex
	pMin    int
	pMax    int
	buckets int
	count   int
	a       []*heap.ImplicitHeapMin //buckets of heaps
	smB     int                     //smallest bucket index that has values, cached for optimization
}

//NewHierarchicalHeap Generates a new HQ. 0 priority = max
func NewHierarchicalHeap(buckets, pMin, pMax int, autoMutexLock bool) (*HierarchicalHeap, error) {
	if buckets < 1 {
		return nil, errors.New("must have at least 1 bucket")
	}

	if pMin < 0 || pMin > pMax {
		return nil, errors.New("pMin must be positive and smaller than pMax")
	}

	return &HierarchicalHeap{
		autoLock: autoMutexLock,
		pMin:     pMin,
		pMax:     pMax,
		buckets:  buckets, //cache
		a:        make([]*heap.ImplicitHeapMin, buckets),
	}, nil
}

func (h *HierarchicalHeap) getBucket(priority int) (*heap.ImplicitHeapMin, int) {
	//not ideal and performant, but we can deal with higher/lower priorities
	if priority > h.pMax {
		priority = h.pMax
	}

	if priority < h.pMin {
		priority = h.pMin
	}

	i := int((float64(priority) - float64(h.pMin)) / float64(h.pMax-h.pMin) * float64(h.buckets))

	//if the buckets are not set correctly, we must fix them
	if i < 0 || i > h.buckets-1 {
		i = h.buckets - 1
	}

	//lazy init
	if h.a[i] == nil {
		h.a[i] = heap.NewImplicitHeapMin(false)
	}

	return h.a[i], i
}

//Enqueue add a new key/value pair in the queue. 0 priority = max. O(log n/k)
func (h *HierarchicalHeap) Enqueue(value interface{}, priority int) error {
	if h.autoLock {
		h.Lock()
		defer h.Unlock()
	}

	bucket, bucketIndex := h.getBucket(priority)

	if bucketIndex < h.smB {
		h.smB = bucketIndex
	}

	// if bucket == nil {
	// 	return errors.New("cannot get bucket for " + string(priority))
	// }
	bucket.Push(priority, value)
	h.count++
	return nil
}

//Dequeue Remove and return the highest key (lowest priority) O(log n/k)
func (h *HierarchicalHeap) Dequeue() (v interface{}, err error) {
	if h.autoLock {
		h.Lock()
		defer h.Unlock()
	}

	for i := h.smB; i < h.buckets; i++ {
		if h.a[i] == nil || h.a[i].IsDepleted() {
			continue
		}
		h.smB = i
		val, _ := h.a[i].Pop()

		//can't replicate this scenario
		// if ok == false {
		// 	err = errors.New("queue pop failed")
		// 	return
		// }

		v = val
		h.count--
		return
	}

	err = errors.New("the queue is empty")
	return
}

//Len Return the count of all values from all priorities. O(1)
func (h *HierarchicalHeap) Len() int {
	if h.autoLock {
		h.Lock()
		defer h.Unlock()
	}

	return h.count
}
