##### Red Black Search Tree

The Red Black Tree is an abstract data structure that stores collections of data in an ordered arrangment based on value and colour heirarchy. In a Red Black Tree, new Red-Black-Tree-Node's are assigned based on the value held by each accessed node and their adjacent colours. To insert a new node to the left of the an accessed node, the argument value must be less than the value stored in the node or the left sub-tree, and the node must be coloured red. For a value to be inserted to the right, the value must be greater than the value stored in the accessed node or right sub-tree, and also coloured red. Unlike the _AVL Search Tree_, the _Red Black Tree_ performs more constrained and checks and balances per insertion each time the Tree is modified. Each rotation is designed to ensure that the leaves of the Red Black Tree are evenly weighted and coloured, with the same number of child nodes across each parent node, and that every node visited touches the same number of opposite coloured nodes as a cursor traverses the tree. The Red Black Tree performs four different kinds of operations to balance its children - the Left Rotation, Right Rotation, Left-Right Rotation and Right-Left Rotation.

###### Average time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| Θ(log(n)) | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) |

###### Worst-case time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| O(log(n)) | O(log(n)) | O(log(n)) | O(log(n)) |
