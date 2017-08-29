package lists

import (
	"container/list"
	"errors"
	"sync"
)

//HierarchicalQueue An O(1)/O(1)* priority queue implementation for small integers
//See the README for more info.
type HierarchicalQueue struct {
	autoLock bool
	q        []*list.List
	lowestP  uint8
	highestP uint8
	sync.Mutex
}

//NewHierarchicalQueue Generates a new HQ
func NewHierarchicalQueue(lowestPriority uint8, autoMutexLock bool) *HierarchicalQueue {
	return &HierarchicalQueue{
		q:        make([]*list.List, lowestPriority+1),
		lowestP:  lowestPriority,
		highestP: 0, //advances to lowestP to empty all queues
		autoLock: autoMutexLock,
	}
}

//Enqueue Add a new element with a priority (0-highest priority, n-lowest)
func (l *HierarchicalQueue) Enqueue(value interface{}, priority uint8) (err error) {
	if l.autoLock {
		l.Lock()
		defer l.Unlock()
	}

	if priority > l.lowestP {
		return errors.New("priority is bigger than max priority")
	}

	if l.q[priority] == nil {
		l.q[priority] = list.New()
	}

	//special exception when we already began to take elements out and empty queues
	//we add all the new elements in the current queue
	//if their priority is smaller than the current one
	//The HQ rule is "when a queue is empty and removed, it cannot be recreated"
	if priority < l.highestP {
		l.q[l.highestP].PushBack(value)
	} else {
		l.q[priority].PushFront(value)
	}

	return nil
}

//removeEmptyQ Advance to the next queue.
//You may experience some performance hickups if you have sparse priority values ex: 0,1,2,3,250,251 ..
func (l *HierarchicalQueue) removeEmptyQ() {
	for {
		//we found a non empty queue, do NOT advance
		if l.q[l.highestP] != nil && l.q[l.highestP].Len() > 0 {
			break
		}

		//remove the empty queue
		l.q[l.highestP] = nil
		l.highestP++

		if l.highestP > l.lowestP {
			break
		}
	}
}

//Dequeue Return the highest priority value (0-highest priority, n-lowest)
//Recommended: start to Dequeue AFTER you Enqueue ALL the elements
func (l *HierarchicalQueue) Dequeue() (interface{}, error) {
	if l.autoLock {
		l.Lock()
		defer l.Unlock()
	}

	if l.highestP > l.lowestP {
		return nil, errors.New("depleted queue") //nothing to do
	}

	//this covers the case when you start to Deq before Enq
	if l.q[l.highestP] == nil || l.q[l.highestP].Len() == 0 {
		l.removeEmptyQ()

		if l.highestP > l.lowestP {
			return nil, errors.New("depleted queue") //nothing to do
		}
	}

	element := l.q[l.highestP].Back()
	l.q[l.highestP].Remove(element)

	//make sure next time we have something to dequeue
	l.removeEmptyQ()

	return element.Value, nil
}

//IsDepleted If all the queues are empty and removed this instance cannot be used anymore
func (l *HierarchicalQueue) IsDepleted() bool {
	if l.autoLock {
		l.Lock()
		defer l.Unlock()
	}
	return l.highestP > l.lowestP
}
