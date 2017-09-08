package graph

import "errors"

//AdjacencyListDirected Directed graphs
type AdjacencyListDirected struct {
	AdjacencyList
}

//AddDirectedEdge Add the x->y vertice with default value.
func (g *AdjacencyListDirected) AddDirectedEdge(x, y *Node) error {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return errors.New("node X doesn't exists")
	}
	if _, ok := g.a[y]; ok == false {
		return errors.New("node Y doesn't exists")
	}

	g.a[x][y] = nil
	return nil
}


//RemoveDirectedEdge Remove the x->y vertice, if exists.
func (g *AdjacencyList) RemoveDirectedEdge(x, y *Node) error {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return errors.New("node X doesn't exists")
	}

	if _, ok := g.a[y]; ok == false {
		return errors.New("node Y doesn't exists")
	}

	delete(g.a[x], y)
	return nil
}


//SetDirectedEdgeValue ... 
func (g *AdjacencyList) SetDirectedEdgeValue(x, y *Node, v interface{}) error {
	g.lazyInit()

	if _, ok := g.a[x]; ok == false {
		return errors.New("node X doesn't exists")
	}
	if _, ok := g.a[y]; ok == false {
		return errors.New("node Y doesn't exists")
	}

	g.a[x][y] = v
	return nil
}
