package tree

//ImplicitHeap A dynamic tree (list) of numbers, stored as a Binary tree in a slice.
type ImplicitHeap interface {
	Push(v int)
	Pop() (v int, ok bool)
	Peek() (v int, ok bool)
	Reset()
}

//ImplicitHeapMin A dynamic tree (list) of numbers, stored as a Binary tree in a slice.
//Used to quickly get the smallest number from a list/queue/priority queue.
type ImplicitHeapMin struct {
	a       []int
	n       int //numbers in the heap
	compare ihCompare
}

//shouldGoUp We keep the min comparasion formula in 1 place
//it is overwritten for Max
func minShouldGoUp(p, c int) bool {
	return c < p
}

//inheritance bypass, the overloading didn't worked :(
//TODO learn how to do a better composition (Parent calls a func from child)
type ihCompare func(p, c int) bool

func (h *ImplicitHeapMin) lazyInit() {
	if h.a == nil {
		h.a = make([]int, 8)
	}

	if h.compare == nil {
		h.compare = minShouldGoUp
	}
}

//Push Insert a new number in the list.
func (h *ImplicitHeapMin) Push(v int) {
	h.lazyInit()

	//if it is full, enlarge it
	if cap(h.a) == h.n {
		newSlice := make([]int, cap(h.a)*2)
		copy(newSlice, h.a)
		h.a = newSlice
	}

	h.a[h.n] = v
	h.n++

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
	for h.compare(h.a[pI], h.a[cI]) && pI >= 0 {
		h.a[pI], h.a[cI] = h.a[cI], h.a[pI]
		cI = pI
		pI = (cI - 1) / 2
	}
}

func (h *ImplicitHeapMin) shouldGoDownAt(pI, lcI, rcI int) int {
	leftCExists := lcI < h.n
	rightCExists := rcI < h.n

	if leftCExists == false && rightCExists == false {
		return -1 //the parent is a leaf
	}

	lcViable := leftCExists && h.compare(h.a[pI], h.a[lcI])
	rcViable := rightCExists && h.compare(h.a[pI], h.a[rcI])

	if lcViable == false && rcViable == false {
		return -1
	}

	if lcViable && rcViable {
		if h.a[lcI] > h.a[rcI] {
			return rcI
		}
		return lcI
	}

	if lcViable {
		return lcI
	}

	return rcI
}

//Peek Find-Min returns the minimum value (root element) O(1)
//Does not mutate the list
func (h *ImplicitHeapMin) Peek() (v int, ok bool) {
	h.lazyInit()

	if h.n <= 0 {
		return 0, false
	}

	return h.a[0], true
}

//Pop Delete-Min, return the minimum value (root element) O(log(n))
//Removes the element from the list
func (h *ImplicitHeapMin) Pop() (v int, ok bool) {
	h.lazyInit()

	if h.n <= 0 {
		return 0, false
	}

	//pop the root, exchange it with the last leaf
	v = h.a[0]
	ok = true

	h.a[0] = 0
	h.n--

	if h.n == 0 {
		return //no elements left, nothing to sort
	}

	h.a[0], h.a[h.n] = h.a[h.n], 0

	//parentI - the poped root
	for pI, switchToI, lcI, rcI := 0, 0, 0, 0; ; {
		lcI = 2*pI + 1 //left child index
		rcI = lcI + 1
		switchToI = h.shouldGoDownAt(pI, lcI, rcI)

		if switchToI < 0 {
			break
		}
		h.a[pI], h.a[switchToI] = h.a[switchToI], h.a[pI]

		pI = switchToI
	}

	//if it is mostly empty (less than 1/4), shrink it
	if cap(h.a) > 8 && h.n <= cap(h.a)/4 {
		newSlice := make([]int, cap(h.a)/2)
		copy(newSlice, h.a)
		h.a = newSlice
	}

	return
}

//Reset Feed all your data to the Garbage Collector.
func (h *ImplicitHeapMin) Reset() {
	h.a = make([]int, 8)
	h.n = 0
}
