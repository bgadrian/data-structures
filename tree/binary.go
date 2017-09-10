package tree

import (
	"errors"
)

//Node ...
type Node struct {
	K     int
	V     interface{}
	Left  *Node
	Right *Node
}

// func (n *Node) String() string {
// 	return string(n.K)
// }

//BinarySearch ...
type BinarySearch struct {
	Root *Node
	n    int
	h    int
}

//Insert ...
func (b *BinarySearch) Insert(key int, value interface{}) (n *Node, err error) {

	if b.Root == nil {
		n = &Node{K: key, V: value}
		b.Root = n
		b.n++
		return
	}

	p := b.Root
	l := &p
	height := 0
	for {
		if key == p.K {
			err = errors.New("The key already exists")
			return
		}

		if key < p.K {
			l = &p.Left
		} else {
			l = &p.Right
		}

		if *l == nil {
			break
		}
		p = *l
		height++
	}

	if height > b.h {
		b.h = height
	}

	*l = &Node{K: key, V: value}
	n = *l

	return
}

//Search ...
func (b *BinarySearch) Search(key int) (n *Node) {
	for c := b.Root; c != nil && n == nil; {
		if c.K == key {
			n = c
		} else if key < c.K {
			c = c.Left
		} else {
			c = c.Right
		}
	}
	return
}

//Height ...
// func (b *BinarySearch) Height() int {
// 	if b.Root == nil {
// 		return -1
// 	}
// 	return b.h
// }
