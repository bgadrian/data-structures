package tree

import (
	"fmt"
)

func todo() {
	minPriority := ImplicitHeapMin{}
	minPriority.Push(3)
	minPriority.Push(1)
	minPriority.Push(9)

	min, ok := minPriority.Peek()
	if ok == false {
		fmt.Print("something went wrong or is empty")
		return
	}

	fmt.Printf("min is %v", min) //min is 1

}
