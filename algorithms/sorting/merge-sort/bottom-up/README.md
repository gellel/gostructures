Bottom-up Merge-sort.

Bottom-up merge sort sequentially partitions adjacent elements in the provided Slice or Array before merging the ordered sequence. Each partition grows by a factor of 2 per iteration cylce for a continously growing range of adjacent elements until N number of partitions been performed. Unlike its sibling Top-down merge sort, Bottom-up does not require recursion to perform the core sorting operation.

| Best        | Average      | Worst       |
| :---        | :---         | :---        |
| Ω(n log(n)) | Θ(n log(n))  | O(n log(n)) |