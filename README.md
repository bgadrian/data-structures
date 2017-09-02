# Data structures in Go [![Build Status](https://travis-ci.org/BTooLs/basic-data-and-algorithms.svg?branch=master)](https://travis-ci.org/BTooLs/basic-data-and-algorithms) [![codecov](https://codecov.io/gh/BTooLs/basic-data-and-algorithms/branch/master/graph/badge.svg)](https://codecov.io/gh/BTooLs/basic-data-and-algorithms) [![Go Report Card](https://goreportcard.com/badge/github.com/BTooLs/basic-data-and-algorithms)](https://goreportcard.com/report/github.com/BTooLs/basic-data-and-algorithms)
I am writing a collection of packages for different data structures in GO.
 
Why? To learn Go, practice basic structures and learning to code fast concurrent algorithms.

All the packages have 100+% test coverage, benchmark tests and godocs. Tested with go 1.9.

#### !! Warning This library wasn't used in production (yet). !!

## priorityqueue  [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/BTooLs/data-structures/priorityqueue)
A collection of performant, concurrent safe, complex abstract data structures used for priority queues.

### Hierarchical Queue [description](https://www.researchgate.net/figure/261191274_fig1_Figure-1-Simple-queue-a-and-hierarchical-queue-b) 
An **O(1)/O(1) priority queue (very fast)** implementation for small integers, that uses an assembly of N simple queues. It is optimized for large amount of data BUT with small value priorities ( <= 255 ). Can store any type of elements/values. 

## heap  [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/BTooLs/data-structures/heap)
A collection of basic abstract heap data structures.

### Implicit Heap [description](http://www.cs.princeton.edu/courses/archive/spr09/cos423/Lectures/i-heaps.pdf) [example](https://www.tutorialspoint.com/data_structures_algorithms/heap_data_structure.htm)
Dynamic Min & Max Implicit heaps.
Insert (push) and remove min/max (pop) have ```O(log n)``` complexity. The keys are ```int```and values can be any type ```interface{}```.


## linear [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/BTooLs/data-structures/linear)
A collection of **(not so performant)** simple linear data structres, that are not in the standard Go lib.

### Stack [description](https://www.tutorialspoint.com/data_structures_algorithms/stack_algorithm.htm)
Basic stack (FILO) using the builtin linked list, can store any type, concurrency safe, no size limit, implements Stringer.

### Queue [description](https://www.tutorialspoint.com/data_structures_algorithms/dsa_queue.htm) 
Basic queue (FIFO) using the builtin linked list, can store any type, concurrency safe (optional mutex), no size limit, implements Stringer.