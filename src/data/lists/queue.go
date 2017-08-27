package lists

import (
	"container/list"
	"strconv"
)

//Queue FIFO list, uses Linked lists
//Please use lists.NewQueue()
type Queue struct {
	common
}

//NewQueue generates a new queue
func NewQueue(autoMutexLock bool) (l *Queue) {
	l = &Queue{}
	l.data = list.New()
	l.autoLock = autoMutexLock
	return
}

//Enqueue (storing) an element on the queue.
func (s *Queue) Enqueue(item interface{}) (ok bool) {
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

//Dequeue Removing (accessing) an element from the queue.
func (s *Queue) Dequeue() (item interface{}, ok bool) {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	if s.data.Len() == 0 {
		return nil, false
	}

	last := s.data.Front()
	if last == nil {
		//don't know how this can happen, just being defensive
		return nil, false
	}
	s.data.Remove(last)

	return last.Value, true
}

//Peek get the top data element of the stack, without removing it.
func (s *Queue) Peek() (item interface{}, ok bool) {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}

	if s.data.Len() == 0 {
		return nil, false
	}

	first := s.data.Front()
	if first == nil {
		//don't know how this can happen, just being defensive
		return nil, false
	}

	return first.Value, true
}

//String returns a string representation of the list
func (s *Queue) String() string {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return "Queue [" + strconv.Itoa(s.data.Len()) + "]"
}
