package graph

import (
	"fmt"
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

	if ok, _ := g.Adjacent(a, c); ok == true {
		t.Error("a and b should not be adjacent")
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

func TestALRemoveEdge(t *testing.T) {
	g := AdjacencyList{}
	a, b, c := &Node{}, &Node{}, &Node{}

	g.AddNode(a, b, c)
	g.AddEdge(a, b)
	g.AddEdge(a, c)

	if ok, _ := g.Adjacent(a, b); ok == false {
		t.Error("edge adjacent didn't worked ")
	}

	g.RemoveEdge(a, b)
	if ok, _ := g.Adjacent(a, b); ok == true {
		t.Error("edge remove didn't worked")
	}

	if ok, _ := g.Adjacent(a, c); ok == false {
		t.Error("edge remove affected other edge")
	}

	quickAssertInt(3, g.LenNodes(), "remove edge removed nodes too", t)

	if err := g.RemoveEdge(&Node{}, b); err == nil {
		t.Error("edge remove non existent A was ok")
	}

	if err := g.RemoveEdge(a, &Node{}); err == nil {
		t.Error("edge remove non existent B was ok")
	}
}

func TestALNodeValues(t *testing.T) {
	g := AdjacencyList{}
	a, b, c := &Node{"a"}, &Node{}, &Node{}

	g.AddNode(a, b)
	v, err := g.GetNodeValue(a)
	if err != nil {
		t.Error(err)
	}

	if v != "a" {
		t.Error("value was removed when inserting")
	}

	_, err = g.GetNodeValue(c)
	if err == nil {
		t.Error("not error when ask for non existing node's value")
	}

	err = g.SetNodeValue(c, "C")
	if err == nil {
		t.Error("not error when setValue for non existing node")
	}

	err = g.SetNodeValue(b, "B")
	if err != nil {
		t.Error(err)
	}

	v, err = g.GetNodeValue(b)
	if err != nil {
		t.Error(err)
	}

	if v != "B" {
		t.Error("SetNodeValue failed")
	}
}

func TestALEdgeValue(t *testing.T) {
	g := AdjacencyList{}
	a, b, c := &Node{"a"}, &Node{}, &Node{}
	var v interface{}

	g.AddNode(a, b, c)
	g.AddEdge(a, b)
	g.AddEdge(a, c)

	err := g.SetEdgeValue(a, b, 0.5)
	if err != nil {
		t.Error(err)
	}

	v, err = g.GetEdgeValue(a, b)

	if v != 0.5 {
		t.Errorf("exp 0.5, got %v getEdgeValue", v)
	}

	if err != nil {
		t.Error(err)
	}

	v, err = g.GetEdgeValue(b, c)

	if err == nil {
		t.Error("GetEdgeValue didn't returned error on missing edge")
	}

	err = g.SetEdgeValue(&Node{}, a, 100)
	if err == nil {
		t.Error("SetEdgeValue didn't returned error when node A is missing")
	}

	err = g.SetEdgeValue(a, &Node{}, 100)
	if err == nil {
		t.Error("SetEdgeValue didn't returned error when node B is missing")
	}

	_, err = g.GetEdgeValue(&Node{}, a)
	if err == nil {
		t.Error("GetEdgeValue didn't returned error when node A is missing")
	}

	_, err = g.GetEdgeValue(a, &Node{})
	if err == nil {
		t.Error("GetEdgeValue didn't returned error when node B is missing")
	}

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

func TestALInterface(t *testing.T) {

	x := func(g Graph) {

	}
	x(&AdjacencyList{})
}

func TestALNeighbours(t *testing.T) {
	g := AdjacencyList{}
	a, b, c, d := &Node{}, &Node{}, &Node{}, &Node{}

	g.AddNode(a, b, c, d)
	g.AddEdge(a, b)
	g.AddEdge(a, c)

	n, err := g.Neighbours(a)

	if err != nil {
		t.Error(err)
	}

	quickAssertInt(2, len(n), "2 neighbours", t)
	if contains(n, b) == false {
		t.Error("neighbours does not contain b")
	}

	if contains(n, c) == false {
		t.Error("neighbours does not contain b")
	}

	_, err = g.Neighbours(&Node{})

	if err == nil {
		t.Error("didn't returned error on a non existing node")
	}
}

func contains(arr []*Node, a *Node) bool {
	for _, n := range arr {
		if n == a {
			return true
		}
	}
	return false
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

func ExampleAdjacencyList() {
	g := AdjacencyList{}
	//node values can be set at init
	a, b, c, d := &Node{"A"}, &Node{"B"}, &Node{}, &Node{}

	g.AddNode(a, b, c, d)
	//node values can be set after init too
	c.v = "C"
	g.SetNodeValue(d, "D")
	fmt.Println("#nodes=", g.LenNodes(), "a value=", a.v)

	g.AddEdge(a, b)
	g.AddEdge(a, c)
	nOfA, _ := g.Neighbours(a)
	fmt.Println("neighbours of A=", nOfA[0].v, nOfA[1].v)

	bAndC, _ := g.Adjacent(b, c)
	fmt.Println("B neighbour with C?", bAndC)

	//add and Edge and set a value of it
	g.SetEdgeValue(b, d, 42)
	BD, _ := g.GetEdgeValue(b, d)
	fmt.Println("edge value BD=", BD)

	// Output:
	//#nodes= 4 a value= A
	//neighbours of A= B C
	//B neighbour with C? false
	//edge value BD= 42
}
