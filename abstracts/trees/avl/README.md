##### AVL Search Tree

AVL Search Tree (**A**delson-**V**elsky and **L**andis) is an abstract data structure that stores collections of data in an ordered arrangment based on value. In a numeric AVL Tree, it assigns new AVL-Tree-Node's based on the value in each accessed node. To insert a new node to the left of the an accessed node, the argument value must be less than the value stored in the node or the left sub-tree. For a value to be inserted to the right, the value must be greater than the value stored in the accessed node or right sub-tree. Unlike the _Binary Search Tree_, the _AVL Tree_ offers guaranteed O(log(_n_)) access, searching, insertion and deletion as a result of balacing operations that are performed the AVL Tree is modified. Each rotation is designed to ensure that the leaves of the AVL Tree are evenly weighted, with the same number of child nodes across each parent node. The AVL Tree performs four different kinds of operations to balance its children - the Left Rotation, Right Rotation, Left-Right Rotation and Right-Left Rotation.

###### Average time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| Θ(log(n)) | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) |

###### Worst-case time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| O(log(n)) | O(log(n)) | O(log(n)) | O(log(n)) |
