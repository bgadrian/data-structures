## heap  [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/bgadrian/data-structures/heap)
A collection of basic abstract heap data structures.

### Implicit Heap [description](http://www.cs.princeton.edu/courses/archive/spr09/cos423/Lectures/i-heaps.pdf) [example](https://www.tutorialspoint.com/data_structures_algorithms/heap_data_structure.htm)
Dynamic Min & Max Implicit heaps.

An implicit heap is an implementation of a heap consisting of a complete binary tree whose
nodes contain the heap items, one node per item.

Insert example:
![insert gif](https://www.tutorialspoint.com/data_structures_algorithms/images/max_heap_animation.gif)

#### Implicit Heap Implementation
Insert (push) and remove min/max (pop) have ```O(log n)``` complexity. The size of the slices (array) is dynamic, with a minimum of **8** elements and doubles each time is full, and shrink to half each time it's ```N < size/4```.

The keys are ```int```and values can be any type ```interface{}```.

**For best perfomance use a small non sparsed Key value distribution. (100-300 incremental values).** 