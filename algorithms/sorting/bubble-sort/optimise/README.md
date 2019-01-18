Optimised Bubble Sort.

Bubble sort iterates sequentially across the length of the Array or Slice. Starting from the beginning of the sequence, the algorithm compares every adjacent pair in the sequence, swapping the compared elements if are not in the correct order. After each iteration, one less element (n - i) is needed to be compared. The operation continues until there are no more elements in the sequence to be compared. Unlike it's unoptimised counterpart, the optimised Bubble Sort can avoid looking at the last n − 1 items when running for the n-th time within the inner (j) loop.

| Best        | Average      | Worst       |
| :---        | :---         | :---        |
| Ω(n)        | Θ(n²)        | O(n²)       |