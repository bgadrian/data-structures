package priorityqueue

import (
	"errors"
	"sync"

	"github.com/karalabe/cookiejar/collections/deque"
)

//HierarchicalQueue An O(1)/O(1)* priority queue implementation for small integers
//See the README for more info.
type HierarchicalQueue struct {
	autoLock bool
	q        []*deque.Deque
	sync.Mutex
	count   int
	lowestP int
	smP     uint8 //smallest queue that has elements,cache for optimization
}

//NewHierarchicalQueue Generates a new HQ
func NewHierarchicalQueue(lowestPriority uint8, autoMutexLock bool) *HierarchicalQueue {
	return &HierarchicalQueue{
		q:        make([]*deque.Deque, uint16(lowestPriority)+1),
		lowestP:  int(lowestPriority),
		autoLock: autoMutexLock,
	}
}

//Enqueue Add a new element with a priority (0-highest priority, n-lowest)
func (l *HierarchicalQueue) Enqueue(value interface{}, priority uint8) (err error) {
	if l.autoLock {
		l.Lock()
		defer l.Unlock()
	}

	if priority > uint8(l.lowestP) {
		priority = uint8(l.lowestP)
	}

	if l.q[priority] == nil {
		l.q[priority] = deque.New()
	}

	//TODO learn how to do this/break dequeue
	// if l.q[priority] == nil {
	// 	return errors.New("cannot create a queue deque")
	// }

	l.q[priority].PushRight(value)

	l.count++

	if priority < l.smP {
		l.smP = priority
	}

	return nil
}

//Dequeue Return the highest priority value (0-highest priority, n-lowest)
func (l *HierarchicalQueue) Dequeue() (v interface{}, err error) {
	if l.autoLock {
		l.Lock()
		defer l.Unlock()
	}

	for i := int(l.smP); i <= l.lowestP; i++ {
		if l.q[i] == nil || l.q[i].Empty() {
			continue
		}

		l.smP = uint8(i)
		v = l.q[i].PopLeft()
		l.count--
		return
	}

	err = errors.New("the queue is empty")
	return
}

//Len Return the count of all values from all priorities
func (l *HierarchicalQueue) Len() int {
	if l.autoLock {
		l.Lock()
		defer l.Unlock()
	}

	return l.count
}

//LenPriority Returns the count of all values from a specific priority queue
func (l *HierarchicalQueue) LenPriority(priority uint8) int {
	if l.autoLock {
		l.Lock()
		defer l.Unlock()
	}

	if l.q[priority] == nil {
		return 0
	}

	return l.q[priority].Size()
}
