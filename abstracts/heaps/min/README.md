##### Minimum Heap

A Heap is an abstract data structure that is used to maintain an ordered collection of elements. Unlike an Array or other sequential Container, Heaps are a form of tree-based structure that enforces that the children of the relative parent (within the Heap) satisfy the required Heap condition. In the case of a Minimum Heap, the condition requires that if _P_ is a parent node of _C_, then the key (the value) of _P_ is either less than or equal the value of _C_. Although a Heap may implement an Array as the underlying structure to build the Heap, the intent is for the first element of the collection to be accessed per transaction and for each removal, the same first element is removed. Once an element is removed from the Heap, the Heap will then check whether the current element at H[0] satisfies the Heap property and if not, shift the current first child down the Heap order until it finds its Heap required position.

###### Average time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| Θ(1)      | Θ(n)      | Θ(log(n)) | Θ(log(n)) |

###### Worst-case time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| O(1)      | O(n)      | O(log(n)) | O(log(n)) |
