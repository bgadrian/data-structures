package tree

//ImplicitHeap A dynamic tree (list) of numbers, stored as a Binary tree in a slice.
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
