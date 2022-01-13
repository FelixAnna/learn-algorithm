# Algorithms implemented by Go

## Categories
### sort
    1. QuickSort (random pivot)
    2. MergeSort
    3. CountingSort
    4. RadixSort
    5. BucketSort

### Search
    1. Binary Search in sorted array
    2. Search Nth min element in unsorted array with QuickSort
    3. Find max / min element in unsorted array

### Dynamic Programming
    1. Solve cutrod issue
    2. Climbing stair issue -- TODO
    3. Fibonacci issue (fib.FibBottonUp in fib/fib.go)

### Divide And Conquer
    1. Merge Sort  (sort.MergeSort in sort/mergesort.go)
    2. others --TODO
## Test

### unit test
    run below in root folder:

```bash
go test ./...
```

### Manually Test

    update algorithm.go, and use:
```bash
go run
```