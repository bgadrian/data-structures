package tree

import (
	"errors"
)

//Node An element of the tree.
type Node struct {
	K     int
	V     interface{}
	Left  *Node
	Right *Node
}

//BST Binary Search Tree simple implementation.
//it does NOT have auto balance algorithm.
type BST struct {
	Root *Node
	n    int
	h    int
}

//InsertKeys Add multiple elements. Doesn't support adding values, only keys.
func (b *BST) InsertKeys(keys ...int) (int, error) {
	i := 0
	for _, k := range keys {
		_, err := b.Insert(k, nil)
		if err != nil {
			return i, err
		}
		i++
	}
	return i, nil
}

//Insert Add a new element.
func (b *BST) Insert(key int, value interface{}) (n *Node, err error) {

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
	b.n++

	return
}

//Search Find a node with a specific key.
func (b *BST) Search(key int) (n *Node) {
	node, _ := b.SearchParent(key)
	return node
}

//SearchParent Return the node with a specific key, and it's parent.
func (b *BST) SearchParent(key int) (n *Node, p *Node) {
	for c := b.Root; c != nil && n == nil; {
		if c.K == key {
			n = c
		} else if key < c.K {
			p = c
			c = c.Left
		} else {
			p = c
			c = c.Right
		}
	}
	return
}

//Remove Delete an element with a specific key.
func (b *BST) Remove(key int) error {
	if b.Root == nil {
		return errors.New("tree is empty")
	}

	var r error

	if b.Root.K == key {
		auxRoot := &Node{}
		auxRoot.Left = b.Root
		r = b.Root.remove(key, auxRoot)
		b.Root = auxRoot.Left
	} else {
		r = b.Root.remove(key, nil)
	}

	if r == nil {
		b.n--
	}
	return r

}

func (n *Node) remove(key int, parent *Node) error {
	if key < n.K {
		if n.Left == nil {
			return errors.New("cannot find key")
		}
		return n.Left.remove(key, n)
	}

	if key > n.K {
		if n.Right == nil {
			return errors.New("cannot find key")
		}
		return n.Right.remove(key, n)
	}

	if n.Left != nil && n.Right != nil {
		min := n.Right.minNode()
		n.K = min.K
		n.V = min.V
		return n.Right.remove(n.K, n)
	}

	if parent.Left == n {
		if n.Left == nil {
			parent.Left = n.Right
		} else {
			parent.Left = n.Left
		}
		return nil
	}

	if parent.Right == n {
		if n.Left == nil {
			parent.Right = n.Right
		} else {
			parent.Right = n.Left
		}
	}
	return nil
}

func (n *Node) minNode() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.minNode()
}

//Size How many elements are now.
func (b *BST) Size() int {
	return b.n
}

//Inorder Get the Inorder list of keys.
func (b *BST) Inorder() []int {
	var traverse func(n *Node, r []int, i *int)
	traverse = func(n *Node, r []int, i *int) {
		if n == nil {
			return
		}
		traverse(n.Left, r, i)
		r[*i] = n.K
		(*i)++
		traverse(n.Right, r, i)
	}
	arr := make([]int, b.n)
	i := 0
	traverse(b.Root, arr, &i)
	return arr
}

//Height ...
// func (b *BinarySearch) Height() int {
// 	if b.Root == nil {
// 		return -1
// 	}
// 	return b.h
// }
