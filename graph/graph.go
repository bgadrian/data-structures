/*Package graph contains a series of simple graph data structures, used for academic purposes.*/
package graph

//Node represent a vertex
type Node struct {
	v interface{} //object stored
}

//Undirected Common interface for all graph implementations
type Undirected interface {
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

//Directed Common interface for all directed graph implementations
type Directed interface {
	Undirected
	AddDirectedEdge(x, y *Node) error
	RemoveDirectedEdge(x, y *Node) error
	SetDirectedEdgeValue(x, y *Node, v interface{}) error
}
