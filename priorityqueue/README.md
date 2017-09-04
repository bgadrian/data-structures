
## priorityqueue  [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/BTooLs/data-structures/priorityqueue)
A collection of high performance, concurrent safe, complex abstract data structures used for priority queues.

*Priority queue is an abstract data type which is like a regular queue or stack data structure, but where additionally each element has a "priority" associated with it. In a priority queue, an element with high priority is served before an element with low priority.*

## Hierarchical Queue [description](https://www.researchgate.net/figure/261191274_fig1_Figure-1-Simple-queue-a-and-hierarchical-queue-b) 
An **O(1)/O(1 + K) priority queue** implementation for small integers, that uses an assembly of N simple queues.

It is optimized for large amount of data BUT with small value priorities ( <= 255 ). Can store any type of elements/values. 

#### Hierarchical Queue usages 
* image/video processing
* networking (routing)
* anywhere you have a small range of priorities/channels.

**The original algorithm has O(1) dequeue complexity, but that adds a very strict limitation: a queue that has been empty it is removed and cannot be recreated. This means the instance cannot be reused once the dequeue finish, and inserting elements during the dequeue decrease the performance. I decided that limits the scope of the algorithm too much and has a low chance to be used in the real world, so I removed the complexity.**

#### Hierarchical Queue implementation:

![HQ example](https://www.researchgate.net/profile/Serge_Beucher/publication/261191274/figure/fig1/AS:296718022266884@1447754497479/Figure-1-Simple-queue-a-and-hierarchical-queue-b.png)

(a) - normal queue, (b) - list of queues

It is an array of buckets. The key is the priority and the bucket is a queue. Queues are ~~linked lists~~ [dynamically growing circular slice of blocks](https://github.com/karalabe/cookiejar/tree/master/collections/queue), the advantage is that no memory preallocation is needed and the queue/dequeue is O(1).

The keys are ```uint8```and values can be any type ```interface{}```.

**Priority: 0 (highest) - n (lowest)**

Inspired by papers:
- [*Revisiting Priority Queues for Image Analysis, Cris L. Luengo Hendriks*](http://www.cb.uu.se/~cris/Documents/Luengo2010a_preprint.pdf)
- [*Hierarchical Queues: general description and implementation in MAMBA Image library, Nicolas Beucher and Serge Beucher*](http://cmm.ensmp.fr/~beucher/publi/HQ_algo_desc.pdf)

#### [Hierarchical Queue benchmarks](benchmark.log)
This [syncronous tests](benchmark.log) were done to demonstrate that Enqueue/Dequeue is **O(1)** regardless of the priority queue size.

The avg value is ```50 ns/op``` using random distributed priorities and ``` 25 ns/op``` using linear priorities on an I7 2.7GHz 64bit windows.

*Previous implementation used list.List linked lists, they were replaced with a queue 10x faster.*

## Hierarhical Heap
It is a version of the Hierarchical Queue structure, adding some complexity (O(log n/k)) but removing it's limitations, proposed by [Cris L. Luengo Hendriks paper](http://www.cb.uu.se/~cris/Documents/Luengo2010a_preprint.pdf)

Unlike HQ that has 1 bucket for each priority value, HH priorities are grouped in K buckets.

**For the best performance benchmark the Enq/Deq functions with your own data, and tweak the buckets,minP,maxP parameters!**


#### Hierarchical Heap implementation

The keys are ```int```and values can be any type ```interface{}```.

**Priority: 0 (highest) - n (lowest)**

The HH is an array of buckets, each bucket contains 0-n number of priority values. The buckets are spread based on the priority with a linear formula:

```bucketIndex = (priority - minPriority)/(maxPriority - minPriority) * bucketsCount```

Buckets are Min Implicit Heaps (special binary tree) that stores all values for a specific range of priorities.
Example: pMin=0, pMax=100 (0-100 priorities), K=15 (buckets)
Enqueue("a", 21) will add "a" to bucket i, where
i = (21 − pmin) / (pmax − pmin) * K = 3

*By dividing the N elements in the queue over k buckets,
the average size of each heap is reduced by a factor k, and
the enqueueing and dequeueing operations thus are of order
O(logN/k) if the buckets are chosen correctly. Of course, the actual algorithmic complexity depends strongly on the distribution of priority values. Calculating the bucket index adds a constant time to the enqueue operation that needs to be amortised by the O(logk) reduction in time to enqueue the element in the implicit heap. It is therefore necessary to choose k large enough.*

#### [Hierarchical Heap benchmarks](benchmark.log)

With the right bucket/priority diversity it can aproach the HQ performance by a factor of 0.5.

