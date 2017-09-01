package tree

//ImplicitHeapMax A dynamic list of numbers, stored as a Binary tree in a slice.
//Used to quickly get the biggest number from a list/queue/priority queue.
type ImplicitHeapMax struct {
	ImplicitHeapMin
}

func maxShouldGoUp(p, c implicitHeapNode) bool {
	return c.priority > p.priority
}

//NewImplicitHeapMax Constructor for IH Max
func NewImplicitHeapMax(autoLockMutex bool) *ImplicitHeapMax {
	h := &ImplicitHeapMax{}
	h.compare = maxShouldGoUp
	h.autoLockMutex = autoLockMutex
	h.Reset()
	return h
}
