package tree

import (
	"reflect"
	"testing"
)

//most of the functionalities are common and tested in min_*test

func TestMaxIHAddOrder1(t *testing.T) {
	h := MaxImplicitHeap{}

	h.Push(1)
	h.Push(3)
	h.Push(4)

	if reflect.DeepEqual(h.a[:3], []int{4, 3, 1}) == false {
		t.Error("push was not ok")
	}

	h.Push(2)

	if reflect.DeepEqual(h.a[:4], []int{4, 3, 1, 2}) == false {
		t.Error("push was not ok 2")
	}

	h.Push(5)

	if reflect.DeepEqual(h.a[:5], []int{5, 4, 1, 2, 3}) == false {
		t.Error("push was not ok 3")
	}

}
