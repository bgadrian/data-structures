package linear

import (
	"container/list"
	"sync"
)

//ListCommon common methods for all list type based
type ListCommon interface {
	Len() int
	HasElement() bool
	IsEmpty() bool
}

type common struct {
	data     *list.List
	autoLock bool
	sync.Mutex
}

//Len get the current length of the list. The complexity is O(1).
func (s *common) Len() int {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len()
}

//IsEmpty Returns true if the list is empty. Use this before Pop or Peek. Opposite of HasElement()
func (s *common) IsEmpty() bool {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len() == 0
}

//HasElement Returns true if the list is NOT empty. Use this before Pop or Peek. Opposite of IsEmpty()
func (s *common) HasElement() bool {
	if s.autoLock {
		s.Lock()
		defer s.Unlock()
	}
	return s.data.Len() > 0
}
