package tree

import (
	"reflect"
	"testing"
)

func TestMinIHBasic(t *testing.T) {
	h := MinImplicitHeap{}

	h.Push(6)
	v, ok := h.Pop()
	if ok == false {
		t.Error("cannot pop")
	}
	quickAssert(6, v, "basic pop 1", t)

	v, ok = h.Pop()
	quickAssertBool(false, ok, "pop empty was ok", t)
}

func TestMinIHAddOrder1(t *testing.T) {
	h := MinImplicitHeap{}

	h.Push(3)
	h.Push(1)
	h.Push(4)

	if reflect.DeepEqual(h.a[:3], []int{1, 3, 4}) == false {
		t.Error("push was not ok")
	}

	h.Push(2)

	if reflect.DeepEqual(h.a[:4], []int{1, 2, 4, 3}) == false {
		t.Error("push was not ok 2")
	}

	h.Push(0)

	if reflect.DeepEqual(h.a[:5], []int{0, 1, 4, 3, 2}) == false {
		t.Error("push was not ok 3")
	}
	h.Push(-1)

	if reflect.DeepEqual(h.a[:6], []int{-1, 1, 0, 3, 2, 4}) == false {
		t.Error("push was not ok 4")
	}
}

func TestMinIHPopOrder1(t *testing.T) {
	h := MinImplicitHeap{}

	h.Push(5)
	h.Push(3)
	h.Push(1)

	v, ok := h.Pop()
	quickAssertBool(true, ok, "pop not ok 1", t)
	quickAssert(1, v, "pop value 1", t)

	v, ok = h.Pop()
	quickAssertBool(true, ok, "pop not ok 2", t)
	quickAssert(3, v, "pop value 2", t)

	v, ok = h.Pop()
	quickAssertBool(true, ok, "pop not ok 3", t)
	quickAssert(5, v, "pop value 3", t)
}

func TestMinIHCapacity(t *testing.T) {
	h := MinImplicitHeap{}

	h.Push(101)
	quickAssert(8, cap(h.a), "capacity is default", t)
	quickAssert(h.a[0], 101, "root is set", t)

	addIHNodes(&h, 8)
	quickAssert(16, cap(h.a), "capacity doubled", t)

	addIHNodes(&h, 17)
	quickAssert(32, cap(h.a), "capacity doubled 2", t)
	quickAssert(26, h.n, "count is 26", t)

	popIHNodes(&h, 19, t)
	quickAssert(16, cap(h.a), "capacity shrinked /2", t)

	popIHNodes(&h, h.n, t)
	quickAssert(8, cap(h.a), "capacity never drop 8", t)

	// quickAssert(0, h.Levels(), "levels() after init", t)
}

func TestMinIHReset(t *testing.T) {
	h := MinImplicitHeap{}

	h.Push(1)
	h.Push(2)

	h.Reset()
	quickAssert(0, h.a[0], "reset forgot about elements", t)
	quickAssert(0, h.a[1], "reset forgot about elements", t)
	quickAssert(0, h.n, "reset forgot about n", t)
}

func TestMinIHLarge(t *testing.T) {
	type tuple struct {
		h         MinImplicitHeap
		toPush    []int
		shouldPop []int
	}

	table := []tuple{
		{MinImplicitHeap{}, []int{2, 1}, []int{1, 2}},
	}

	for i := 0; i < len(table); i++ {
		spamIMCheck(table[i].h, table[i].toPush, table[i].shouldPop, t)
	}
}

func addIHNodes(h ImplicitHeap, c int) {
	for i := 0; i < c; i++ {
		h.Push(i)
	}
}

func popIHNodes(h ImplicitHeap, c int, t *testing.T) {
	for i := 0; i < c; i++ {
		_, ok := h.Pop()

		if ok == false {
			t.Error("cannot pop")
		}
	}
}

func quickAssert(expected int, got int, fail string, t *testing.T) {
	if expected == got {
		return
	}

	t.Errorf("expected %v, got %v : %v", expected, got, fail)
}

func quickAssertBool(expected bool, got bool, fail string, t *testing.T) {
	if expected == got {
		return
	}

	t.Errorf("expected %v, got %v : %v", expected, got, fail)
}

func spamIMCheck(h MinImplicitHeap, toPush []int, shouldPop []int, t *testing.T) {
	for i := 0; i < len(toPush); i++ {
		h.Push(toPush[i])
	}

	for i := 0; i < len(shouldPop); i++ {
		v, ok := h.Pop()
		quickAssertBool(true, ok, "pop failed for ", t)
		quickAssert(shouldPop[i], v, "pop failed for ", t)
	}
}
