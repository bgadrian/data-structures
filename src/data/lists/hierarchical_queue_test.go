package lists

import (
	"testing"
)

func testHQenq(l *HierarchicalQueue, exp interface{}, t *testing.T) {
	val, err := l.Dequeue()

	if err != nil {
		t.Errorf("expected %v, got error %v", exp, err)
	} else if val != exp {
		t.Errorf("deq expected %v, got %v", exp, val)
	}
}

func testHQdeq(l *HierarchicalQueue, elem interface{}, p uint8, t *testing.T) {
	err := l.Enqueue(elem, p)
	if err != nil {
		t.Errorf("enq failed at elem [%v] priority [%v] with %v", elem, p, err)
	}
}

func TestHQOne(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	testHQdeq(l, "a", 0, t)
	testHQenq(l, "a", t)
}
func TestHQTwo(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	testHQdeq(l, "a", 0, t)
	testHQdeq(l, "b", 0, t)

	testHQenq(l, "a", t)
	testHQenq(l, "b", t)
}

func TestHQReverse(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	testHQdeq(l, "b", 1, t)
	testHQdeq(l, "a", 0, t)

	testHQenq(l, "a", t)
	testHQenq(l, "b", t)
}

func TestHQDepletion(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	testHQdeq(l, "a", 0, t)
	testHQenq(l, "a", t)
	if l.IsDepleted() == false {
		t.Error("is not depleted after empty")
	}

	if _, err := l.Dequeue(); err == nil {
		t.Error("enq on a depleted HQ does not return error")
	}
}

func TestHQEnqAfterDeqBigger(t *testing.T) {
	l := NewHierarchicalQueue(2, false)
	testHQdeq(l, "b", 1, t)
	testHQdeq(l, "a", 0, t)

	testHQenq(l, "a", t)
	testHQdeq(l, "c", 2, t)

	testHQenq(l, "b", t)
	testHQenq(l, "c", t)
}

//TestHQEnqAfterDeqSmaller Edgecase, enq a smaller priority than the current one
func TestHQEnqAfterDeqSmaller(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	testHQdeq(l, "c", 1, t)
	testHQdeq(l, "a", 0, t)

	testHQenq(l, "a", t) //should advance to 1 now
	if l.highestP != 1 {
		t.Errorf("highestP expected %v, got %v", 1, l.highestP)
	}
	testHQdeq(l, "b", 0, t) //add 0 < 1

	testHQenq(l, "b", t)
	testHQenq(l, "c", t)
}

func TestHQOrderedMultipleTypes(t *testing.T) {
	table := [][]interface{}{
		{"a", "b", "c"},
		{1},
		{nil, nil},
		{1, "a", 2, "b"},
		{true, false, true, true},
		{true, "a", 1, true},
	}

	for _, arr := range table {
		hqTestArrKeyIsPriority(arr, t)
	}
}

//enqAllAndDeq Enqueue and dequeue a list of elements
//for simplicity value == priority
func hqTestArrKeyIsPriority(arr []interface{}, t *testing.T) {
	l := NewHierarchicalQueue(uint8(len(arr)-1), false)

	for p, v := range arr {
		err := l.Enqueue(v, uint8(p))

		if err != nil {
			t.Errorf("enq failed for %v in %v", v, arr)
		}
	}

	//should be N elements in the queue
	for _, v := range arr {
		elem, err := l.Dequeue()

		if err != nil {
			t.Error(err)
		} else if v != elem {
			t.Errorf("deq failed, expected %v got %v for %v", v, elem, arr)
		}
	}

	if l.IsDepleted() == false {
		t.Errorf("HQ is not depleted after deq all elem %v", arr)
	}
}
