/*Package stack contains a simple FIFO list implementation, no max size, generic use, store data in linked lists.

 Faster stack, but not safe for concurrency.
 var StackNotSafe := Stack.New(false)

If you use goroutines create one using
var StackNotSafe := Stack.New(true)

Most common error is "stack was emtpy", check Stack.Empty() or ignore it in highly-concurrent funcs.
*/
package stack

import (
	"container/list"
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
func (s *Stack) Push(element interface{}) bool {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}

	listElement := s.data.PushBack(element)

	if listElement == nil {
		return false
	}

	return true
}

//Pop Removing (accessing) an element from the stack.
func (s *Stack) Pop() (element interface{}, ok bool) {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	if s.data.Len() == 0 {
		return nil, false
	}

	first := s.data.Back()
	if first == nil {
		return nil, false
	}
	s.data.Remove(first)

	return first.Value, true
}

//Peek get the top data element of the stack, without removing it.
func (s *Stack) Peek() (interface{}, bool) {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}

	if s.data.Len() == 0 {
		return nil, false
	}

	return s.data.Back().Value, true
}

//Len get the current length of the stack. The complexity is O(1).
func (s *Stack) Len() int {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len()
}

//IsEmpty Returns true if the stack is empty. Use this before Pop or Peek. Opposite of HasElement()
func (s *Stack) IsEmpty() bool {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len() == 0
}

//HasElement Returns true if the stack is NOT empty. Use this before Pop or Peek. Opposite of IsEmpty()
func (s *Stack) HasElement() bool {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len() > 0
}

func (s *Stack) String() string {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	return "Stack len " + string(s.data.Len())
}
