package tree

//ImplicitHeapMax A dynamic list of numbers, stored as a Binary tree in a slice.
//Used to quickly get the biggest number from a list/queue/priority queue.
type ImplicitHeapMax struct {
	ImplicitHeapMin
}

func maxShouldGoUp(p, c int) bool {
	return c > p
}

//Push Push a new number in the list.
func (h *ImplicitHeapMax) Push(v int) {
	if h.compare == nil {
		h.compare = maxShouldGoUp
	}

	h.ImplicitHeapMin.Push(v)
}

//Pop Delete-Max, return the maximum value (root element) O(log(n))
//Removes the element from the list
func (h *ImplicitHeapMax) Pop() (v int, ok bool) {
	if h.compare == nil {
		h.compare = maxShouldGoUp
	}

	return h.ImplicitHeapMin.Pop()
}
