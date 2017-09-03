package heap

import "sync"

//ImplicitHeapMin A dynamic tree (list) of numbers, stored as a Binary tree in a dynamic slice.
//Used to quickly get the smallest number from a list/queue/priority queue.
//
//It is a base struct for ImplicitHeapMax.
type ImplicitHeapMin struct {
	a             []implicitHeapNode
	n             int       //numbers in the heap
	compare       ihCompare //different compare func for Min/Max
	autoLockMutex bool      //auto locks the mutex for each func call
	sync.Mutex
}

//shouldGoUp We keep the min comparasion formula in 1 place
//it is overwritten for Max
func minShouldGoUp(p, c implicitHeapNode) bool {
	return c.priority < p.priority
}

//NewImplicitHeapMin Builds an empty ImplicitHeapMin
func NewImplicitHeapMin(autoLockMutex bool) *ImplicitHeapMin {
	h := &ImplicitHeapMin{
		compare:       minShouldGoUp,
		autoLockMutex: autoLockMutex}
	h.Reset()
	return h
}

//Push Insert a new key/value pair in the list.
func (h *ImplicitHeapMin) Push(priority int, value interface{}) {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	//if it is full, enlarge it
	if cap(h.a) == h.n {
		newSlice := make([]implicitHeapNode, cap(h.a)*2)
		copy(newSlice, h.a)
		h.a = newSlice
	}

	h.a[h.n] = implicitHeapNode{priority, value}
	h.n++

	if h.n <= 1 {
		return //nothing to sort
	}

	//rebalance the tree, check the new value parents
	/*
		parentIndex = (childIndex - 1 ) / 2
		[0,1,2,3,4,5,6,7]
			0 = root node
			1 = left child; 0 = (1-1) / 2
			2 = right child; 0 = (2-1) / 2
			3 = left child of 1 ; 1 = (3-1) / 2
			4 = right child of 1 ; 1 = (4-1) / 2
			5 = left child of 2 ; 2 = (5 - 1) / 2
			6 = right child of 2 ; 2 = (6 - 1) / 2
	*/
	cI := h.n - 1      //childIndex, newest number
	pI := (cI - 1) / 2 //parentIndex
	for cI > 0 && h.compare(h.a[pI], h.a[cI]) {
		h.a[pI], h.a[cI] = h.a[cI], h.a[pI]
		cI = pI
		pI = (cI - 1) / 2
	}
}

//Peek Find-* returns the first value (root element) O(1).
//For ImplicitHeapMin returns the value for the smallest key(priority).
//For ImplicitHeapMax returns the value for the largest key(priority).
//Does not mutate the list
func (h *ImplicitHeapMin) Peek() (v interface{}, ok bool) {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	if h.n <= 0 {
		return 0, false
	}

	return h.a[0].value, true
}

//Pop Delete-*, return the first value (root element) O(log(n))
//For ImplicitHeapMin returns the value for the smallest key(priority).
//For ImplicitHeapMax returns the value for the largest key(priority).
//Removes the element from the list
func (h *ImplicitHeapMin) Pop() (v interface{}, ok bool) {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	if h.n <= 0 {
		return
	}

	//pop the root, exchange it with the last leaf
	v = h.a[0].value
	ok = true

	h.n--

	h.a[0] = h.a[h.n]

	//mark it as delete, for testing purposes
	h.a[h.n].priority = 0
	h.a[h.n].value = nil

	if h.n <= 1 {
		return //no use to sort
	}

	pI, isLc, isRc, leftChildIndex, rightChildIndex := 0, false, false, 0, 0

	for {
		leftChildIndex = 2*pI + 1
		rightChildIndex = leftChildIndex + 1

		//should the parent switch to left chid?
		isLc = leftChildIndex < h.n && h.compare(h.a[pI], h.a[leftChildIndex])
		isRc = rightChildIndex < h.n && h.compare(h.a[pI], h.a[rightChildIndex])

		if isLc == false && isRc == false {
			break
		}

		if isLc && isRc {
			if h.compare(h.a[leftChildIndex], h.a[rightChildIndex]) {
				isLc = false
			}
			isRc = false
		}

		if isLc {
			h.a[pI], h.a[leftChildIndex] = h.a[leftChildIndex], h.a[pI]
			pI = leftChildIndex
			continue
		}

		//isRC
		h.a[pI], h.a[rightChildIndex] = h.a[rightChildIndex], h.a[pI]
		pI = rightChildIndex
	}

	//if it is mostly empty (less than 1/4), shrink it
	if cap(h.a) > 8 && h.n <= cap(h.a)/4 {
		newSlice := make([]implicitHeapNode, cap(h.a)/2)
		copy(newSlice, h.a)
		h.a = newSlice
	}

	return
}

//Reset Feed all your data to the Garbage Collector.
func (h *ImplicitHeapMin) Reset() {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	h.a = make([]implicitHeapNode, 8)
	h.n = 0
}

//IsDepleted Check if the list is empty
func (h *ImplicitHeapMin) IsDepleted() bool {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	return h.n == 0
}

//HasElement Check if the list has at least 1 elment left
func (h *ImplicitHeapMin) HasElement() bool {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	return h.n > 0
}

//Len How many elements are in the heap
func (h *ImplicitHeapMin) Len() int {
	if h.autoLockMutex {
		h.Lock()
		defer h.Unlock()
	}

	return h.n
}
