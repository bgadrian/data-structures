
## priorityqueue  [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/BTooLs/data-structures/priorityqueue)
A collection of performant, concurrent safe, complex abstract data structures used for priority queues.

### Hierarchical Queue [description](https://www.researchgate.net/figure/261191274_fig1_Figure-1-Simple-queue-a-and-hierarchical-queue-b) 
An **O(1)/O(1) priority queue** implementation for small integers, that uses an assembly of N simple queues.

It is optimized for large amount of data BUT with small value priorities ( <= 255 ). Can store any type of elements/values. 

**Priority: 0 (highest) - n (lowest)**

For best performance **Enqueue ALL the elements before starting to Dequeue**.
The downsides:
- the instance cannot be reused because ...
- once a priority queue is empty and removed, it cannot be recreated
#### Hierarchical Queue usages 
* image/video processing
* networking (routing)
* anywhere you have a small range of priorities/channels.

#### Hierarchical Queue implementation:

![HQ example](https://www.researchgate.net/profile/Serge_Beucher/publication/261191274/figure/fig1/AS:296718022266884@1447754497479/Figure-1-Simple-queue-a-and-hierarchical-queue-b.png)

(a) - normal queue, (b) - list of queues

It is an array of buckets. The key is the priority and the bucket is a queue. Queues are ~~linked lists~~ [dynamically growing circular slice of blocks](https://github.com/karalabe/cookiejar/tree/master/collections/queue), the advantage is that no memory preallocation is needed and the queue/dequeue is O(1).
We dequeue from highest priority (0) until it's bucket (queue) is empty and we remove it. We move to the next priority (1) and so on until we deplete the structure.

The keys are ```uint8```and values can be any type ```interface{}```.

Inspired by papers:
- [*Revisiting Priority Queues for Image Analysis, Cris L. Luengo Hendriks*](http://www.cb.uu.se/~cris/Documents/Luengo2010a_preprint.pdf)
- [*Hierarchical Queues: general description and implementation in MAMBA Image library, Nicolas Beucher and Serge Beucher*](http://cmm.ensmp.fr/~beucher/publi/HQ_algo_desc.pdf)

#### Hierarchical Queue benchmarks
This syncronous tests were done to demonstrate that Enqueue/Dequeue is **O(1)** regardless of the priority queue size. A queue is filled with N elements and equally distributed priorities. The data stored is 1 character. 

Each pass consists of : 1 enqueue with increasing priority (0,1,2,3...255,0,1...) and 1 dequeue. K is the priority lowest value (0 - K). 

```bash
go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/btools/basic-data-and-algorithms/src/data/lists
```

|K = 50 | | |
|---|:---:|:---:|
|N = 1000            |20000000               |24.4 ns/op|
|N = 100000          |10000000               |24.4 ns/op|
|N = 1000000         |10000000               |24.4 ns/op|
|N = 10000000        |10000000               |26.6 ns/op|
|N = 100000000       |10000000               |33.4 ns/op|

*Previous implementation used list.List linked lists, they were replaced with a queue 10x faster.*
