package stack

import "testing"
import "time"
import "fmt"
import "sync"

func TestConcurrency(t *testing.T) {

	megaStack := New(true)
	var group sync.WaitGroup

	//spam peek
	for i := 0; i <= 10; i++ {
		go func() {
			if megaStack.IsEmpty() == false {
				value, err := megaStack.Peek()

				if err != nil {
					//we ignore the stack was empty, because of the fast times this will happen, 1 stack vs lots of workers
					//fmt.Printf(err.Error())
				}
				fmt.Println(value)
			}
			time.Sleep(time.Millisecond * 3)
		}()
	}

	//spam push
	for i := 0; i <= 100; i++ {
		group.Add(1)
		go func() {
			for times := 0; times < 200; times++ {
				err := megaStack.Push(times)

				if err != nil {
					t.Error(err)
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
					_, err := megaStack.Pop()

					if err != nil {
						//we ignore the stack was empty, because of the fast times this will happen, 1 stack vs lots of workers
						//fmt.Printf(err.Error())
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
		err := s.Push(v)

		if err != nil {
			t.Error(err)
		}
		len := s.Len()

		if len != i+1 {
			t.Errorf("Len expected %v, got %v, for %v", i, len, toPush)
		}

		peek, err := s.Peek()

		if err != nil {
			t.Error(err)
		}

		if peek != v {
			t.Errorf("Peek failed, expected %v, got %v ", v, peek)
		}
	}

	for i := len(toPush) - 1; i >= 0; i-- {
		el, err := s.Pop()

		if err != nil {
			t.Error(err)
		}

		if el != toPush[i] {
			t.Errorf("Pop failed, expected %v, got %v", toPush[i], el)
		}
	}

	if s.IsEmpty() == false {
		t.Errorf("Stack is not empty after all Pop, size=%v", s.Len())
	}
}
func TestInitPeekIsNil(t *testing.T) {
	s := New(false)
	peek, err := s.Peek()

	if err == nil {
		t.Error("Peek should return an error when used on an empty stack")
	}

	if peek != nil {
		t.Errorf("expected nil, got %v ", peek)
	}
}

func TestInitPopIsNil(t *testing.T) {
	s := New(false)
	var pop, err = s.Pop()

	if err == nil {
		t.Error("Pop should return an error when used on an empty stack")
	}

	if pop != nil {
		t.Errorf("expected nil, got %v ", pop)
	}
}

func TestInitIsEmpty(t *testing.T) {
	s := New(false)

	if s.IsEmpty() == false {
		t.Errorf("Stack is not empty after created")
	}
}
