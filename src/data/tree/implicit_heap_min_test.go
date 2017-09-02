package tree

import (
	"testing"
)

type testIHTuple struct {
	h         ImplicitHeap
	toPush    []int
	shouldPop []int
}

func TestIHMinBasic(t *testing.T) {
	h := NewImplicitHeapMin(false)

	h.Push(6, 6)

	v, ok := h.Peek()
	if ok == false {
		t.Error("cannot Peek")
	}
	quickAssert(6, v, "basic Peek failed", t)

	quickAssertBool(false, h.IsDepleted(), "depleted with elements", t)
	quickAssertBool(true, h.HasElement(), "hasElement() is wrong", t)

	v, ok = h.Pop()
	if ok == false {
		t.Error("cannot pop")
	}
	quickAssert(6, v, "basic pop failed", t)

	v, ok = h.Pop()
	quickAssertBool(false, ok, "pop empty was ok", t)

	v, ok = h.Peek()
	quickAssertBool(false, ok, "peek empty was ok", t)

	quickAssertBool(true, h.IsDepleted(), "was empty but depleted says false", t)
	quickAssertBool(false, h.HasElement(), "was empty but hasElement is true", t)
}

func TestIHMinAddOrder1(t *testing.T) {
	h := NewImplicitHeapMin(false)

	h.Push(3, 3)
	h.Push(1, 1)
	h.Push(4, 4)
	testIMPriorityOrder(h.a, []int{1, 3, 4}, "push 1", t)

	h.Push(2, 2)
	testIMPriorityOrder(h.a, []int{1, 2, 4, 3}, "push 2", t)

	h.Push(0, 0)
	testIMPriorityOrder(h.a, []int{0, 1, 4, 3, 2}, "push 3", t)

	h.Push(-1, -1)
	testIMPriorityOrder(h.a, []int{-1, 1, 0, 3, 2, 4}, "push 4", t)
}

func TestIHMinDuplicates(t *testing.T) {

	table := []testIHTuple{
		{NewImplicitHeapMin(false),
			[]int{3, 1, 1},
			[]int{1, 1, 3}},
		{NewImplicitHeapMin(false),
			[]int{3, 3, 1},
			[]int{1, 3, 3}},
		{NewImplicitHeapMin(false),
			[]int{5, 1, 4, 3, 4, 1, 1},
			[]int{1, 1, 1, 3, 4, 4, 5}},
	}

	for i := 0; i < len(table); i++ {
		testIMPopOrder(table[i].h, table[i].toPush, table[i].shouldPop, t)
	}
}

func TestIHMinPopOrder1(t *testing.T) {
	h := NewImplicitHeapMin(false)

	h.Push(5, 5)
	h.Push(3, 3)
	h.Push(1, 1)

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

func TestIHMinCapacity(t *testing.T) {
	h := NewImplicitHeapMin(false)

	h.Push(101, 101)
	quickAssert(8, cap(h.a), "capacity is default", t)

	addIHNodes(h, 8)
	quickAssert(16, cap(h.a), "capacity doubled", t)

	addIHNodes(h, 17)
	quickAssert(32, cap(h.a), "capacity doubled 2", t)
	quickAssert(26, h.n, "count is 26", t)

	popIHNodes(h, 19, t)
	quickAssert(16, cap(h.a), "capacity shrinked /2", t)

	popIHNodes(h, int(h.n), t)
	quickAssert(8, cap(h.a), "capacity never drop 8", t)

	// quickAssert(0, h.Levels(), "levels() after init", t)
}

func TestIHMinReset(t *testing.T) {
	h := NewImplicitHeapMin(false)

	h.Push(1, 1)
	h.Push(2, 2)

	h.Reset()
	quickAssert(0, h.a[0].priority, "reset forgot about elements", t)
	quickAssert(0, h.a[1].priority, "reset forgot about elements", t)
	quickAssert(0, h.n, "reset forgot about n", t)
}

func TestIHMinLarge(t *testing.T) {

	a35 := []int{1, 2, 2, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 11, 12, 13, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 34, 34, 35, 35}

	table := []testIHTuple{
		{NewImplicitHeapMin(false),
			[]int{2, 3, 13, 7, 10, 11, 20, 4, 2, 14, 12, 2, 22, 10, 18, 34, 5, 24, 34, 25, 2, 35, 32, 35, 34, 23, 26, 28, 13, 16, 9, 8, 33, 27, 2, 6, 1, 29, 10, 21, 19, 15, 30, 31, 17},
			a35},
		{NewImplicitHeapMin(false),
			[]int{1, 33, 4, 24, 12, 13, 18, 3, 35, 32, 27, 10, 13, 2, 2, 35, 34, 2, 17, 9, 10, 20, 29, 2, 8, 30, 21, 22, 26, 28, 25, 34, 7, 5, 23, 19, 15, 16, 2, 14, 34, 10, 6, 11, 31},
			a35},
		{NewImplicitHeapMin(false),
			[]int{5, 34, 35, 10, 23, 24, 6, 15, 32, 29, 12, 31, 18, 2, 27, 13, 34, 25, 2, 10, 1, 2, 20, 22, 16, 9, 13, 30, 7, 10, 11, 33, 28, 4, 3, 2, 17, 2, 19, 14, 21, 34, 35, 8, 26},
			a35},
	}

	for i := 0; i < len(table); i++ {
		testIMPopOrder(table[i].h, table[i].toPush, table[i].shouldPop, t)
	}
}

func addIHNodes(h ImplicitHeap, c int) {
	for i := 0; i < c; i++ {
		h.Push(i, i)
	}
}

func addIHNodesWithMaxP(h ImplicitHeap, c int, maxP int) {
	for i := 0; i < c; i++ {
		h.Push(i%maxP, i)
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

func quickAssert(expected interface{}, got interface{}, fail string, t *testing.T) {
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

func testIMPopOrder(h ImplicitHeap, toPush []int, shouldPop []int, t *testing.T) {
	for _, to := range toPush {
		h.Push(to, to)
	}

	for _, should := range shouldPop {
		v, ok := h.Pop()
		if ok == false {
			t.Errorf("pop failed for %v", toPush)
			break
		}
		if should != v {
			t.Errorf("expected %v, got %v , from %v", should, v, toPush)
			break
		}
	}
}

func testIMPriorityOrder(a []implicitHeapNode, shouldBe []int, msg string, t *testing.T) {
	for i, should := range shouldBe {
		if a[i].priority != should {
			t.Errorf("%v : priority order in %v is not good %v", msg, i, a)
			break
		}
	}
}
