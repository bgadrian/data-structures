package tree

import "testing"

func TestBinaryBasic(t *testing.T) {
	b := BinarySearch{}

	n, err := b.Insert(42, nil)
	quickAssertErr(err, t)
	quickAssertInt(42, n.K, "Node returned has wrong value", t)

	n, err = b.Insert(7, nil)
	quickAssertErr(err, t)
	quickAssertInt(7, n.K, "Node returned has wrong value", t)

	n, err = b.Insert(84, nil)
	quickAssertErr(err, t)
	quickAssertInt(84, n.K, "Node returned has wrong value", t)

	_, err = b.Insert(7, nil)

	if err == nil {
		t.Error("inserting same value should have throw error")
	}
}

func TestBinaryInsertOrder(t *testing.T) {
	b := BinarySearch{}

	var r, rl, rr *Node          //first level of the tree
	var rll, rlr, rrl, rrr *Node //2nd level
	var rlll, rllr *Node         //3rd level

	/*
						42
				30				50
			20		35		46		60
		10		25
	*/

	r, _ = b.Insert(42, nil)
	rl, _ = b.Insert(30, nil)
	rr, _ = b.Insert(50, nil)

	rll, _ = b.Insert(20, nil)
	rlr, _ = b.Insert(35, nil)
	rrl, _ = b.Insert(46, nil)
	rrr, _ = b.Insert(60, nil)

	rlll, _ = b.Insert(10, nil)
	rllr, _ = b.Insert(25, nil)

	quickAssert(b.Root, r, "root", t)
	quickAssert(b.Root.Left, rl, "root left", t)
	quickAssert(b.Root.Right, rr, "root right", t)

	quickAssert(b.Root.Left.Left, rll, "r left left", t)
	quickAssert(b.Root.Left.Right, rlr, "r left right", t)
	quickAssert(b.Root.Right.Left, rrl, "r Right left", t)
	quickAssert(b.Root.Right.Right, rrr, "r Right right", t)

	quickAssert(b.Root.Left.Left.Left, rlll, "r L L L", t)
	quickAssert(b.Root.Left.Left.Right, rllr, "r L L R", t)
}

func TestBinarySearch(t *testing.T) {
	b := BinarySearch{}

	b.Insert(42, nil)
	b.Insert(30, nil)
	b.Insert(50, nil)
	b.Insert(20, nil)
	b.Insert(35, nil)
	b.Insert(46, nil)

	n := b.Search(42)
	if n == nil {
		t.Error("cannot find root by key")
	}

	b.Search(35)
	if n == nil {
		t.Error("cannot find element by key")
	}

	n = b.Search(99)
	if n != nil {
		t.Error("found a missing node ?!")
	}
}

func quickAssert(exp, got *Node, m string, t *testing.T) {
	if exp == got {
		return
	}
	t.Errorf("exp %v, got %v : %v", exp, got, m)
}

func quickAssertInt(exp, got int, m string, t *testing.T) {
	if exp == got {
		return
	}
	t.Errorf("exp %v, got %v : %v", exp, got, m)
}

func quickAssertErr(err error, t *testing.T) {
	if err == nil {
		return
	}
	t.Error(err)
}
