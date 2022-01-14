package sort

import (
	"fmt"
	"time"

	"github.com/felixanna/algorithm-go/lib"
)

/* sort
ways to sort:
	1. quick sort - O(nlogn) , sort in original array, less memory cost
	2. merge sort - O(nlogn) , need more memory
	3. counting sort - O(n+k) , for elements all in [m, n],  m and n are close numbers
	4. radix sort - O(n*k) , when all elements are positive, k normally not very large (like 32 / 4)
	5. bucket sort - O(n*k), split element in n bucket, then use counting sort to sort each bucket, finally append all of them
*/
func TestSort(n int, action int) {
	fmt.Println("TestSort:", n, action)
	arr := lib.Seed(n)

	start := time.Now()
	switch action {
	case 1:
		{
			QuickSort(arr)
		}
	case 2:
		{
			arr = MergeSort(arr)
		}
	case 3:
		{
			bottom, up := 10000, 35000
			arr = lib.SeedWithinRange(n, bottom, up)
			//fmt.Println(arr)
			arr = CountingSort(arr, bottom, up)
			//fmt.Println(arr)
		}
	case 4:
		{
			arr = RadixSort(arr)
		}
	case 5:
		{
			arr = BucketSort(arr)
		}
	default:
		fmt.Println("not support")
	}

	cost := time.Now().UnixMilli() - start.UnixMilli()
	fmt.Println("cost: ", cost)

	fmt.Println("sorted, checking")
	if err := lib.CheckSorted(arr, n); err != nil {
		fmt.Println(err)
	}

	fmt.Println("done")
}
