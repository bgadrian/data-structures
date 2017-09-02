package priorityqueue

import (
	"math"
	"runtime"
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
	l := NewHierarchicalQueue(math.MaxUint8, false)
	testHQenq(l, "a", 0, t)
	testHQenq(l, "a", 1, t)
	testHQenq(l, "a", math.MaxUint8-1, t)
	testHQenq(l, "a", math.MaxUint8, t)
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

func TestHQLenghts(t *testing.T) {
	l := NewHierarchicalQueue(2, false)

	quickAssert(0, l.Len(), "Len() after init", t)
	quickAssert(0, l.LenPriority(0), "LenPriority() after init", t)
	quickAssert(0, l.LenPriority(1), "LenPriority() after init", t)
	quickAssert(0, l.LenPriority(2), "LenPriority() after init", t)

	l.Enqueue("a", 0)

	quickAssert(1, l.Len(), "Len() after 1enq", t)
	quickAssert(1, l.LenPriority(0), "LenPriority(0) after 1enq", t)
	quickAssert(0, l.LenPriority(1), "LenPriority(1) after 1enq", t)
	quickAssert(0, l.LenPriority(2), "LenPriority(2) after 1enq", t)

	l.Enqueue("a", 1)
	l.Enqueue("a", 1)
	l.Enqueue("a", 2)

	quickAssert(4, l.Len(), "Len() after nenq", t)
	quickAssert(1, l.LenPriority(0), "LenPriority(0) after nenq", t)
	quickAssert(2, l.LenPriority(1), "LenPriority(1) after nenq", t)
	quickAssert(1, l.LenPriority(2), "LenPriority(2) after nenq", t)

	l.Dequeue()
	l.Dequeue()

	quickAssert(2, l.Len(), "Len() after deq", t)
	quickAssert(0, l.LenPriority(0), "LenPriority(0) after deq", t)
	quickAssert(1, l.LenPriority(1), "LenPriority(1) after deq", t)
	quickAssert(1, l.LenPriority(2), "LenPriority(2) after deq", t)
}

func quickAssert(expected int, got int, fail string, t *testing.T) {
	if expected == got {
		return
	}

	t.Errorf("expected %v, got %v : %v", expected, got, fail)
}

func TestHQDeqFirst(t *testing.T) {
	l := NewHierarchicalQueue(1, false)

	if _, err := l.Dequeue(); err == nil {
		t.Error("enq on an empty HQ does not return error")
	}

	if l.IsDepleted() == false {
		t.Error("deq an empty HQ should have depleted it")
	}
}
func TestHQDeqDepleted(t *testing.T) {
	l := NewHierarchicalQueue(1, false)
	l.Dequeue()

	if _, err := l.Dequeue(); err == nil {
		t.Error("deq on a depleted HQ does not return error")
	}

	if err := l.Enqueue("a", 0); err == nil {
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

	runtime.GOMAXPROCS(1)
	testHQLocks(true, t)
	testHQLocks(false, t)

	runtime.GOMAXPROCS(runtime.NumCPU())
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
			//we must wait for at least a few Enqeue, otherwise it will finish before it started
			time.Sleep(time.Millisecond * 30)

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

	//spam lengths
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			//we must wait for at least a few Enqeue, otherwise it will finish before it started
			time.Sleep(time.Millisecond * 30)

			var times uint8
			for ; times < 50; times++ {
				if autoLock == false {
					megaHQ.Lock()
				}

				megaHQ.Len()
				megaHQ.LenPriority(2)

				if autoLock == false {
					megaHQ.Unlock()
				}
				time.Sleep(time.Millisecond * 20)
			}

			group.Done()
		}()
	}

	group.Wait()
}

/* ********************** priorities between 0 - 50 **********************************************/

//BenchmarkHQSyncEnqDeqOn1000Size50P 1 Enqueue and 1 Dequeue cost in a HQ with 1.000 elements
func BenchmarkHQSyncEnqDeqOn1000Size50P(b *testing.B) {
	benchmarkHQSyncEnqDeqOne(1000, 50, b)
}

//BenchmarkHQSyncEnqDeqOn100000Size50P 1 Enqueue and 1 Dequeue cost in a HQ with 100.000 elements
func BenchmarkHQSyncEnqDeqOn100000Size50P(b *testing.B) {
	benchmarkHQSyncEnqDeqOne(100000, 50, b)
}

//BenchmarkHQSyncEnqDeqOn1000000Size50P 1 Enqueue and 1 Dequeue cost in a HQ with 1.000.000 elements
func BenchmarkHQSyncEnqDeqOn1000000Size50P(b *testing.B) {
	benchmarkHQSyncEnqDeqOne(1000000, 50, b)
}

//BenchmarkHQSyncEnqDeqOn10000000Size50P 1 Enqueue and 1 Dequeue cost in a HQ with 10.000.000 elements
func BenchmarkHQSyncEnqDeqOn10000000Size50P(b *testing.B) {
	benchmarkHQSyncEnqDeqOne(10000000, 50, b)
}

//BenchmarkHQSyncEnqDeqOn100000000Size50P 1 Enqueue and 1 Dequeue cost in a HQ with 1000.000.000 elements
func BenchmarkHQSyncEnqDeqOn100000000Size50P(b *testing.B) {
	benchmarkHQSyncEnqDeqOne(100000000, 50, b)
}

/* ********************** priorities between 0 - 255 **********************************************/

//BenchmarkHQSyncEnqDeqOn1000Size255P 1 Enqueue and 1 Dequeue cost in a HQ with 1.000 elements
func BenchmarkHQSyncEnqDeqOn1000Size255P(b *testing.B) {
	benchmarkHQSyncEnqDeqOne(1000, 255, b)
}

//BenchmarkHQSyncEnqDeqOn100000Size255P 1 Enqueue and 1 Dequeue cost in a HQ with 100.000 elements
func BenchmarkHQSyncEnqDeqOn100000Size255P(b *testing.B) {
	benchmarkHQSyncEnqDeqOne(100000, 255, b)
}

//BenchmarkHQSyncEnqDeqOn1000000Size255P 1 Enqueue and 1 Dequeue cost in a HQ with 1.000.000 elements
func BenchmarkHQSyncEnqDeqOn1000000Size255P(b *testing.B) {
	benchmarkHQSyncEnqDeqOne(1000000, 255, b)
}

//BenchmarkHQSyncEnqDeqOn10000000Size255P 1 Enqueue and 1 Dequeue cost in a HQ with 10.000.000 elements
func BenchmarkHQSyncEnqDeqOn10000000Size255P(b *testing.B) {
	benchmarkHQSyncEnqDeqOne(10000000, 255, b)
}

//BenchmarkHQSyncEnqDeqOn100000000Size255P 1 Enqueue and 1 Dequeue cost in a HQ with 100.000.000 elements
func BenchmarkHQSyncEnqDeqOn100000000Size255P(b *testing.B) {
	benchmarkHQSyncEnqDeqOne(100000000, 255, b)
}

//benchmarkHQSyncEnqDeqOne Measure time for 1 Enqueue and 1 Dequeue
func benchmarkHQSyncEnqDeqOne(count int, lowestP uint8, b *testing.B) {

	l := NewHierarchicalQueue(lowestP, false)
	var err error

	for i := 0; i < count; i++ {
		l.Enqueue("a", uint8(i)%lowestP)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = l.Enqueue("a", uint8(i)%lowestP)

		if err != nil {
			b.Error(err)
		}

		_, err = l.Dequeue()

		if err != nil {
			b.Error(err)
		}
	}
}

//BenchmarkHQSync1000E100P Time to build, fill and deplete a HQ with 1.000 elements
// func BenchmarkHQSync1000E100P(b *testing.B) {
// 	benchHQSyncOne(10, 100, b)
// }

// //BenchmarkHQSync100000E100P Time to build, fill and deplete a HQ with 100.000 elements
// func BenchmarkHQSync100000E100P(b *testing.B) {
// 	benchHQSyncOne(1000, 100, b)
// }

//BenchmarkHQSync1000000E100P Time to build, fill and deplete a HQ with 1.000.000 elements
// func BenchmarkHQSync1000000E100P(b *testing.B) {
// 	benchHQSyncOne(1000000, 100, b)
// }

//benchHQSyncOne run b.N tests of: HQ list with [elements] count of "a" valued nodes with priorities between 0-lowestP
//the priorities are added 0-lowestP,0-lowestP...
// func benchHQSyncOne(elements int, lowestP uint8, b *testing.B) {
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {

// 		l := NewHierarchicalQueue(lowestP, false)
// 		var err error
// 		var p uint8

// 		for times := 0; times < elements; times++ {
// 			for p = 0; p <= lowestP; p++ {
// 				err = l.Enqueue("a", p)

// 				if err != nil {
// 					b.Error(err)
// 				}
// 			}
// 		}

// 		for l.IsDepleted() == false {
// 			_, err = l.Dequeue()

// 			if err != nil {
// 				b.Error(err)
// 			}
// 		}
// 	}
// }
