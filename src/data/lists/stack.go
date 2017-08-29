package lists

import (
	"container/list"
	"strconv"
)

//Stack FILO list, uses Linked lists
//Please use lists.NewStack()
type Stack struct {
	common
}

//NewStack generates a new stack
func NewStack(autoMutexLock bool) (l *Stack) {
	l = &Stack{}
	l.data = list.New()
	l.autoLock = autoMutexLock
	return
}

//Push (storing) an element on the stack.
func (s *Stack) Push(item interface{}) (ok bool) {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}

	s.data.PushBack(item)

	return true
}

//Pop Removing (accessing) an element from the stack.
func (s *Stack) Pop() (item interface{}, ok bool) {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}

	first := s.data.Back()
	if first == nil {
		//the list is empty
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

	first := s.data.Back()

	return first.Value, true
}

//String returns a string representation of the list
func (s *Stack) String() string {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return "Stack [" + strconv.Itoa(s.data.Len()) + "]"
}
