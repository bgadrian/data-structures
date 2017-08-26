package stack

import "testing"
import "time"

import "sync"

func TestConcurrency(t *testing.T) {

	megaStack := New(true)
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
				if megaStack.IsEmpty() == false {
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

func TestBasicTypes(t *testing.T) {
	type Dummy struct {
		a int
	}

	a, b, c := 1, 1, 1

	table := [][]interface{}{
		{1.0, 2.2, 3.14},
		{-1000, 0, 1000},
		{"str1", "str2"},
		{true, false, true},
		{Dummy{1}, Dummy{2}},
		{1, true, "str"},
		{nil},
		{nil, nil},
		{nil, nil, nil, nil, nil},
		{&a, &b, &c},
	}

	for _, r := range table {
		one(t, r)
	}
}

func one(t *testing.T, toPush []interface{}) {
	s := New(false)

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

	if s.IsEmpty() == false {
		t.Errorf("stack is not empty after all Pop(), size=%v", s.Len())
	}
}
func TestInitPeekIsNil(t *testing.T) {
	s := New(false)
	peek, ok := s.Peek()

	if ok {
		t.Error("peek should be false when used on an empty stack")
	}

	if peek != nil {
		t.Errorf("expected nil, got %v ", peek)
	}
}

func TestInitPopIsNil(t *testing.T) {
	s := New(false)
	var pop, ok = s.Pop()

	if ok {
		t.Error("Pop should be false when used on an empty stack")
	}

	if pop != nil {
		t.Errorf("expected nil, got %v ", pop)
	}
}

func TestInitIsEmpty(t *testing.T) {
	s := New(false)

	if s.IsEmpty() == false {
		t.Errorf("Stack is not empty after created (isEmpty)")
	}

	if s.HasElement() {
		t.Errorf("Stack is not empty after created (hasElement)")
	}
}
