package tree

import (
	"fmt"
)

func ExampleImplicitHeapMin() {
	autoLockMutex := false
	minPriority := NewImplicitHeapMin(autoLockMutex)
	//add 3 priorities, with any data types
	minPriority.Push(3, false)
	minPriority.Push(1, "alfa")
	minPriority.Push(9, []int{0, 0})

	min, _ := minPriority.Pop()
	fmt.Printf("min is %v", min)
	// Output: min is alfa
}

func ExampleImplicitHeapMax() {
	autoLockMutex := false
	maxPriority := NewImplicitHeapMax(autoLockMutex)
	//add 3 priorities, with any data types
	maxPriority.Push(3, false)
	maxPriority.Push(1, "alfa")
	maxPriority.Push(9, []int{0, 0})

	max, _ := maxPriority.Pop()
	fmt.Printf("max is %v", max)
	// Output: max is [0 0]
}
