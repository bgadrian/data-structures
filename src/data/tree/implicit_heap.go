package tree

/*ImplicitHeap is an interface for dynamic min & max implicit heaps.

An implicit heap is an implementation of a heap consisting of a complete binary tree whose
nodes contain the heap items, one node per item.

It can be used as a Priority queue as well.
It can store a series of objects (of any type) associated with a key (priority).
Use ImplicitHeapMin to always get the smallest key (priority)
and ImplicitHeapMax for the largest key.

For best perfomance use a small non sparsed Key value distribution. (100-300 incremental values).
*/
type ImplicitHeap interface {
	Push(priority int, v interface{})
	Pop() (v interface{}, ok bool)
	Peek() (v interface{}, ok bool)
	Reset()
	Lock()
	Unlock()
	IsDepleted() bool
	HasElement() bool
}

//inheritance bypass, the overloading didn't worked :(
//TODO learn how to do a better composition (Parent calls a func from child)
type ihCompare func(p, c implicitHeapNode) bool

//implicitHeapNode Elements of the Heap.
//No much use of heaps just with numbers.
//We usually use them to store ... stuff.
type implicitHeapNode struct {
	priority int
	value    interface{}
}
