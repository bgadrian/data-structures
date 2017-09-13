package priorityqueue

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

/* ********************** priorities between 0 - 50 **********************************************/

func TestHHBasic(t *testing.T) {
	h, err := NewHierarchicalHeap(1, 0, 1, false)

	if err != nil {
		t.Error(err)
	}
	err = h.Enqueue("a", 1)

	if err != nil {
		t.Error(err)
	}

	quickAssert(1, h.Len(), "length after 1 enqueue", t)

	var v interface{}
	v, err = h.Dequeue()

	if err != nil {
		t.Error(err)
		return
	}

	if v != "a" {
		t.Errorf("expected %v, got %v", "a", v)
	}
}

func TestHHLen(t *testing.T) {
	h, _ := NewHierarchicalHeap(2, 0, 1, false)

	h.Enqueue("a", 0)
	h.Enqueue("a", 1)

	quickAssert(2, h.Len(), "length bucket after 2 enqueue", t)
	quickAssert(1, h.a[0].Len(), "length bucket 1 after 1 enqueue", t)
	quickAssert(1, h.a[1].Len(), "length bucket 2 after 1 enqueue", t)
}

func TestHHWrongInit(t *testing.T) {
	h, err := NewHierarchicalHeap(0, 0, 1, false)

	if err == nil || h != nil {
		t.Error("doesn't return error with 0 buckets")
	}

	h, err = NewHierarchicalHeap(1, -1, 1, false)

	if err == nil || h != nil {
		t.Error("doesn't return error with negative min priority")
	}

	h, err = NewHierarchicalHeap(1, 0, -1, false)

	if err == nil || h != nil {
		t.Error("doesn't return error with smaller max priority")
	}
}

func TestHHOverflow(t *testing.T) {
	h, err := NewHierarchicalHeap(5, 0, 5, false)

	if err != nil {
		t.Error(err)
	}

	_, err = h.Dequeue()

	if err == nil {
		t.Error("dequeue didn't returned error when empty")
	}

	err = h.Enqueue("a", -1)

	if err != nil {
		t.Error(err)
	}

	err = h.Enqueue("a", 5)
	if err != nil {
		t.Error(err)
	}

	err = h.Enqueue("a", 10)
	if err != nil {
		t.Error(err)
	}
}

func TestHHO1Buckets(t *testing.T) {
	minP := 0 //inclusive
	maxP := 5 //inclusive
	buckets := maxP - minP + 1
	h, err := NewHierarchicalHeap(buckets, minP, maxP, false)
	if err != nil {
		t.Error(err)
	}

	for i := minP; i <= maxP; i++ {
		h.Enqueue("a", i)
		h.Enqueue("a", i)
	}

	for i := 0; i < buckets; i++ {
		if h.a[i] == nil {
			t.Errorf("bucket %v is not inited", i)
			continue
		}

		len := h.a[i].Len()
		if len != 2 {
			t.Errorf("each bucket should have 2 el, got %v", len)
		}
	}
}

func TestHHLarge(t *testing.T) {
	testHHSyncEnqDeqOne(1, 100, 100, 100, t)
	testHHSyncEnqDeqOne(2, 100, 100, 1, t)
	testHHSyncEnqDeqOne(1, 1000, 1000, 100, t)
	testHHSyncEnqDeqOne(1, 1000, 100, 100, t)
	testHHSyncEnqDeqOne(1, 1000, 10, 100, t)
}

func testHHSyncEnqDeqOne(seed int64, count int, lowestP int, buckets int, b *testing.T) {
	h, err := NewHierarchicalHeap(buckets, 0, lowestP, false)
	if err != nil {
		b.Error(err)
		return
	}

	rand.Seed(seed)

	for i := 0; i < count; i++ {
		h.Enqueue("a", rand.Intn(lowestP))
	}

	for i := 0; i < 100; i++ {
		h.Enqueue("a", rand.Intn(lowestP))

		_, err = h.Dequeue()

		if err != nil {
			b.Error(err)
		}
	}
}

func TestHHConcurrencyLock(t *testing.T) {

	runtime.GOMAXPROCS(1)
	testHHLocks(true, t)
	testHHLocks(false, t)

	runtime.GOMAXPROCS(runtime.NumCPU())
	testHHLocks(true, t)
	testHHLocks(false, t)
}

func testHHLocks(autoLock bool, t *testing.T) {
	lowestP := 50
	megaHH, err := NewHierarchicalHeap(10, 0, 100, autoLock)

	if err != nil {
		t.Error(err)
		return
	}

	var group sync.WaitGroup

	lock := func() {
		if autoLock == false {
			megaHH.Lock()
		}
	}

	unlock := func() {
		if autoLock == false {
			megaHH.Unlock()
		}
	}

	//spam enqueue
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 20; times++ {
				lock()

				megaHH.Enqueue("a", times%lowestP)

				if err != nil {
					t.Error(err)
				}
				unlock()
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

			for times := 0; times < 20; times++ {
				lock()
				_, err := megaHH.Dequeue()

				if err != nil {
					t.Error(err)
				}

				unlock()
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

			for times := 0; times < 10; times++ {
				lock()
				megaHH.Len()

				unlock()
				time.Sleep(time.Millisecond * 20)
			}

			group.Done()
		}()
	}

	group.Wait()
}

func BenchmarkHHSyncEnqDeqOn1000Size50P5B(b *testing.B) {
	benchmarkHHSyncEnqDeqOne(1000, 50, 5, b)
}

func BenchmarkHHSyncEnqDeqOn1000Size50P25B(b *testing.B) {
	benchmarkHHSyncEnqDeqOne(1000, 50, 25, b)
}

func BenchmarkHHSyncEnqDeqOn100000Size50P10B(b *testing.B) {
	benchmarkHHSyncEnqDeqOne(100000, 50, 10, b)
}

func BenchmarkHHSyncEnqDeqOn1000000Size50P10B(b *testing.B) {
	benchmarkHHSyncEnqDeqOne(1000000, 50, 10, b)
}

func BenchmarkHHSyncEnqDeqOn10000000Size500P5B(b *testing.B) {
	benchmarkHHSyncEnqDeqOne(10000000, 500, 5, b)
}

func BenchmarkHHSyncEnqDeqOn10000000Size500P50B(b *testing.B) {
	benchmarkHHSyncEnqDeqOne(10000000, 500, 50, b)
}

func BenchmarkHHSyncEnqDeqOn10000000Size500P501B(b *testing.B) {
	benchmarkHHSyncEnqDeqOne(10000000, 500, 501, b)
}

//benchmarkHHSyncEnqDeqOne Measure time for 1 Enqueue and 1 Dequeue
func benchmarkHHSyncEnqDeqOne(count int, lowestP int, buckets int, b *testing.B) {
	h, err := NewHierarchicalHeap(buckets, 0, lowestP, false)
	if err != nil {
		b.Error(err)
		return
	}
	rand.Seed(2)

	for i := 0; i < count; i++ {
		h.Enqueue("a", rand.Intn(lowestP))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Enqueue("a", rand.Intn(lowestP))

		_, err = h.Dequeue()

		if err != nil {
			b.Error(err)
			return
		}

	}
}
