package lists

import (
	"errors"
	"sync"
)

//HierarchicalQueueSlices An O(1)/O(1)* priority queue implementation for small integers using slices.
//See the README for more info.
type HierarchicalQueueSlices struct {
	autoLock bool
	q        [][]interface{} //a fixed of array of slices
	lowestP  uint8
	highestP uint8
	sync.Mutex
	depleted bool
}

//NewHierarchicalQueueSlices Generates a new HQ
func NewHierarchicalQueueSlices(lowestPriority uint8, autoMutexLock bool) *HierarchicalQueueSlices {
	return &HierarchicalQueueSlices{
		q:        make([][]interface{}, uint16(lowestPriority)+1),
		lowestP:  lowestPriority,
		highestP: 0, //advances to lowestP to empty all queues
		autoLock: autoMutexLock,
		depleted: false,
	}
}

//Enqueue Add a new element with a priority (0-highest priority, n-lowest)
func (l *HierarchicalQueueSlices) Enqueue(value interface{}, priority uint8) (err error) {
	if l.autoLock {
		l.Lock()
		defer l.Unlock()
	}

	if l.depleted {
		return errors.New("depleted queue") //nothing to do
	}

	if priority > l.lowestP {
		return errors.New("priority is bigger than max priority")
	}

	// if l.q[priority] == nil {
	// 	l.q[priority] = []interface{}
	// }

	//special exception when we already began to take elements out and empty queues
	//we add all the new elements in the current queue
	//if their priority is smaller than the current one
	//The HQ rule is "when a queue is empty and removed, it cannot be recreated"
	if priority < l.highestP {
		l.q[l.highestP] = append([]interface{}{value}, l.q[l.highestP]...)
	} else {
		l.q[priority] = append(l.q[priority], value)
	}

	return nil
}

//removeEmptyQ Advance to the next queue.
//You may experience some performance hickups if you have sparse priority values ex: 0,1,2,3,250,251 ..
func (l *HierarchicalQueueSlices) removeEmptyQ() {
	for {
		//we found a non empty queue, do NOT advance
		if l.q[l.highestP] != nil && len(l.q[l.highestP]) > 0 {
			break
		}

		//remove the empty queue
		l.q[l.highestP] = nil

		if l.highestP == l.lowestP {
			l.depleted = true
			break
		}

		l.highestP++
	}
}

//Dequeue Return the highest priority value (0-highest priority, n-lowest)
//Recommended: start to Dequeue AFTER you Enqueue ALL the elements
func (l *HierarchicalQueueSlices) Dequeue() (interface{}, error) {
	if l.autoLock {
		l.Lock()
		defer l.Unlock()
	}

	if l.depleted {
		return nil, errors.New("depleted queue") //nothing to do
	}

	//this covers the case when you start to Deq before Enq
	if l.q[l.highestP] == nil || len(l.q[l.highestP]) == 0 {
		l.removeEmptyQ()

		if l.depleted {
			return nil, errors.New("depleted queue") //nothing to do
		}
	}

	element := l.q[l.highestP][0]
	l.q[l.highestP] = l.q[l.highestP][1:]

	//make sure next time we have something to dequeue
	l.removeEmptyQ()

	return element, nil
}

//IsDepleted If all the queues are empty and removed this instance cannot be used anymore
func (l *HierarchicalQueueSlices) IsDepleted() bool {
	if l.autoLock {
		l.Lock()
		defer l.Unlock()
	}
	return l.depleted
}
