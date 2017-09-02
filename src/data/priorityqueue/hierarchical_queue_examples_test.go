package priorityqueue

import (
	"fmt"
)

func ExampleHierarchicalQueue() {
	autoLockMutex := false
	var lowestPriority uint8 = 10 //highest is 0
	l := NewHierarchicalQueue(lowestPriority, autoLockMutex)

	l.Enqueue("a", 2)
	l.Enqueue("b", 0)

	first, _ := l.Dequeue()
	fmt.Printf("first is %v", first)
	// Output: first is b
}
