/*Package stack contains a simple FIFO list implementation, no max size, generic use, store data in linked lists.

Scenario 1:
Faster stack, but not safe for concurrency.
var StackNotSafe := Stack.New(false)

Scenario 2:
If you use goroutines create one using
var StackNotSafe := Stack.New(true)
Most common error is "stack was emtpy", check Stack.Empty() or ignore it in highly-concurrent funcs.
Because the state may change between the HasElement() call and Pop/Peek.

Scenario 3:
Manual lock the struct, 100% reability, prune to mistakes/bugs
var Stack := Stack.New(false)
Stack.Lock()
//do stuff with Stack
Stack.Unlock()

*/
package stack

import (
	"container/list"
	"strconv"
	"sync"
)

//Stack a dynamic FIFO list, uses Linked lists for memory efficiency.
type Stack struct {
	data     *list.List
	autoLock bool
	sync.Mutex
}

//New generates a new stack
func New(autoMutexLock bool) *Stack {
	n := &Stack{}
	n.data = list.New()
	n.autoLock = autoMutexLock
	return n
}

//Push (storing) an element on the stack.
func (s *Stack) Push(item interface{}) (ok bool) {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}

	element := s.data.PushBack(item)
	if element == nil {
		//don't know how this can happen, just being defensive
		return false
	}

	return true
}

//Pop Removing (accessing) an element from the stack.
func (s *Stack) Pop() (item interface{}, ok bool) {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	if s.data.Len() == 0 {
		return nil, false
	}

	first := s.data.Back()
	if first == nil {
		//don't know how this can happen, just being defensive
		return nil, false
	}
	s.data.Remove(first)

	return first.Value, true
}

//Peek get the top data element of the stack, without removing it.
func (s *Stack) Peek() (item interface{}, ok bool) {
	if s.autoLock {
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
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len()
}

//IsEmpty Returns true if the stack is empty. Use this before Pop or Peek. Opposite of HasElement()
func (s *Stack) IsEmpty() bool {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len() == 0
}

//HasElement Returns true if the stack is NOT empty. Use this before Pop or Peek. Opposite of IsEmpty()
func (s *Stack) HasElement() bool {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len() > 0
}

func (s *Stack) String() string {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return "Stack [" + strconv.Itoa(s.data.Len()) + "]"
}
