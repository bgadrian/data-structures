# Basic data structures and algorithms in Go [![Build Status](https://travis-ci.org/BTooLs/basic-data-and-algorithms.svg?branch=master)](https://travis-ci.org/BTooLs/basic-data-and-algorithms)
Learning Go and TDD while making efficient concurrent algorithms and data structures.

The package is meant to be used as a library. If you have any advice/tip please let me know (ex: open an issue).


### Data
I will skip the data structures already implemented in the standard libraries (like linked lists).

**Stack** - basic stack (FILO) using the builtin linked list, can store any type, concurrency safe, no size limit, implements Stringer.

**Queue** - basic queue (FIFO) using the builtin linked list, can store any type, concurrency safe (optional mutex), no size limit, implements Stringer.