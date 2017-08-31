package tree

//MaxImplicitHeap A dynamic list of numbers, stored as a Binary tree in a slice.
//Used to quickly get the biggest number from a list/queue/priority queue.
type MaxImplicitHeap struct {
	MinImplicitHeap
}

func (h *MaxImplicitHeap) shouldGoUp(p, c int) bool {
	return c > p
}
