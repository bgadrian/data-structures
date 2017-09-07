package graph

import "errors"

//AdjacencyList is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a vertex in the graph.
//This is one of several commonly used representations of graphs for use in computer programs.
type AdjacencyList struct {
	a map[*Node]map[*Node]empty
}

func (g *AdjacencyList) lazyInit() {
	if g.a != nil {
		return
	}

	g.a = make(map[*Node]map[*Node]empty)
}

//AddNode Add a vertex, if doesn't exists already.
func (g *AdjacencyList) AddNode(x ...*Node) {
	g.lazyInit()

	for _, n := range x {
		if g.a[n] == nil {
			g.a[n] = make(map[*Node]empty)
		}
	}
}

//RemoveNode ...
func (g *AdjacencyList) RemoveNode(x ...*Node) {

	g.lazyInit()

	for _, n := range x {
		for neighbour := range g.a[n] {
			delete(g.a[neighbour], n)
		}

		delete(g.a, n)
	}
}

//AddEdge ...
func (g *AdjacencyList) AddEdge(x, y *Node) error {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return errors.New("node X doesn't exists")
	}
	if _, ok := g.a[y]; ok == false {
		return errors.New("node Y doesn't exists")
	}

	g.a[x][y] = empty{}
	g.a[y][x] = empty{}
	return nil
}

//LenNodes ...
func (g *AdjacencyList) LenNodes() int {
	g.lazyInit()

	return len(g.a)
}

//Adjacent ...
func (g *AdjacencyList) Adjacent(x, y *Node) (bool, error) {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return false, errors.New("node X doesn't exists")
	}
	if _, ok := g.a[y]; ok == false {
		return false, errors.New("node Y doesn't exists")
	}

	_, ok := g.a[x][y]
	return ok, nil
}
