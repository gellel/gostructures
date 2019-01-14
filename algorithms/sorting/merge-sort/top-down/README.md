Top-down Merge-sort.

Top-down merge sort generates a duplicate of the original Slice or Array before recursively splitting the minimum and maximum range of available iterations by N÷2. The algorithm then continiously divides current subset into futher N÷2 subsets until subset has a final length of 1. After there are no further subdivisions to perform, each subset is ordered, merged and merged again to produce a sorted sequence.

| Best        | Average      | Worst       |
| :---        | :---         | :---        |
| Ω(n log(n)) | Θ(n log(n))  | O(n log(n)) |