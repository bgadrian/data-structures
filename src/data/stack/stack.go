package stack

import (
	"container/list"
	"errors"
	"sync"
)

//Stack a dynamic FIFO list, uses Linked lists for memory efficiency.
type Stack struct {
	data *list.List
	safe bool
	sync.Mutex
}

//New generates a new stack
func New(concurrencySafe bool) *Stack {
	n := &Stack{}
	n.data = list.New()
	n.safe = concurrencySafe
	return n
}

//Push (storing) an element on the stack.
func (s *Stack) Push(element interface{}) error {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}

	listElement := s.data.PushBack(element)

	if listElement == nil {
		return errors.New("insert failed")
	}

	return nil
}

//Pop Removing (accessing) an element from the stack.
func (s *Stack) Pop() (element interface{}, err error) {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	if s.data.Len() == 0 {
		return nil, errors.New("The stack was empty")
	}

	first := s.data.Back()
	s.data.Remove(first)

	return first.Value, nil
}

//Peek get the top data element of the stack, without removing it.
func (s *Stack) Peek() (interface{}, error) {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}

	if s.data.Len() == 0 {
		return nil, errors.New("The stack was empty")
	}

	return s.data.Back().Value, nil
}

//Len get the current length of the stack. The complexity is O(1).
func (s *Stack) Len() int {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len()
}

//IsEmpty Returns true if the stack is empty. Use this before Pop or Peek
func (s *Stack) IsEmpty() bool {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len() == 0
}

func (s *Stack) String() string {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	return "Stack len " + string(s.data.Len())
}
