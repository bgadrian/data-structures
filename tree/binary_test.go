package tree

import "testing"
import "sort"
import "math/rand"

import "fmt"
import "strconv"

func TestBSTBasic(t *testing.T) {
	b := BST{}

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

func TestBSTInsertOrder(t *testing.T) {
	b := BST{}

	var r, rl, rr *Node          //first level of the tree
	var rll, rlr, rrl, rrr *Node //2nd level
	var rlll, rllr *Node         //3rd level

	/*
						42
				30				50
			20		35		46		60
		10		25

		inorder: 10,20,25,30,35,42,46,50,60
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

func TestBSTInOrder(t *testing.T) {
	b := BST{}

	b.Insert(42, nil)
	b.Insert(30, nil)
	b.Insert(50, nil)
	b.Insert(20, nil)
	b.Insert(35, nil)
	b.Insert(46, nil)
	b.Insert(60, nil)
	b.Insert(10, nil)
	b.Insert(25, nil)

	should := []int{10, 20, 25, 30, 35, 42, 46, 50, 60}
	got := b.Inorder()

	quickAssertArrInt(should, got, fmt.Sprintf("inorder exp %v, got %v", should, got), t)
}

func TestBSTSearch(t *testing.T) {
	b := BST{}

	b.InsertKeys(42, 30, 50, 20, 35, 46)

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
		t.Error("found a ghost node ?!")
	}
}

func TestBSTRemoveAndSize(t *testing.T) {
	b := BST{}

	err := b.Remove(42)
	if err == nil {
		t.Error("removing an existing node didn't returned error")
	}

	b.Insert(42, nil)
	quickAssertInt(1, b.Size(), "1 size after add root", t)

	err = b.Remove(42)
	if err != nil {
		t.Error(err)
	}

	if b.Root != nil {
		t.Errorf("removing the root, only node didn't worked, %+v", b.Root)
	}

	quickAssertInt(0, b.Size(), "0 size after remove root", t)

	var count int
	count, err = b.InsertKeys(3, 1, 2, 3)

	if err == nil {
		t.Error("InsertKeys duplicate key didn't returned error")
	}

	quickAssertInt(3, count, "InsertKeys returned wrong value", t)

	b.Remove(3)
	quickAssertArrInt([]int{1, 2}, b.Inorder(), "after remove root with 2 kids", t)
	b.Remove(1)
	quickAssertArrInt([]int{2}, b.Inorder(), "after remove root with 1 kids", t)
	b.Remove(2)

	//remove not existing keys
	b.InsertKeys(2, 1, 3)
	err = b.Remove(0)
	if err == nil {
		t.Error("not existing key remove should return error")
	}
	err = b.Remove(4)
	if err == nil {
		t.Error("not existing key remove should return error")
	}
}

func TestBSTRemoveTable(t *testing.T) {
	table := [][]int{
		{5, 1, 3},
		{5, 1, 3, 9, 10, 2, 4, 11},
		{42, 30, 11, 23, 50, 60, 70, 33, 90, 21, 99},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8, 0, -1, -2, -3, -4},
		{40, 20, 60, 10, 30, 50, 70},
		{40, 20, 60, 10, 30, 50, 70, 25, 35, 15, 45, 55, 65},
		{9, 8, 7, 6, 5, 4, 3},
		{20, 19, 21, 22, 18, 23, 17, 16, 15, 24, 25, 26},
		{20, 19, 21, 22, 18, 23, 17, 16, 15, 24, 25, 26, 1, 2, 3, 4, 5, 6},
	}

	for _, v := range table {
		bstRemoveAutoTest(v, t)
	}
}

func bstRemoveAutoTest(k []int, t *testing.T) {
	b := BST{}
	b.InsertKeys(k...)

	rand.Seed(1)

	//choose a random value and remove it from both lists
	removeValue := func(slice []int, v int) []int {
		s := 0
		for i, val := range slice {
			if val == v {
				s = i
				break
			}
		}
		return append(slice[:s], slice[s+1:]...)
	}

	//sort the entries, so we can easily compare with
	//the inorder array
	sort.Ints(k)

	for b.n > 0 {
		val := k[rand.Intn(b.n)]
		k = removeValue(k, val)

		b.Remove(val)
		got := b.Inorder()

		quickAssertArrInt(k, got, "fail after removing "+strconv.Itoa(val), t)
	}
}

func quickAssertStr(exp, got string, m string, t *testing.T) {
	if exp == got {
		return
	}
	t.Errorf("exp %v, got %v : %v", exp, got, m)
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

func quickAssertArrInt(exp, got []int, m string, t *testing.T) {
	if len(exp) != len(got) {
		t.Error(m)
		return
	}

	for i, v := range exp {
		if v != got[i] {
			t.Error(m)
			return
		}
	}
}

func ExampleBST() {
	b := BST{}

	_, err := b.Insert(42, nil)

	if err != nil {
		//do something
	}

	b.InsertKeys(30, 50, 24, 60)
	fmt.Printf("elements in order:%+v\n", b.Inorder())

	b.Remove(50)
	fmt.Printf("elements in order:%+v\n", b.Inorder())

	fmt.Printf("root is:%v\n", b.Root.K)
	fmt.Printf("left of root is:%v\n", b.Root.Left.K)

	// Output:
	//elements in order:[24 30 42 50 60]
	//elements in order:[24 30 42 60]
	//root is:42
	//left of root is:30
}
