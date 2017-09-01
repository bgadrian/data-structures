package tree

import (
	"fmt"
)

func todo() {
	minPriority := ImplicitHeapMin{}
	minPriority.Push(3, "a")
	minPriority.Push(1, "b")
	minPriority.Push(9, "c")

	min, ok := minPriority.Peek()
	if ok == false {
		fmt.Print("something went wrong or is empty")
		return
	}

	fmt.Printf("min is %v", min) //min is "b"

}
