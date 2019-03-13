##### Graph Adjacent List

Adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbours of a vertex (node) in the graph. In the adjacency list implementation, each vertex is held as a key within a graph. When a key is accesssed (should it exist), it returns a list of its connections to other verticies in the graph. Generally, the unordered list will contain a collection of strings, with each string being a reference to a key in the graph.

###### Average time performance.

| Access    | Search    | Insert   | Delete     |
| :---      | :---      | :---     | :---       |
| Θ(1)      | Θ(V)      | Θ(1)     | Θ(|V|+|E|) |

###### Worst-case time performance.

| Access   | Search    | Insert    | Delete     |
| :---     | :---      | :---      | :---       |
| O(1)     | O(V)      | O(1)      | O(|V|+|E|) |