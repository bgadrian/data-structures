package graph

import (
	"testing"
)

func TestALBasic(t *testing.T) {
	g := AdjacencyList{}

	a, b, c := &Node{}, &Node{}, &Node{}
	g.AddNode(a, b, c)

	err := g.AddEdge(a, b)
	if err != nil {
		t.Error(err)
	}

	if ok, _ := g.Adjacent(a, b); ok == false {
		t.Error("a and b should be adjacent")
	}
}

func TestALLen(t *testing.T) {
	g := AdjacencyList{}
	a, b, c := Node{}, Node{}, Node{}

	quickAssertInt(0, g.LenNodes(), "map is not empty after init", t)

	g.AddNode(&a)
	quickAssertInt(1, g.LenNodes(), "after 1 insert", t)

	g.AddNode(&b)
	quickAssertInt(2, g.LenNodes(), "after 2 insert", t)

	g.AddNode(&c)
	quickAssertInt(3, g.LenNodes(), "after 3 insert", t)

	g.AddNode(&c)
	quickAssertInt(3, g.LenNodes(), "after same insert", t)

	g.RemoveNode(&a, &c)
	quickAssertInt(1, g.LenNodes(), "after 2 remove", t)
}

func TestALRemove(t *testing.T) {
	g := AdjacencyList{}
	a, b, c := &Node{}, &Node{}, &Node{}

	g.AddNode(a, b, c)
	g.AddEdge(a, b)
	g.AddEdge(a, c)
	g.RemoveNode(c)
	quickAssertInt(2, g.LenNodes(), "len after remove", t)
	quickAssertInt(1, len(g.a[a]), "1 neighbour left", t)

	g.RemoveNode(a, b)
	quickAssertInt(0, g.LenNodes(), "len after all ", t)

	//removing a non existing node should be ok
	g.RemoveNode(a)
}

func TestALNotExistent(t *testing.T) {
	g := AdjacencyList{}
	a, b := &Node{}, &Node{}

	_, err := g.Adjacent(a, b)
	if err == nil {
		t.Error("Adjacent did not returned error when A not existent")
	}
	err = g.AddEdge(a, b)
	if err == nil {
		t.Error("AddVertex did not returned error when A not existent")
	}

	g.AddNode(a)
	_, err = g.Adjacent(a, b)
	if err == nil {
		t.Error("Adjacent did not returned error when B not existent")
	}
	err = g.AddEdge(a, b)
	if err == nil {
		t.Error("AddVertex did not returned error when B not existent")
	}
}

func quickAssertInt(exp, got int, m string, t *testing.T) {
	if exp == got {
		return
	}

	t.Errorf("expected %v, got %v : %v", exp, got, m)
}

func quickAssertBool(exp, got bool, m string, t *testing.T) {
	if exp == got {
		return
	}

	t.Errorf("expected %v, got %v : %v", exp, got, m)
}
