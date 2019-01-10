Lomuto-partition Quick-sort.

The modified partition scheme. The algorithm maintains index i as it scans the seqeuence using another index j to compare that the elements Array[j] are less than or equal to the pivot. Less efficient than Hoare's original scheme. This scheme degrades to O(n2) when the sequence is already in order.

| Best        | Average      | Worst     |
| :---        | :---         | :---      |
| Ω(n log(n)) | Θ(n log(n))  | O(n²)     |