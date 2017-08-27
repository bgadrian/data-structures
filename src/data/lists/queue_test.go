package lists

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestQueueConcurrencyManualLock(t *testing.T) {
	megaQueue := NewQueue(false)

	var group sync.WaitGroup

	//spam enqueue
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				megaQueue.Lock()
				ok := megaQueue.Enqueue(times)

				if ok == false {
					t.Error("insert failed " + string(times))
				}

				megaQueue.Unlock()

				time.Sleep(time.Millisecond * 10)
			}
			group.Done()
		}()
	}

	//spam dequeue
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				megaQueue.Lock()
				if megaQueue.HasElement() {
					_, ok := megaQueue.Dequeue()

					if ok == false {
						t.Error("lock failed, hasElement() but Pop() failed")
					}
				}
				megaQueue.Unlock()
				time.Sleep(time.Millisecond * 8)
			}
			group.Done()
		}()
	}

	group.Wait()
}

func TestQueueConcurrencyAutoLock(t *testing.T) {

	megaStack := NewQueue(true)
	var group sync.WaitGroup

	//spam peek
	for i := 0; i <= 10; i++ {
		go func() {
			if megaStack.IsEmpty() == false {
				_, ok := megaStack.Peek()

				if ok == false {
					//we ignore the stack was empty, because of the fast times this will happen, 1 stack vs lots of workers
					//TODO learn a better way to do this
				}
			}
			//spam for test coverage
			megaStack.String()
			megaStack.Len()

			time.Sleep(time.Millisecond * 3)
		}()
	}

	//spam enqueue
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				ok := megaStack.Enqueue(times)

				if ok == false {
					t.Error("insert failed " + string(times))
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
			for times := 0; times < 200; times++ {
				if megaStack.HasElement() {
					_, ok := megaStack.Dequeue()

					if ok == false {
						//we ignore the stack was empty, because of the fast times this will happen, 1 stack vs lots of workers
						//TODO learn a better way to do this
					}
				}
				time.Sleep(time.Millisecond * 8)
			}
			group.Done()
		}()
	}

	group.Wait()
}

func TestQueueBasicTypes(t *testing.T) {
	for _, r := range fakeTable {
		testQueueFunctionality(t, r)
	}
}

func testQueueFunctionality(t *testing.T, toPush []interface{}) {
	s := NewQueue(false)

	for i, v := range toPush {
		ok := s.Enqueue(v)

		if ok == false {
			t.Error("Enqueue failed ")
		}
		len := s.Len()

		if len != i+1 {
			t.Errorf("len failed, expected %v, got %v, for %v", i, len, toPush)
		}

		value, ok := s.Peek()

		if ok == false {
			t.Error("peek failed")
		}

		if value != toPush[0] {
			t.Errorf("peek failed, expected %v, got %v ", v, value)
		}
	}

	for _, v := range toPush {
		el, ok := s.Dequeue()

		if ok == false {
			t.Error("dequeue failed")
		}

		if el != v {
			t.Errorf("dequeue failed, expected %v, got %v", v, el)
		}
	}

	if s.HasElement() {
		t.Errorf("stack is not empty after all Pop(), size=%v", s.Len())
	}
}

func TestQueueInitPeekIsNil(t *testing.T) {
	s := NewQueue(false)
	peek, ok := s.Peek()

	if ok {
		t.Error("peek should be false when used on an empty queue")
	}

	if peek != nil {
		t.Errorf("expected nil, got %v ", peek)
	}
}

func TestQueueInitPopIsNil(t *testing.T) {
	s := NewQueue(false)
	var pop, ok = s.Dequeue()

	if ok {
		t.Error("Pop should be false when used on an empty queue")
	}

	if pop != nil {
		t.Errorf("expected nil, got %v ", pop)
	}
}

func TestQueueInitIsEmpty(t *testing.T) {
	helperInitIsEmpty("queue", NewQueue(false), t)
}

func TestQueueSkippingNewShouldPanic(t *testing.T) {
	//TODO Should Queue implement lazyInit() method (like list.go has) so this should never happen?
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the code did not panic")
		}
	}()
	s := Queue{}

	c := s.Len()

	fmt.Print(c)
}

func TestQueueStringer(t *testing.T) {
	s := NewQueue(false)

	v := s.String()
	if v != "Queue [0]" {
		t.Error("stringer was incorrect for 0 length" + v)
	}

	s.Enqueue(1)
	s.Enqueue(1)

	v = s.String()
	if v != "Queue [2]" {
		t.Error("stringer was incorrect for 2 length" + v)
	}
}
