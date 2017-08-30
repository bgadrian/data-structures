package tree

import (
	"testing"
)

func TestMinIHBasic(t *testing.T) {
	h := MinImplicitHeap{}

	h.AddNode(6)
	v, ok := h.Pop()
	if ok == false {
		t.Error("cannot pop")
	}
	quickAssert(6, v, "basic pop 1", t)
}

func TestMinIHPopOrder1(t *testing.T) {
	h := MinImplicitHeap{}

	h.AddNode(1)
	h.AddNode(3)
	h.AddNode(5)

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

	h.AddNode(101)
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

func addIHNodes(h *MinImplicitHeap, c int) {
	for i := 0; i < c; i++ {
		h.AddNode(i)
	}
}

func popIHNodes(h *MinImplicitHeap, c int, t *testing.T) {
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
