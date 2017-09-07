/*Package graph ...*/
package graph

type empty struct{}

//Edge a non-directed edge/vertice
type Edge struct {
	a Node
	b Node
	v float64 //value, if any
	// d interface{} //object stored
}

//Arrow a directed edge/vertice
type Arrow struct {
	Edge
}

//Node represent a vertex
type Node struct {
	v float64 //value, if any
	// d interface{} //object stored
}

//Graph ...
type Graph interface {
	Adjacent(x, y Node) bool           // tests whether there is an edge from the vertex x to the vertex y;
	Neighbours(x Node) []Node          //: lists all vertices y such that there is an edge from the vertex x to the vertex y;
	AddNode(x ...Node)                 //: adds the vertex x, if it is not there;
	RemoveNode(x ...Node)              //: removes the vertex x, if it is there;
	AddEdge(x, y Node)                 //: adds the edge from the vertex x to the vertex y, if it is not there;
	RemoveEdge(x, y Node)              //: removes the edge from the vertex x to the vertex y, if it is there;
	GetVertexValue(x Node) float64     //: returns the value associated with the vertex x;
	SetVertexValue(x, v Node)          //: sets the value associated with the vertex x to v.
	GetEdgeValue(x, y Edge)            //: returns the value associated with the edge (x, y);
	SetEdgeValue(x, y Edge, v float64) //: sets the value associated with the edge (x, y) to v.
}
