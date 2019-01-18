Single Linked-List

Sequenced linear collection of data. Each linked list node points to an optional adjacent linked list node. Appending to the linked list places the newest item to end of the linked list (known as the Tail). The appended data then becomes the new Tail in the linked list. Prepending puts the newest item to the beginning of the linked list (known as the Head). The current Head then becomes the new nodes adjacent node before the new node is set as the current Head.

Best time performance.

| Access              | Search  | Insert  | Delete  |
| :---                | :---    | :---    | :---    |
| Ω(1) (Head or Tail) | Ω(n)    | Ω(1)    | Ω(n)    |

Average time performance.

| Access    | Search    | Insert    | Delete    |
| :---      | :---      | :---      | :---      |
| Θ(n)      | Θ(n)      | Θ(1)      | Θ(n)      |

Worst-case time performance.

| Access   | Search    | Insert    | Delete    |
| :---     | :---      | :---      | :---      |
| O(n)     | O(n)      | O(1)      | O(n)      |