package sort

import (
	"math/rand"
)

/* Quicksort
if start>end, no need to sort, otherwise:
	1. use RandQuickSortOneTime to sort for one round, keep the return value partition Index
	2. sort( arr, start, partition index - 1), sort(arr, partition index + 1, end) use RandQuickSortOneTime
*/
func QuickSort(arr []int) {
	//only have one element or no element : no need to sort
	if len(arr) <= 1 {
		return
	}

	quickSortRecursive(arr, 0, len(arr)-1)
}

func quickSortRecursive(arr []int, start, end int) {

	//only have one element or no element : no need to sort
	if start > end {
		return
	}

	partitionIndex := RandQuickSortOneTime(arr, start, end)

	quickSortRecursive(arr, start, partitionIndex-1)
	quickSortRecursive(arr, partitionIndex+1, end)
}

/* RandQuickSortOneTime
use a random picked element as pivot, sort one round for this pivot:
    1. random find a pivot in the slice;
	2. swap pivot element with start;  ##[pivot, e2, e3, ... ,e1, ... , ei]
	3. sort and partition by start element: set partitionIndex = start，then from start + 1 to end (included), compare to start :  ## [pivot (partition index), e2, e3, ... , ei]
		a. partitionIndex is expected to the final location of pivot in a sorted array;
		b. if x <= pivot,
			* swap partitionIndex + 1 (now is bigger than start) with x;
			* set partitionIndex +=1 (now is the last element smaller or equal than start);
		c. if bigger, do nothing	 												e3>pivot  => ## [pivot, e2  (old / new partition index), e3, ... , ei]

	4. swap start（pivot) with partitionIndex, (we find pivot's real location, now left sides all smaller than pivot, right side all bigger than pivot)
return the random pivot partitionIndex in array
*/
func RandQuickSortOneTime(arr []int, start, end int) int {
	if start == end { //no need to search
		return start
	}

	pivotIndex := rand.Intn(end-start) + start

	swap(arr, start, pivotIndex)
	partitionIndex := start //record pivot location in array

	for i := start + 1; i <= end; i++ {
		if arr[i] <= arr[start] {
			partitionIndex += 1
			swap(arr, i, partitionIndex)
		}
	}

	swap(arr, start, partitionIndex)

	return partitionIndex
}

func swap(arr []int, start, target int) {
	arr[start], arr[target] = arr[target], arr[start]
}
