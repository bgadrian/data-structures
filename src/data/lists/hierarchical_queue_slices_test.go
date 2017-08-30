package lists

import (
	"math"
	"runtime"
	"sync"
	"testing"
	"time"
)

func testHQSlicesdeq(l *HierarchicalQueueSlices, exp interface{}, t *testing.T) {
	val, err := l.Dequeue()

	if err != nil {
		t.Errorf("expected %v, got error %v", exp, err)
	} else if val != exp {
		t.Errorf("deq expected %v, got %v", exp, val)
	}
}

func testHQSlicesenq(l *HierarchicalQueueSlices, elem interface{}, p uint8, t *testing.T) {
	err := l.Enqueue(elem, p)
	if err != nil {
		t.Errorf("enq failed at elem [%v] priority [%v] with %v", elem, p, err)
	}
}

func TestHQSlicesOne(t *testing.T) {
	l := NewHierarchicalQueueSlices(1, false)
	testHQSlicesenq(l, "a", 0, t)
	testHQSlicesdeq(l, "a", t)
}

func TestHQSlicesPriorityBounds(t *testing.T) {
	l := NewHierarchicalQueueSlices(math.MaxUint8, false)
	testHQSlicesenq(l, "a", 0, t)
	testHQSlicesenq(l, "a", 1, t)
	testHQSlicesenq(l, "a", math.MaxUint8-1, t)
	testHQSlicesenq(l, "a", math.MaxUint8, t)
}
func TestHQSlicesOverflowPriority(t *testing.T) {
	l := NewHierarchicalQueueSlices(1, false)
	if err := l.Enqueue("a", 2); err == nil {
		t.Error("enq a priority > max didn't returned an error")
	}
}
func TestHQSlicesTwo(t *testing.T) {
	l := NewHierarchicalQueueSlices(1, false)
	testHQSlicesenq(l, "a", 0, t)
	testHQSlicesenq(l, "b", 0, t)

	testHQSlicesdeq(l, "a", t)
	testHQSlicesdeq(l, "b", t)
}

func TestHQSlicesReverse(t *testing.T) {
	l := NewHierarchicalQueueSlices(1, false)
	testHQSlicesenq(l, "b", 1, t)
	testHQSlicesenq(l, "a", 0, t)

	testHQSlicesdeq(l, "a", t)
	testHQSlicesdeq(l, "b", t)
}

func TestHQSlicesDeqFirst(t *testing.T) {
	l := NewHierarchicalQueueSlices(1, false)

	if _, err := l.Dequeue(); err == nil {
		t.Error("enq on an empty HQ does not return error")
	}

	if l.IsDepleted() == false {
		t.Error("deq an empty HQ should have depleted it")
	}
}
func TestHQSlicesDeqDepleted(t *testing.T) {
	l := NewHierarchicalQueueSlices(1, false)
	l.Dequeue()

	if _, err := l.Dequeue(); err == nil {
		t.Error("deq on a depleted HQ does not return error")
	}

	if err := l.Enqueue("a", 0); err == nil {
		t.Error("enq on a depleted HQ does not return error")
	}
}

func TestHQSlicesEnqAfterDeqBigger(t *testing.T) {
	l := NewHierarchicalQueueSlices(2, false)
	testHQSlicesenq(l, "b", 1, t)
	testHQSlicesenq(l, "a", 0, t)

	testHQSlicesdeq(l, "a", t)
	testHQSlicesenq(l, "c", 2, t)

	testHQSlicesdeq(l, "b", t)
	testHQSlicesdeq(l, "c", t)
}

//TestHQSlicesEnqAfterDeqSmaller Edgecase, enq a smaller priority than the current one
func TestHQSlicesEnqAfterDeqSmaller(t *testing.T) {
	l := NewHierarchicalQueueSlices(1, false)
	testHQSlicesenq(l, "c", 1, t)
	testHQSlicesenq(l, "a", 0, t)

	testHQSlicesdeq(l, "a", t) //should advance to 1 now
	if l.highestP != 1 {
		t.Errorf("highestP expected %v, got %v", 1, l.highestP)
	}
	testHQSlicesenq(l, "b", 0, t) //add 0 < 1

	testHQSlicesdeq(l, "b", t)
	testHQSlicesdeq(l, "c", t)
}

func TestHQSlicesOrderedMultipleTypes(t *testing.T) {
	table := [][]interface{}{
		{"a", "b", "c"},
		{1},
		{nil, nil},
		{1, "a", 2, "b"},
		{true, false, true, true},
		{true, "a", 1, true},
	}

	for _, arr := range table {
		testHQSlicesArrKeyIsPriority(arr, t)
	}
}

//enqAllAndDeq Enqueue and dequeue a list of elements
//for simplicity value == priority
func testHQSlicesArrKeyIsPriority(arr []interface{}, t *testing.T) {
	l := NewHierarchicalQueueSlices(uint8(len(arr)-1), false)

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

func TestHQSlicesConcurrencyManualLock(t *testing.T) {

	runtime.GOMAXPROCS(1)
	testHQSlicesLocks(true, t)
	testHQSlicesLocks(false, t)

	runtime.GOMAXPROCS(runtime.NumCPU())
	testHQSlicesLocks(true, t)
	testHQSlicesLocks(false, t)
}

func testHQSlicesLocks(autoLock bool, t *testing.T) {
	var lowestP uint8 = 50
	megaHQ := NewHierarchicalQueueSlices(lowestP, autoLock)

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

	group.Wait()
}

/* ********************** priorities between 0 - 50 **********************************************/

//BenchmarkHQSlicesSyncEnqDeqOn1000Size50P 1 Enqueue and 1 Dequeue cost in a HQ with 1.000 elements
func BenchmarkHQSlicesSyncEnqDeqOn1000Size50P(b *testing.B) {
	benchmarkHQSlicesSyncEnqDeqOne(1000, 50, b)
}

//BenchmarkHQSlicesSyncEnqDeqOn100000Size50P 1 Enqueue and 1 Dequeue cost in a HQ with 100.000 elements
func BenchmarkHQSlicesSyncEnqDeqOn100000Size50P(b *testing.B) {
	benchmarkHQSlicesSyncEnqDeqOne(100000, 50, b)
}

//BenchmarkHQSlicesSyncEnqDeqOn1000000Size50P 1 Enqueue and 1 Dequeue cost in a HQ with 1.000.000 elements
// func BenchmarkHQSlicesSyncEnqDeqOn1000000Size50P(b *testing.B) {
// 	benchmarkHQSlicesSyncEnqDeqOne(1000000, 50, b)
// }

// //BenchmarkHQSlicesSyncEnqDeqOn10000000Size50P 1 Enqueue and 1 Dequeue cost in a HQ with 10.000.000 elements
// func BenchmarkHQSlicesSyncEnqDeqOn10000000Size50P(b *testing.B) {
// 	benchmarkHQSlicesSyncEnqDeqOne(10000000, 50, b)
// }

// //BenchmarkHQSlicesSyncEnqDeqOn100000000Size50P 1 Enqueue and 1 Dequeue cost in a HQ with 1000.000.000 elements
// func BenchmarkHQSlicesSyncEnqDeqOn100000000Size50P(b *testing.B) {
// 	benchmarkHQSlicesSyncEnqDeqOne(100000000, 50, b)
// }

/* ********************** priorities between 0 - 255 **********************************************/

//BenchmarkHQSlicesSyncEnqDeqOn1000Size255P 1 Enqueue and 1 Dequeue cost in a HQ with 1.000 elements
func BenchmarkHQSlicesSyncEnqDeqOn1000Size255P(b *testing.B) {
	benchmarkHQSlicesSyncEnqDeqOne(1000, 255, b)
}

//BenchmarkHQSlicesSyncEnqDeqOn100000Size255P 1 Enqueue and 1 Dequeue cost in a HQ with 100.000 elements
func BenchmarkHQSlicesSyncEnqDeqOn100000Size255P(b *testing.B) {
	benchmarkHQSlicesSyncEnqDeqOne(100000, 255, b)
}

//BenchmarkHQSlicesSyncEnqDeqOn1000000Size255P 1 Enqueue and 1 Dequeue cost in a HQ with 1.000.000 elements
// func BenchmarkHQSlicesSyncEnqDeqOn1000000Size255P(b *testing.B) {
// 	benchmarkHQSlicesSyncEnqDeqOne(1000000, 255, b)
// }

// //BenchmarkHQSlicesSyncEnqDeqOn10000000Size255P 1 Enqueue and 1 Dequeue cost in a HQ with 10.000.000 elements
// func BenchmarkHQSlicesSyncEnqDeqOn10000000Size255P(b *testing.B) {
// 	benchmarkHQSlicesSyncEnqDeqOne(10000000, 255, b)
// }

// //BenchmarkHQSlicesSyncEnqDeqOn100000000Size255P 1 Enqueue and 1 Dequeue cost in a HQ with 100.000.000 elements
// func BenchmarkHQSlicesSyncEnqDeqOn100000000Size255P(b *testing.B) {
// 	benchmarkHQSlicesSyncEnqDeqOne(100000000, 255, b)
// }

//benchmarkHQSlicesSyncEnqDeqOne Measure time for 1 Enqueue and 1 Dequeue
func benchmarkHQSlicesSyncEnqDeqOne(count int, lowestP uint8, b *testing.B) {

	l := NewHierarchicalQueueSlices(lowestP, false)
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

//BenchmarkHQSlicesSync1000E100P Time to build, fill and deplete a HQ with 1.000 elements
// func BenchmarkHQSlicesSync1000E100P(b *testing.B) {
// 	benchHQSyncOne(10, 100, b)
// }

// //BenchmarkHQSlicesSync100000E100P Time to build, fill and deplete a HQ with 100.000 elements
// func BenchmarkHQSlicesSync100000E100P(b *testing.B) {
// 	benchHQSyncOne(1000, 100, b)
// }

//BenchmarkHQSlicesSync1000000E100P Time to build, fill and deplete a HQ with 1.000.000 elements
// func BenchmarkHQSlicesSync1000000E100P(b *testing.B) {
// 	benchHQSyncOne(1000000, 100, b)
// }

//benchHQSyncOne run b.N tests of: HQ list with [elements] count of "a" valued nodes with priorities between 0-lowestP
//the priorities are added 0-lowestP,0-lowestP...
// func benchHQSyncOne(elements int, lowestP uint8, b *testing.B) {
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {

// 		l := NewHierarchicalQueueSlices(lowestP, false)
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
