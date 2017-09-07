/*Package graph ...*/
package graph

//Node represent a vertex
type Node struct {
	v interface{} //object stored
}

//Graph ...
type Graph interface {
	Adjacent(x, y *Node) (bool, error)            // tests whether there is an edge from the vertex x to the vertex y;
	Neighbours(x *Node) ([]*Node, error)          //: lists all vertices y such that there is an edge from the vertex x to the vertex y;
	AddNode(x ...*Node)                           //: adds the vertex x, if it is not there;
	RemoveNode(x ...*Node)                        //: removes the vertex x, if it is there;
	AddEdge(x, y *Node) error                     //: adds the edge from the vertex x to the vertex y, if it is not there;
	RemoveEdge(x, y *Node) error                  //: removes the edge from the vertex x to the vertex y, if it is there;
	GetNodeValue(x *Node) (interface{}, error)    //: returns the value associated with the vertex x;
	SetNodeValue(x *Node, v interface{}) error    //: sets the value associated with the vertex x to v.
	GetEdgeValue(x, y *Node) (interface{}, error) //: returns the value associated with the edge (x, y);
	SetEdgeValue(x, y *Node, v interface{}) error //: sets the value associated with the edge (x, y) to v.
}
