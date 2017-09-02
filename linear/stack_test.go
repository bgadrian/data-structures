package linear

import "testing"
import "time"
import "sync"
import "fmt"

func TestConcurrencyManualLock(t *testing.T) {
	megaStack := NewStack(false)

	var group sync.WaitGroup

	//spam push
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				megaStack.Lock()
				ok := megaStack.Push(times)

				if ok == false {
					t.Error("insert failed " + string(times))
				}

				megaStack.Unlock()

				time.Sleep(time.Millisecond * 10)
			}
			group.Done()
		}()
	}

	//spam pop and IsEmpty
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				megaStack.Lock()
				if megaStack.HasElement() {
					_, ok := megaStack.Pop()

					if ok == false {
						t.Error("lock failed, hasElement() but Pop() failed")
					}
				}
				megaStack.Unlock()
				time.Sleep(time.Millisecond * 8)
			}
			group.Done()
		}()
	}

	group.Wait()
}

func TestConcurrencyAutoLock(t *testing.T) {

	megaStack := NewStack(true)
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
			if megaStack.String() == "" {
				t.Error("String() failed")
			}

			//spam for test coverage
			if megaStack.Len() < 0 {
				t.Error("Len() failed")
			}

			time.Sleep(time.Millisecond * 3)
		}()
	}

	//spam push
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				ok := megaStack.Push(times)

				if ok == false {
					t.Error("insert failed " + string(times))
				}

				time.Sleep(time.Millisecond * 10)
			}
			group.Done()
		}()
	}

	//spam pop and IsEmpty
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				if megaStack.HasElement() {
					_, ok := megaStack.Pop()

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

func TestStackBasicTypes(t *testing.T) {
	for _, r := range fakeTable {
		testStackFunctionality(t, r)
	}
}

func testStackFunctionality(t *testing.T, toPush []interface{}) {
	s := NewStack(false)

	for i, v := range toPush {
		ok := s.Push(v)

		if ok == false {
			t.Error("push failed ")
		}
		len := s.Len()

		if len != i+1 {
			t.Errorf("len failed, expected %v, got %v, for %v", i, len, toPush)
		}

		value, ok := s.Peek()

		if ok == false {
			t.Error("peek failed")
		}

		if value != v {
			t.Errorf("peek failed, expected %v, got %v ", v, value)
		}
	}

	for i := len(toPush) - 1; i >= 0; i-- {
		el, ok := s.Pop()

		if ok == false {
			t.Error("pop failed")
		}

		if el != toPush[i] {
			t.Errorf("pop failed, expected %v, got %v", toPush[i], el)
		}
	}

	if s.HasElement() {
		t.Errorf("stack is not empty after all Pop(), size=%v", s.Len())
	}
}

func TestInitPeekIsNil(t *testing.T) {
	s := NewStack(false)
	peek, ok := s.Peek()

	if ok {
		t.Error("peek should be false when used on an empty stack")
	}

	if peek != nil {
		t.Errorf("expected nil, got %v ", peek)
	}
}

func TestStackInitPopIsNil(t *testing.T) {
	s := NewStack(false)
	var pop, ok = s.Pop()

	if ok {
		t.Error("Pop should be false when used on an empty stack")
	}

	if pop != nil {
		t.Errorf("expected nil, got %v ", pop)
	}
}

func TestStackInitIsEmpty(t *testing.T) {
	helperInitIsEmpty("stack", NewStack(false), t)
}

func TestStackSkippingNewShouldPanic(t *testing.T) {
	//TODO Should Stack implement lazyInit() method (like list.go has) so this should never happen?
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the code did not panic")
		}
	}()
	s := Stack{}

	c := s.Len()

	fmt.Print(c)
}

func TestStackStringer(t *testing.T) {
	s := NewStack(false)

	v := s.String()
	if v != "Stack [0]" {
		t.Error("stringer was incorrect for 0 length" + v)
	}

	s.Push(1)
	s.Push(1)
	s.Push(1)

	v = s.String()
	if v != "Stack [3]" {
		t.Error("stringer was incorrect for 3 length" + v)
	}
}

func BenchmarkStackSync1000(b *testing.B) {
	benchStackSync(1000, b)
}

func BenchmarkStackSync100000(b *testing.B) {
	benchStackSync(100000, b)
}

func BenchmarkStackSync1000000(b *testing.B) {
	benchStackSync(1000000, b)
}

func benchStackSync(count int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := NewStack(false)

		for c := 0; c < count; c++ {
			if q.Push("a") == false {
				b.Error("s push failed")
			}
		}

		for c := 0; c < count; c++ {
			if _, ok := q.Pop(); ok == false {
				b.Error("s pop failed")
			}
		}
	}
}
