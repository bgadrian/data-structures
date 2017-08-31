package tree

//ImplicitHeap A dynamic list of numbers, stored as a Binary tree in a slice.
type ImplicitHeap interface {
	Push(v int)
	Pop() (v int, ok bool)
}

//MinImplicitHeap A dynamic list of numbers, stored as a Binary tree in a slice.
//Used to quickly get the smallest number from a list/queue/priority queue.
type MinImplicitHeap struct {
	a []int
	n int //numbers in the heap
}

func (h *MinImplicitHeap) lazyInit() {
	if h.a == nil {
		h.a = make([]int, 8)
	}
}

//Push Insert a new number in the list.
func (h *MinImplicitHeap) Push(v int) {
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
	for h.shouldGoUp(h.a[pI], h.a[cI]) && pI >= 0 {
		h.a[pI], h.a[cI] = h.a[cI], h.a[pI]
		cI = pI
		pI = (cI - 1) / 2
	}
}

func (h *MinImplicitHeap) shouldGoUp(p, c int) bool {
	return c < p
}

func (h *MinImplicitHeap) shouldGoDownAt(pI, lcI, rcI int) int {
	leftCExists := lcI < h.n
	rightCExists := rcI < h.n

	if leftCExists == false && rightCExists == false {
		return -1 //the parent is a leaf
	}

	lcViable := leftCExists && h.shouldGoUp(h.a[pI], h.a[lcI])
	rcViable := rightCExists && h.shouldGoUp(h.a[pI], h.a[rcI])

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

//Pop Pop the root element (min/max) O(log(n))
func (h *MinImplicitHeap) Pop() (v int, ok bool) {
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
func (h *MinImplicitHeap) Reset() {
	h.a = make([]int, 8)
	h.n = 0
}
