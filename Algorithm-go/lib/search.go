package lib

import (
	"fmt"
	"time"
)

/* sort
ways to search:
	1. bin search in a sorted array
	2. find xth max / min element in an unsorted array
*/
func TestSearch(n int, action int) {
	fmt.Println("TestSearch:", n, action)
	arr := seed(n)
	var element int
	var index int = -1

	start := time.Now()
	switch action {
	case 1:
		{
			element = arr[n/2]
			arr = radixSort(arr)
			index = binSearch(arr, element)

		}
	case 2:
		{
			index, element = minNthSearch(arr, 5, 0, len(arr)-1)
		}
	case 3:
		{
			index, element = minSearch(arr)
		}
	default:
		fmt.Println("not support")
	}

	cost := time.Now().UnixMilli() - start.UnixMilli()
	fmt.Println("cost: ", cost)

	fmt.Println("result: ", arr, index, element)

	fmt.Println("search, checking")
	if index >= 0 && arr[index] != element {
		fmt.Println("Not same incorrect:", index, arr[index], element)
		return
	}

	fmt.Println("done")
}

func minSearch(arr []int) (index, min int) {
	min = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			index = i
			min = arr[i]
		}
	}

	return
}

/* min Nth search
search the min nth [1, len(arr)] element in array (unsorted):
1. use quick search methods, random pick an element
2. swap with first element
3. loop from the second to end, to find all elements smaller than the first one
4. now we know the position of the first element in a sorted array
	a. if the posistion is desired rank, just return
	b. if greater, search nTh in the left slice branch
	c. if smaller, search n - position th elements in the right slice
*/
func minNthSearch(arr []int, n int, start, end int) (int, int) {
	//skip if nth is not in range [0, len(arr)]
	if start > end {
		return -1, -1
	}

	partitionIndex := randQuickSort(arr, start, end)

	//fmt.Println(arr, pivotIndex, partitionIndex, n)
	switch {
	case partitionIndex == n-1:
		return partitionIndex, arr[partitionIndex]
	case partitionIndex > n-1:
		return minNthSearch(arr, n, start, partitionIndex-1)
	default:
		return minNthSearch(arr, n, partitionIndex+1, end)
	}
}

/* binary search:
can search in a sorted array;
1. find the middle element in a array;
2. compare the middle one with target:
	a. if equal, return the index;
	b. if greater than target, then search in the left sub-slice
	c. if smaller than target, then search in the right sub-slice
	d. if array is empty before find one, return -1 (not found)
*/
func binSearch(arr []int, element int) int {
	if len(arr) == 0 {
		return -1
	}

	index := len(arr) / 2

	switch {
	case arr[index] == element:
		return index
	case arr[index] > element && index > 0:
		return binSearch(arr[:index], element)
	case arr[index] < element && index+1 < len(arr):
		return binSearch(arr[index+1:], element) + index + 1
	default:
		return -1
	}
}
