package tree

import (
	"reflect"
	"testing"
)

//most of the functionalities are common and tested in min_*test

func TestIHMaxAddOrder1(t *testing.T) {
	h := ImplicitHeapMax{}

	h.Push(1)
	h.Push(3)
	h.Push(4)

	if reflect.DeepEqual(h.a[:3], []int{4, 1, 3}) == false {
		t.Error("push was not ok")
	}

	h.Push(2)

	if reflect.DeepEqual(h.a[:4], []int{4, 2, 3, 1}) == false {
		t.Error("push was not ok 2")
	}

	h.Push(5)

	if reflect.DeepEqual(h.a[:5], []int{5, 4, 3, 1, 2}) == false {
		t.Error("push was not ok 3")
	}

}

func TestIHMaxPopFirst(t *testing.T) {
	h := ImplicitHeapMax{}

	_, ok := h.Pop()

	if ok == true {
		t.Error("pop didn't returned error when empty")
	}
}

func TestIHMaxDuplicates(t *testing.T) {

	table := []testIHTuple{
		{&ImplicitHeapMax{},
			[]int{1, 1, 3, 3, 3},
			[]int{3, 3, 3, 1, 1}},
		{&ImplicitHeapMax{},
			[]int{5, 1, 4, 3, 4, 1, 1},
			[]int{5, 4, 4, 3, 1, 1, 1}},
	}

	for i := 0; i < len(table); i++ {
		testIMPopOrder(table[i].h, table[i].toPush, table[i].shouldPop, t)
	}
}

func TestIHMaxLarge(t *testing.T) {

	a35 := []int{35, 35, 34, 34, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 13, 12, 11, 10, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 2, 2, 2, 2, 1}

	table := []testIHTuple{
		{&ImplicitHeapMax{},
			[]int{2, 3, 13, 7, 10, 11, 20, 4, 2, 14, 12, 2, 22, 10, 18, 34, 5, 24, 34, 25, 2, 35, 32, 35, 34, 23, 26, 28, 13, 16, 9, 8, 33, 27, 2, 6, 1, 29, 10, 21, 19, 15, 30, 31, 17},
			a35},
		{&ImplicitHeapMax{},
			[]int{1, 33, 4, 24, 12, 13, 18, 3, 35, 32, 27, 10, 13, 2, 2, 35, 34, 2, 17, 9, 10, 20, 29, 2, 8, 30, 21, 22, 26, 28, 25, 34, 7, 5, 23, 19, 15, 16, 2, 14, 34, 10, 6, 11, 31},
			a35},
		{&ImplicitHeapMax{},
			[]int{5, 34, 35, 10, 23, 24, 6, 15, 32, 29, 12, 31, 18, 2, 27, 13, 34, 25, 2, 10, 1, 2, 20, 22, 16, 9, 13, 30, 7, 10, 11, 33, 28, 4, 3, 2, 17, 2, 19, 14, 21, 34, 35, 8, 26},
			a35},
	}

	for i := 0; i < len(table); i++ {
		testIMPopOrder(table[i].h, table[i].toPush, table[i].shouldPop, t)
	}
}
