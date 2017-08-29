package lists

import (
	"sync"
	"testing"
	"time"
)

func testHQdeq(l *HierarchicalQueue, exp interface{}, t *testing.T) {
	val, err := l.Dequeue()

	if err != nil {
		t.Errorf("expected %v, got error %v", exp, err)
	} else if val != exp {
		t.Errorf("deq expected %v, got %v", exp, val)
	}
}

func testHQenq(l *HierarchicalQueue, elem interface{}, p uint8, t *testing.T) {
	err := l.Enqueue(elem, p)
	if err != nil {
		t.Errorf("enq failed at elem [%v] priority [%v] with %v", elem, p, err)
	}
}

func TestHQOne(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	testHQenq(l, "a", 0, t)
	testHQdeq(l, "a", t)
}

func TestHQPriorityBounds(t *testing.T) {
	l := NewHierarchicalQueue(3, false)
	testHQenq(l, "a", 0, t)
	testHQenq(l, "a", 1, t)
	testHQenq(l, "a", 2, t)
	testHQenq(l, "a", 3, t)
}
func TestHQOverflowPriority(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	if err := l.Enqueue("a", 2); err == nil {
		t.Error("enq a priority > max didn't returned an error")
	}
}
func TestHQTwo(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	testHQenq(l, "a", 0, t)
	testHQenq(l, "b", 0, t)

	testHQdeq(l, "a", t)
	testHQdeq(l, "b", t)
}

func TestHQReverse(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	testHQenq(l, "b", 1, t)
	testHQenq(l, "a", 0, t)

	testHQdeq(l, "a", t)
	testHQdeq(l, "b", t)
}

func TestHQDepletion(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	testHQenq(l, "a", 0, t)
	testHQdeq(l, "a", t)
	if l.IsDepleted() == false {
		t.Error("is not depleted after empty")
	}

	if _, err := l.Dequeue(); err == nil {
		t.Error("enq on a depleted HQ does not return error")
	}
}

func TestHQEnqAfterDeqBigger(t *testing.T) {
	l := NewHierarchicalQueue(2, false)
	testHQenq(l, "b", 1, t)
	testHQenq(l, "a", 0, t)

	testHQdeq(l, "a", t)
	testHQenq(l, "c", 2, t)

	testHQdeq(l, "b", t)
	testHQdeq(l, "c", t)
}

//TestHQEnqAfterDeqSmaller Edgecase, enq a smaller priority than the current one
func TestHQEnqAfterDeqSmaller(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	testHQenq(l, "c", 1, t)
	testHQenq(l, "a", 0, t)

	testHQdeq(l, "a", t) //should advance to 1 now
	if l.highestP != 1 {
		t.Errorf("highestP expected %v, got %v", 1, l.highestP)
	}
	testHQenq(l, "b", 0, t) //add 0 < 1

	testHQdeq(l, "b", t)
	testHQdeq(l, "c", t)
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

func TestHQConcurrencyManualLock(t *testing.T) {
	testHQLocks(true, t)
	testHQLocks(false, t)
}

func testHQLocks(autoLock bool, t *testing.T) {
	var lowestP uint8 = 50
	megaHQ := NewHierarchicalQueue(lowestP, autoLock)

	var group sync.WaitGroup

	//spam enqueue
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			var times uint8
			for ; times < 200; times++ {
				if autoLock == false {
					megaHQ.Lock()
				}
				if megaHQ.IsDepleted() {
					if autoLock == false {
						megaHQ.Unlock()
					}
					break
				}

				err := megaHQ.Enqueue("a", times%lowestP)

				if err != nil {
					t.Error(err)
				}
				if autoLock == false {

					megaHQ.Unlock()
				}
				time.Sleep(time.Millisecond * 10)
			}
			group.Done()
		}()
	}

	//spam dequeue
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			var times uint8
			for ; times < 200; times++ {
				if autoLock == false {
					megaHQ.Lock()
				}

				if megaHQ.IsDepleted() {
					if autoLock == false {
						megaHQ.Unlock()
					}
					break
				}
				_, err := megaHQ.Dequeue()

				if err != nil {
					t.Error(err)
				}

				if autoLock == false {
					megaHQ.Unlock()
				}
				time.Sleep(time.Millisecond * 10)
			}

			group.Done()
		}()
	}

	group.Wait()
}
