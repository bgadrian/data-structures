## graph [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/BTooLs/data-structures/graph)
A collection of simple graph data structures, used for **academic purposes**. More info on [wiki](https://en.wikipedia.org/wiki/Graph_(abstract_data_type)) or [visual example](https://www.tutorialspoint.com/data_structures_algorithms/graph_data_structure.htm)

Nodes and edges can have values, so **weighted** graphs can be built using the same data structures.


### AdjacencyList [description](https://en.wikipedia.org/wiki/Adjacency_list)
AdjacencyList is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a vertex in the graph. This is one of several commonly used representations of graphs for use in computer programs.

The graph is undirected with values on nodes and edges.

#### AdjacencyList implementation
Each node and edge can have a value ```interface{}```.

Internally the data is stored as a map of maps ```map[*Node]map[*Node]interface{}```. 

### AdjacencyListDirected 
It is a AdjacencyList with 3 extra functions, that allow 1 direction edge control.