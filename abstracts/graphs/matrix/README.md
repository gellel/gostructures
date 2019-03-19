##### Graph Adjacent Matrix

Adjacency matrix is a two-dimensonal collection of lists used to represent a finite graph. Each list describes the set of neighbours of a vertex (node) in the graph. In the adjacency matrix implementation, each vertex is held as a address within the top level container. When a key is accesssed (should it exist), it returns a list of its connections to other verticies in the graph. Generally, the unordered list will contain a collection of integers (or booleans), with each value being a either true (1) or false (0) to denote a connection. For a stable adjacency matrix, the number of rows and columns must be equal. Although it is described as unordered, the values must correspond with the outer-level container to accurately pair rows and columns.

###### Average time performance.

| Access    | Search    | Insert   | Delete |
| :---      | :---      | :---     | :---   |
| Θ(1)      | Θ(V)      | Θ(1)     | Θ(V+E) |

###### Worst-case time performance.

| Access   | Search    | Insert    | Delete |
| :---     | :---      | :---      | :---   |
| O(1)     | O(V)      | O(1)      | O(V+E) |
