package graph

import "errors"

//AdjacencyList is a collection of unordered lists used to represent a finite graph.
//Each list describes the set of neighbors of a vertex in the graph.
//This is one of several commonly used representations of graphs for use in computer programs.
type AdjacencyList struct {
	a map[*Node]map[*Node]interface{}
}

func (g *AdjacencyList) lazyInit() {
	if g.a != nil {
		return
	}

	g.a = make(map[*Node]map[*Node]interface{})
}

//AddNode Add a vertex, if doesn't exists already.
func (g *AdjacencyList) AddNode(x ...*Node) {
	g.lazyInit()

	for _, n := range x {
		if g.a[n] == nil {
			g.a[n] = make(map[*Node]interface{})
		}
	}
}

//RemoveNode Remove a vertex from the graph.
func (g *AdjacencyList) RemoveNode(x ...*Node) {

	g.lazyInit()

	for _, n := range x {
		for neighbour := range g.a[n] {
			delete(g.a[neighbour], n)
		}

		delete(g.a, n)
	}
}

//AddEdge Add a vertice between 2 nodes.
//if is a directed graph adds 2 edges (x->y and x<-y)
func (g *AdjacencyList) AddEdge(x, y *Node) error {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return errors.New("node X doesn't exists")
	}
	if _, ok := g.a[y]; ok == false {
		return errors.New("node Y doesn't exists")
	}

	g.a[x][y] = nil
	g.a[y][x] = nil
	return nil
}

//RemoveEdge Remove a vertice between two nodes
//if is a directed graph removes 2 edges (x->y and x<-y)
func (g *AdjacencyList) RemoveEdge(x, y *Node) error {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return errors.New("node X doesn't exists")
	}

	if _, ok := g.a[y]; ok == false {
		return errors.New("node Y doesn't exists")
	}

	delete(g.a[x], y)
	delete(g.a[y], x)
	return nil
}

//LenNodes How many distinct nodes are in the graph.
func (g *AdjacencyList) LenNodes() int {
	g.lazyInit()

	return len(g.a)
}

//Adjacent Check if two nodes are connected with an edge.
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

//GetNodeValue Returns the nodes value (alias x.v)
func (g *AdjacencyList) GetNodeValue(x *Node) (interface{}, error) {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return nil, errors.New("node X doesn't exists")
	}

	return x.v, nil
}

//SetNodeValue Set value of a node (alias a := Node{0.4})
func (g *AdjacencyList) SetNodeValue(x *Node, v interface{}) error {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return errors.New("node X doesn't exists")
	}
	x.v = v
	return nil
}

//GetEdgeValue Gets an edge value, if exists.
func (g *AdjacencyList) GetEdgeValue(x, y *Node) (interface{}, error) {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return nil, errors.New("node X doesn't exists")
	}
	if _, ok := g.a[y]; ok == false {
		return nil, errors.New("node Y doesn't exists")
	}

	v, ok := g.a[x][y]
	if ok == false {
		return nil, errors.New("edge does not exists")
	}

	return v, nil
}

//SetEdgeValue Create or updates a vertice's value
func (g *AdjacencyList) SetEdgeValue(x, y *Node, v interface{}) error {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return errors.New("node X doesn't exists")
	}
	if _, ok := g.a[y]; ok == false {
		return errors.New("node Y doesn't exists")
	}

	g.a[x][y] = v
	g.a[y][x] = v
	return nil
}

//Neighbours ...
func (g *AdjacencyList) Neighbours(x *Node) ([]*Node, error) {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return nil, errors.New("node X doesn't exists")
	}

	r := make([]*Node, len(g.a[x]))
	i := 0
	for n := range g.a[x] {
		r[i] = n
		i++
	}
	return r, nil
}
