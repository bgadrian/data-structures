package tree

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestIHMinAutoLock(t *testing.T) {
	testIHConcurrentSpam(NewImplicitHeapMin(true), true, t)
	testIHConcurrentSpam(NewImplicitHeapMin(false), false, t)
}

func testIHConcurrentSpam(h ImplicitHeap, autoLock bool, t *testing.T) {
	var group sync.WaitGroup

	pushes, pops := 0, 0

	for i := 0; i < 200; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
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

	for i := 0; i < 200; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 300; times++ {
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
	fmt.Printf("pushes vs pops, %v vs %v", pushes, pops)
}
