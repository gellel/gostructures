##### Double Linked-List

Sequenced bi-linear collection of data. Each linked list node points to an optional previous and next linked list node. Appending to the linked list places the newest item to end of the linked list (known as the Tail). The appended data then becomes the new Tail in the linked list. Prepending puts the newest item to the beginning of the linked list (known as the Head). The current Head then becomes the new nodes adjacent node before the new node is set as the current Head. Unlike the singly linked list, the doubly linked list can move both forward and backwards between linked list nodes without implementing a reverse sort.

###### Best time performance.

| Access              | Search | Insert  | Delete  |
| :---                | :---   | :---    | :---    |
| Ω(1) (Head or Tail) | Ωn)    | Ω(1)    | Ω(n)    |

###### Average time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| Θ(n)      | Θ(n)      | Θ(1)      | Θ(n)      |

###### Worst-case time performance.

| Access   | Search    | Insert    | Delete    |
| :---     | :---      | :---      | :---      |
| O(n)     | O(n)      | O(1)      | O(n)      |