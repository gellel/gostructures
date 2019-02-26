##### Priority Queue (Max Priority)

The Priority Queue is an abstract data structure that stores collections of data in an ordered arrangment based on priority heirarchy. Similar to the _Heap_ (which most Priority Queues will implement), Priority Queues are tree-based data structures that enforces that the children of the relative parent satisfy the required priority condition. In the case of a Maximum Priority Queue, the condition requires that if _P_ is a parent node of _C_, then the priority (the value) of node _P_ is either greater than or equal the value of node _C_. Again, like the _Heap_, a Priority Queue offers transactions on the first element in the collection. Once an element is removed from the Priority Queue the Queue will then check whether the current element at Q[0] satisfies the priority property and if not, shift the current first child down the Queue order until it finds its required position. Unlike a basic _Heap_ a Priority Queue can hold any form of data, but requires each item to be submitted with a defined numeric priority.

###### Average time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| Θ(1)      | Θ(n)      | Θ(log(n)) | Θ(log(n)) |

###### Worst-case time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| O(1)      | O(n)      | O(log(n)) | O(log(n)) |

###### Average time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| Θ(log(n)) | Θ(log(n)) | Θ(log(n)) | Θ(log(n)) |

###### Worst-case time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| O(log(n)) | O(log(n)) | O(log(n)) | O(log(n)) |
