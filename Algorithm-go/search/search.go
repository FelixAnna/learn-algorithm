package search

import (
	"github.com/felixanna/algorithm-go/sort"
)

/* FindMin
find the smallest element in an array
*/
func FindMin(arr []int) (index, min int) {
	index = -1
	min = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			index = i
			min = arr[i]
		}
	}

	return
}

/* FindMax
find the largest element in an array
*/
func FindMax(arr []int) (index, max int) {
	index = -1
	max = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			index = i
			max = arr[i]
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
func FindNthMin(arr []int, n int) (int, int) {
	return findNthMinInternal(arr, n, 0, len(arr)-1)
}

func findNthMinInternal(arr []int, n int, start, end int) (int, int) {
	//skip if nth is not in range [0, len(arr)]
	if start > end {
		return -1, -1
	}

	partitionIndex := sort.RandQuickSortOneTime(arr, start, end)

	//fmt.Println(arr, pivotIndex, partitionIndex, n)
	switch {
	case partitionIndex == n-1:
		return partitionIndex, arr[partitionIndex]
	case partitionIndex > n-1:
		return findNthMinInternal(arr, n, start, partitionIndex-1)
	default:
		return findNthMinInternal(arr, n, partitionIndex+1, end)
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
func BinarySearch(arr []int, element int) int {
	if len(arr) == 0 {
		return -1
	}

	index := len(arr) / 2

	switch {
	case arr[index] == element:
		return index
	case arr[index] > element && index > 0:
		return BinarySearch(arr[:index], element)
	case arr[index] < element && index+1 < len(arr):
		return BinarySearch(arr[index+1:], element) + index + 1
	default:
		return -1
	}
}
