package priorityqueue

import "container/list"

//import "github.com/karalabe/cookiejar/collections/deque"

/*LadderQueue An O(1) amortized, Priority Queue Structure for Large-Scale Discrete Event Simulation

It an improved Sorted-discipline Calendar Queue (SCQ) and Lazy Queue structure,
Inspired from papers [Ladder queue: An O (1) priority queue structure for large-scale discrete event simulation](https://www.researchgate.net/publication/234787413_Ladder_queue_An_O_1_priority_queue_structure_for_large-scale_discrete_event_simulation)
We will use timestamps as Priorities, for a more broad scope

Basically, the structure consists of three tiers: a sorted linked list called
Bottom; the middle layer, called Ladder, consisting of several rungs of buckets
where each bucket may contain an unsorted linked list; and a simple unsorted
linked list called Top

The Ladder Queue [8] is a variant of the Calendar Queue
which is more suited for skewed distributions of the timestamps
of the events, thanks to the possibility of dynamically
splitting an individual bucket in sub-intervals (i.e. sublists
of records) if the number of elements associated with the
1 Source code available at https://github.com/HPDCS/NBCQ.
bucket exceeds a given threshold*/
type LadderQueue struct {
	top          lqTop
	ladder       lqLadder
	bot          lqBottom
	epochStarted bool
}

type lqNode struct {
	ts float64     //priority
	v  interface{} //value of the event
}

type lqTop struct {
	//Maximum timestamp of all events in Top. Its value is updated as events are enqueued into Top.
	maxTS float64
	//Minimum timestamp of all events in Top. Its value is updated as events are enqueued into Top.
	minTS float64
	//NTop Number of events in Top.
	n int
	//TopStart Minimum timestamp threshold of events which must be enqueued in Top
	start float64
	//unsorted list of events
	a *list.List
}

type lqLadder struct {
	//Bucketwidth [x] Bucketwidth of Rung [x].
	bucketWidth []float64
	//NBc Number of events in current dequeue bucket Bc.
	nbC int
	//NBucket[ j, k] Number of events in Bucket[k] of Rung [j].
	nBucket [][]int
	//NRung Number of rungs currently in active use.
	nRung int

	//RCur [x] Starting timestamp threshold of the first valid bucket in Rung [x]
	//which subsequent dequeue operations will start.
	//Minimum timestamp threshold of events which can be enqueued in Rung [x].
	rCur []float64

	//RStart[x] Starting timestamp threshold of the first bucket in Rung [x].
	//Used for calculating the bucket-index when inserting an event in Rung [x] of Ladder. See Eq. (2).
	rStart []float64

	//THRES If the number of events in a bucket or Bottom exceeds this threshold, then a spawning action would be initiated
	thres int

	//[rungs][buckets] with events
	a [][]*list.List
}

type lqBottom struct {
	//NBot Number of events in Bottom.
	n int

	//sorted list of events, instead of linked list we use a fast circular structure
	a *list.List
}

//NewLadderQueue Generates a new empty LadderQueue
func NewLadderQueue() *LadderQueue {
	// buckets := 10
	return &LadderQueue{
		top:    lqTop{a: list.New()},
		ladder: lqLadder{thres: 50},
		bot:    lqBottom{a: list.New()},
	}
}

func (q *LadderQueue) pushToTop(timestamp float64, value interface{}) {
	q.top.a.PushFront(lqNode{timestamp, value})
	q.top.n++

	if timestamp > q.top.maxTS {
		q.top.maxTS = timestamp
	}

	if timestamp < q.top.minTS {
		q.top.minTS = timestamp
	}
}

func (q *LadderQueue) pushToBottom(timestamp float64, value interface{}) {

	for current := q.bot.a.Back(); ; {
		node := current.Value.(lqNode)

		if timestamp < node.ts {
			q.bot.a.InsertBefore(lqNode{timestamp, value}, current)
			break
		}

		//we reched the end
		if current.Next() == nil {
			q.bot.a.InsertAfter(lqNode{timestamp, value}, current)
			break
		}

		current = current.Next()
	}
	q.bot.n++

	//TODO
	if q.bot.n > q.ladder.thres {
		/*  rung spawning process somewhat similar to that found in the dequeue
		operation. To describe this process briefly, assume that the number of
		rungs currently in use (NRung) is i. A new rung, Rung [i+1], is created with
		Bucketwidth [i+1] set using Eq. (3). Events that belonged to the Bottom list will
		then be redistributed into Rung [i+1]. The previous Bc in Rung [i] then changes
		to Bsp. A new Bc is identified in Rung [i+1] and events in Bc will then be sorted
		to form a new Bottom list */

		newRungIndex := q.ladder.nRung
		q.ladder.a[newRungIndex] = make([]*list.List, 100)
		q.eq3SetNextBucketWidth(newRungIndex - 1)

		q.transferEventsToRung(q.bot.a, q.ladder.a[newRungIndex], newRungIndex)

		//TODO
		// The previous Bc in Rung [i] then changes to Bsp.
		//A new Bc is identified in Rung [i+1] and events in Bc will then be sorted to form a new Bottom list
	}
}

//Enqueue ...
func (q *LadderQueue) Enqueue(timestamp float64, value interface{}) {
	// If TS  TopStart, the event is appended at Top.
	if q.epochStarted == false || timestamp > q.top.start {
		q.pushToTop(timestamp, value)
		return
	}

	// TS is compared with thresholds RCur [1], RCur [2], ... , RCur [NRung]
	// to determine which rung the event should be in. Once TS  RCur [i], the event
	// is inserted into Rung [i].
	for i := 0; i < q.ladder.nRung; i++ {
		if timestamp >= q.ladder.rCur[i] {
			bucketIndex := q.eq2BucketIndex(timestamp, i)
			q.ladder.a[i][bucketIndex].PushFront(lqNode{timestamp, value})
			//TODo increment counters and stuff
		}
	}

	//fallback, we put it to bottom
	q.pushToBottom(timestamp, value)
}

func (q *LadderQueue) eq2BucketIndex(timestamp float64, rungIndex int) int {
	return int((timestamp - q.ladder.rStart[rungIndex]) / q.ladder.bucketWidth[rungIndex])
}

func (q *LadderQueue) eq3SetNextBucketWidth(bucketIndex int) {
	// Bucketwidth [i + 1] = Bucketwidth [i] /THRES
	q.ladder.bucketWidth[bucketIndex+1] = q.ladder.bucketWidth[bucketIndex] / float64(q.ladder.thres)
}

//transferTop transfer the top events into the first RUNG[1]
func (q *LadderQueue) transferTop() {
	bucketWidth := (q.top.maxTS - q.top.minTS) / float64(q.top.n)

	rungIndex := 0

	// RStart[1] and RCur [1] are set = MinTS, and TopStart is set = MaxTS + Bucketwidth [1].
	q.ladder.bucketWidth[rungIndex] = bucketWidth
	q.ladder.rStart[rungIndex] = q.top.minTS
	q.ladder.rCur[rungIndex] = q.top.minTS
	q.top.start = q.top.maxTS + bucketWidth

	//transfer the top events in Bucket k = (TS − RStart[i]) /Bucketwidth [i]
	q.transferEventsToRung(q.top.a, q.ladder.a[rungIndex], rungIndex)

	//ss
}

func (q *LadderQueue) transferEventsToRung(from *list.List, toRung []*list.List, toRungIndex int) {
	for e := from.Front(); e != nil; e = e.Next() {
		oneNode := e.Value.(lqNode)
		//distribute in buckets, according to eq2
		bucketIndex := q.eq2BucketIndex(oneNode.ts, toRungIndex)
		toRung[bucketIndex].PushFrontList(q.top.a)
		from.Remove(e.Prev())
	}
}

//Dequeue ...
func (q *LadderQueue) Dequeue() (v interface{}, err error) {

	if q.epochStarted == false {
		q.epochStarted = true
		//when the first Deq happens ... TODO
		q.transferTop()

		//traversal process to find the next non empty bucket

	}

	//usually, first bottom item is the highest priority
	e := q.bot.a.Back()
	q.bot.a.Remove(e)
	node := e.Value.(lqNode)
	return node.v, nil
}
