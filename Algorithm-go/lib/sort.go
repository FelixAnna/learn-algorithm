package lib

import (
	"fmt"
	"math/rand"
	"time"
)

/* sort
ways to sort:
	1. quick sort - O(nlogn) , sort in original array, less memory cost
	2. merge sort - O(nlogn) , need more memory
*/
func TestSort(n int, action int) {
	fmt.Println("TestSort:", n, action)
	arr := seed(n)

	start := time.Now()
	switch action {
	case 1:
		{
			quicksort(arr, 0, n-1)

		}
	case 2:
		{
			arr = mergeSort(arr)
		}
	default:
		fmt.Println("not support")
	}

	cost := time.Now().UnixMilli() - start.UnixMilli()
	fmt.Println("cost: ", cost)

	fmt.Println("sorted, checking")
	if len(arr) != n {
		fmt.Println("Length incorrect:", len(arr), n)
		return
	}

	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			fmt.Println("Failed:", arr[i-1], arr[i])
			return
		}
	}

	fmt.Println("done")
}

/* quick sort：
if start>=end, no need to sort(at most 1 element)
	1. random find a pivot in the slice;
	2. swap pivot element with start;  ##[pivot, e2, e3, ... ,e1, ... , ei]
	3. sort and partition by start element: set partitionIndex = start，then from start + 1 to end (included), compare to start :  ## [pivot (partition index), e2, e3, ... , ei]
		a. partitionIndex is expected to the final location of pivot in a sorted array;
		b. if x <= pivot,
			* swap partitionIndex + 1 (now is bigger than start) with x;
			* set partitionIndex +=1 (now is the last element smaller or equal than start);
		c. if bigger, do nothing	 												e3>pivot  => ## [pivot, e2  (old / new partition index), e3, ... , ei]

	4. swap start（pivot) with partitionIndex, (we find pivot's real location, now left sides all smaller than pivot, right side all bigger than pivot)
	5. sort( arr, start, partition index - 1), sort(arr, partition index + 1, end)
*/
func quicksort(arr []int, start, end int) {

	//only have one element or no element
	if start >= end {
		return
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

	quicksort(arr, start, partitionIndex-1)
	quicksort(arr, partitionIndex+1, end)
}

func swap(arr []int, start, target int) {
	arr[start], arr[target] = arr[target], arr[start]
}

/* merge sort:
1. sort：
	a. if only 1 element, no need to sort, just retun;
	b. if more than one, split into 2 arrays, sort left， sort right, then merge left and right
2. merge :
	loop:
	a. pick element from first and second array,
	b. insert the smaller one to slice
	c. smaller one related array index +1 (ready to pick next in the smaller one array)
	d. if any array is empty, append another array to result slice
*/
func mergeSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen == 1 {
		return arr
	}

	middle := arrLen / 2

	/* user go routine and chan is quite slow here
	flag := make(chan bool, 2)
	left := make([]int, 0)
	right := make([]int, 0)
	go func() {
		left = mergeSort(arr[:middle])
		flag <- true
	}()

	go func() {
		right = mergeSort(arr[middle:])
		flag <- true
	}()

	<-flag
	<-flag
	*/
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	return merge(left, right)
}

func merge(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))

	i, j, t := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		fir, sec := arr1[i], arr2[j]
		if fir <= sec {
			result[t] = fir
			i++
		} else {
			result[t] = sec
			j++
		}

		t++
	}

	if i == len(arr1) { //first arr empty, append all second array
		for k := j; k < len(arr2); k++ {
			result[t] = arr2[k]
			t++
		}
	}

	if j == len(arr2) { //second arr empty, append all first array
		for k := i; k < len(arr1); k++ {
			result[t] = arr1[k]
			t++
		}
	}

	return result
}

func seed(n int) (arr []int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < n; i++ {
		arr = append(arr, r1.Intn(i+n))
	}

	return
}
