Binary Search Tree

Binary Search Tree is an abstract data structure that stores collections of data in an ordered arrangment based on value. In the case of a numeric Binary-Search-Tree, it assigns new Binary-Tree-Node's based on the value in each node. To insert a new node to the left of the root node, the argument value must be less than the value stored in the root or the left sub-tree. For a value to be inserted to the right of the root node the value must be greater than the value stored in the root node or right sub-tree.

Average time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| Θ(log(n)) | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) |

Worst-case time performance.

| Access   | Search    | Insert    | Delete    |
| :---     | :---      | :---      | :---      |
| Ω(n)     | Ω(n)      | Ω(n)      | Ω(n)      |