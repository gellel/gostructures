Hoare-partition Quick-sort.

Original partition scheme. Hoare partition creates two indices (i(left), j(right)) that start at the ends of the Array or Slice being partitioned. Per each iteration these indexes are shifted inward until index i is greater than or equal to j. 

| Best        | Average      | Worst     |
| :---        | :---         | :---      |
| Ω(n log(n)) | Θ(n log(n))  | O(n²)     |