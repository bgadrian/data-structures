package graph

import "testing"

//most of the functionalities are tested in adjacency_list_test.go

func TestALDAddDEdge(t *testing.T) {
	g := AdjacencyListDirected{}

	a, b, c := &Node{}, &Node{}, &Node{}

	g.AddNode(a, b, c)
	err := g.AddDirectedEdge(a, b)
	var ok bool

	if err != nil {
		t.Error(err)
	}

	ok, err = g.Adjacent(a, b)
	if err != nil {
		t.Error(err)
	}
	if ok == false {
		t.Error("adjacent didnt worked a->b")
	}

	ok, err = g.Adjacent(b, c)
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Error("adjacent worked in reverse a->b")
	}

	err = g.AddDirectedEdge(&Node{}, a)
	if err == nil {
		t.Error("AddDirectedEdge should return error on non existing A node")
	}

	err = g.AddDirectedEdge(a, &Node{})
	if err == nil {
		t.Error("AddDirectedEdge should return error on non existing B node")
	}
}

func TestALDRemoveDEdge(t *testing.T) {
	g := AdjacencyListDirected{}

	a, b, c := &Node{}, &Node{}, &Node{}

	g.AddNode(a, b, c)
	err := g.AddDirectedEdge(a, b)
	err = g.AddDirectedEdge(a, c)
	err = g.RemoveDirectedEdge(a, b)

	if err != nil {
		t.Error(err)
	}

	var ok bool
	ok, err = g.Adjacent(a, b)
	if ok {
		t.Error("RemoveDirectedEdge didn't worked")
	}

	ok, err = g.Adjacent(a, c)
	if ok == false {
		t.Error("RemoveDirectedEdge removed other edge")
	}

	err = g.RemoveDirectedEdge(&Node{}, a)
	if err == nil {
		t.Error("RemoveDirectedEdge should return error on non existing A node")
	}

	err = g.RemoveDirectedEdge(a, &Node{})
	if err == nil {
		t.Error("RemoveDirectedEdge should return error on non existing B node")
	}
}

func TestALDValueDEdge(t *testing.T) {
	g := AdjacencyListDirected{}

	a, b, c := &Node{}, &Node{}, &Node{}

	g.AddNode(a, b, c)
	var v interface{}
	var err error
	g.AddDirectedEdge(a, b)
	g.AddDirectedEdge(a, c)

	v, err = g.GetEdgeValue(a, b)

	if err != nil {
		t.Error(err)
	}

	if v != nil {
		t.Error("edge value default is not nil")
	}

	err = g.SetDirectedEdgeValue(a, b, 42)

	if err != nil {
		t.Error(err)
	}
	v, err = g.GetEdgeValue(a, b)

	if err != nil {
		t.Error(err)
	}

	if v.(int) != 42 {
		t.Error("SetDirectedEdgeValue didn't worked")
	}

	err = g.SetDirectedEdgeValue(&Node{}, a, 42)
	if err == nil {
		t.Error("SetDirectedEdgeValue should return error on non existing A node")
	}

	err = g.SetDirectedEdgeValue(a, &Node{}, 42)
	if err == nil {
		t.Error("SetDirectedEdgeValue should return error on non existing B node")
	}
}

func TestALDInterface(t *testing.T) {

	x := func(g Undirected) {

	}
	x(&AdjacencyListDirected{})

	y := func(g Directed) {

	}
	y(&AdjacencyListDirected{})
}
