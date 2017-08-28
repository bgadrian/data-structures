package lists

import (
	"container/list"
	"errors"
	"sync"
)

//HierarchicalQueue A priority ordered queues.
type HierarchicalQueue struct {
	autoLock        bool
	queues          []*list.List
	maxPriority     uint8
	currentPriority uint8
	sync.Mutex
}

//NewHierarchicalQueue Generates ...
func NewHierarchicalQueue(maxPriority uint8, autoMutexLock bool) *HierarchicalQueue {
	return &HierarchicalQueue{
		queues:          make([]*list.List, maxPriority),
		maxPriority:     maxPriority,
		currentPriority: 0,
		autoLock:        autoMutexLock,
	}
}

//Enqueue ss
func (l *HierarchicalQueue) Enqueue(value interface{}, priority uint8) (err error) {
	//if we already began to take elements out and empty queues
	//we add all the new elements in the current queue
	//if their priority is smaller than the current one
	if priority < l.currentPriority {
		priority = l.currentPriority
	}

	if priority > l.maxPriority {
		return errors.New("priority is bigger than max priority")
	}

	if l.queues[priority] == nil {
		l.queues[priority] = list.New()
	}

	element := l.queues[priority].PushFront(value)

	if element == nil {
		return errors.New("cannot insert to list, internal error")
	}
	return nil
}
