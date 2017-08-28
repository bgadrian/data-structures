/*Package lists contains a simple FIFO list implementation, no max size, generic use, store data in linked lists.

Scenario 1:
Faster stack, but not safe for concurrency.
var listNotSafe := Stack.NewStack(false) //or queue

Scenario 2:
If you use goroutines create one using
var listSafe := Stack.NewStack(true) //or queue
Most common error is "stack was empty", check Stack.Empty() or ignore it in highly-concurrent funcs.
Because the state may change between the HasElement() call and Pop/Peek.

Scenario 3:
Manual lock the struct, 100% reability, prune to mistakes/bugs
var listNotSafe := Stack.NewStack(false) //or queue
listNotSafe.Lock()
//do stuff with the list
listNotSafe.Unlock()*/
package lists

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
