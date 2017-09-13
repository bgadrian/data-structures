package heap

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

//most of the functionalities are common and tested in min_*test

func TestIHMaxAddOrder1(t *testing.T) {
	h := NewImplicitHeapMax(false)

	h.Push(1, 1)
	h.Push(3, 3)
	h.Push(4, 4)

	testIMPriorityOrder(h.a, []int{4, 1, 3}, "push 1", t)

	h.Push(2, 2)
	testIMPriorityOrder(h.a, []int{4, 2, 3, 1}, "push 2", t)

	h.Push(5, 5)
	testIMPriorityOrder(h.a, []int{5, 4, 3, 1, 2}, "push 3", t)
}

func TestIHMaxPopFirst(t *testing.T) {
	h := NewImplicitHeapMax(false)

	_, ok := h.Pop()

	if ok == true {
		t.Error("pop didn't returned error when empty")
	}
}

func TestIHMaxDuplicates(t *testing.T) {

	table := []testIHTuple{
		{NewImplicitHeapMax(false),
			[]int{1, 1, 3, 3, 3},
			[]int{3, 3, 3, 1, 1}},
		{NewImplicitHeapMax(false),
			[]int{5, 1, 4, 3, 4, 1, 1},
			[]int{5, 4, 4, 3, 1, 1, 1}},
	}

	for i := 0; i < len(table); i++ {
		testIMPopOrder(table[i].h, table[i].toPush, table[i].shouldPop, t)
	}
}

func TestIHMaxLarge(t *testing.T) {

	a35 := []int{35, 35, 34, 34, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 13, 12, 11, 10, 10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 2, 2, 2, 2, 1}

	table := []testIHTuple{
		{NewImplicitHeapMax(false),
			[]int{2, 3, 13, 7, 10, 11, 20, 4, 2, 14, 12, 2, 22, 10, 18, 34, 5, 24, 34, 25, 2, 35, 32, 35, 34, 23, 26, 28, 13, 16, 9, 8, 33, 27, 2, 6, 1, 29, 10, 21, 19, 15, 30, 31, 17},
			a35},
		{NewImplicitHeapMax(false),
			[]int{1, 33, 4, 24, 12, 13, 18, 3, 35, 32, 27, 10, 13, 2, 2, 35, 34, 2, 17, 9, 10, 20, 29, 2, 8, 30, 21, 22, 26, 28, 25, 34, 7, 5, 23, 19, 15, 16, 2, 14, 34, 10, 6, 11, 31},
			a35},
		{NewImplicitHeapMax(false),
			[]int{5, 34, 35, 10, 23, 24, 6, 15, 32, 29, 12, 31, 18, 2, 27, 13, 34, 25, 2, 10, 1, 2, 20, 22, 16, 9, 13, 30, 7, 10, 11, 33, 28, 4, 3, 2, 17, 2, 19, 14, 21, 34, 35, 8, 26},
			a35},
	}

	for i := 0; i < len(table); i++ {
		testIMPopOrder(table[i].h, table[i].toPush, table[i].shouldPop, t)
	}
}

func TestIHMinAutoLock(t *testing.T) {
	runtime.GOMAXPROCS(1)
	testIHConcurrentSpam(NewImplicitHeapMin(true), true, t)
	testIHConcurrentSpam(NewImplicitHeapMin(false), false, t)
	testIHConcurrentSpam(NewImplicitHeapMax(true), true, t)
	testIHConcurrentSpam(NewImplicitHeapMax(false), false, t)

	runtime.GOMAXPROCS(runtime.NumCPU())
	testIHConcurrentSpam(NewImplicitHeapMin(true), true, t)
	testIHConcurrentSpam(NewImplicitHeapMin(false), false, t)
	testIHConcurrentSpam(NewImplicitHeapMax(true), true, t)
	testIHConcurrentSpam(NewImplicitHeapMax(false), false, t)
}

func testIHConcurrentSpam(h ImplicitHeap, autoLock bool, t *testing.T) {
	var group sync.WaitGroup

	pushes, pops := 0, 0

	for i := 0; i < 50; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 100; times++ {
				if autoLock == false {
					h.Lock()
				}

				h.Push(times, "a")
				pushes++

				if autoLock == false {
					h.Unlock()
				}
				time.Sleep(time.Millisecond * 10)
			}
			group.Done()
		}()
	}

	for i := 0; i < 50; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 150; times++ {
				if h.Len() == 0 {
					time.Sleep(1 * time.Millisecond)
				}

				if autoLock == false {
					h.Lock()
				}
				if h.HasElement() && h.IsDepleted() == false {
					_, ok := h.Peek()

					if ok == false && autoLock == false {
						t.Error("peek failed")
					}

					_, ok = h.Pop()
					pops++

					if ok == false && autoLock == false {
						t.Error("pop failed")
					}
				}

				if autoLock == false {
					h.Unlock()
				}
				time.Sleep(time.Millisecond * 10)
			}
			group.Done()
		}()
	}

	group.Wait()
	// fmt.Printf("pushes vs pops, %v vs %v", pushes, pops)
}
