package lib

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
	case 3:
		{
			bottom, up := 10000, 35000
			arr = seedWithinRange(n, bottom, up)
			//fmt.Println(arr)
			arr = countingSort(arr, bottom, up)
			//fmt.Println(arr)
		}
	case 4:
		{
			arr = radixSort(arr)
		}
	case 5:
		{
			arr = bucketSort(arr)
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

/* counting sort:
suppose all elements in an array are in a relative small range 0~k， then we can use counting sort to sort the array in O(n+k)
steps:
	1. define counting array[k+1] : ca[k+1]     				example: [0, 0, 0]
	2. scan all elements in input array, set: ca[input[i]]++    example: [1, 3, 2]
	3. scan ca from 0 - k: make ca[i]=ca[i-1] + ca[i] 			example: [1, 4, 6]
	4. create a new array as the input array len: result array;
	5. scan input array (from last to first), result[ca[input[i]] - 1] = input[i]， set: ca[input[i]] = ca[input[i]] - 1
*/
func countingSort(arr []int, bottom, up int) []int {
	ca := make([]int, up-bottom+1)
	for i := 0; i < len(arr); i++ {
		ca[arr[i]-bottom]++
	}

	for i := 1; i < len(ca); i++ {
		ca[i] += ca[i-1]
	}

	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result[ca[arr[i]-bottom]-1] = arr[i]
		ca[arr[i]-bottom] -= 1
	}

	return result
}

/* radix sort:
when element are positive numbers and distributed a little evenly in a large range, we can use radix sort:
step:
	1. define the max bit in numbers, and the base bit we use for split
	2. get the base digit we used for counting: baseDigit = 1 << baseBit
	3. loop for largest maxBit/baseBit times:
		a. sort element on k th digit with counting sort
		b. copy sorted array to original array as new input for next round loop
*/
func radixSort(arr []int) []int {
	result := make([]int, len(arr))
	baseBit := 4
	maxBit := 32
	baseDigit := 1 << 8
	for r := 0; r < maxBit/baseBit; r++ {
		shift := math.Pow(float64(baseDigit), float64(r))
		count := make([]int, baseDigit)

		for i := 0; i < len(arr); i++ {
			countIndex := getCount(arr, i, int(shift), baseDigit) //get rank of special position
			count[countIndex]++
		}

		for i := 1; i < len(count); i++ {
			count[i] += count[i-1]
		}
		//fmt.Println(count)
		for i := len(arr) - 1; i >= 0; i-- {
			countIndex := getCount(arr, i, int(shift), baseDigit)
			result[count[countIndex]-1] = arr[i]
			count[countIndex]--
		}

		for i := 0; i < len(arr); i++ {
			arr[i] = result[i]
		}
		//fmt.Println(result)
	}

	return result
}

func getCount(arr []int, i, shift, baseDigit int) int {
	return arr[i] / int(shift) % baseDigit
}

/* bucket sort:
when all elements are in a not very large range, we can split them in some sub-range - subset of the range, and then sort and merge
	1. find the max and min element in buckets;
	2. define bucket number, and calculate bucket range for every bucket;
	3. loop array, find the bucket for every element;
	4. loop buckets, sort for every buckets;
	5. merge all bucket and get the sorted results.
*/
func bucketSort(arr []int) []int {
	min, max := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

		if arr[i] < min {
			min = arr[i]
		}
	}

	bucketCount := 11
	delta := (max-min)/(bucketCount-1) + 1
	bucket := make([][]int, bucketCount)
	var index map[int]int = make(map[int]int)
	for i := 0; i < bucketCount; i++ {
		index[i] = 0
		bucket[i] = make([]int, 0)
	}

	for i := 0; i < len(arr); i++ {
		bucketIndex := (arr[i] - min) / delta

		bucket[bucketIndex] = append(bucket[bucketIndex], arr[i])
		index[bucketIndex]++
	}

	result := make([]int, 0)
	for i := 0; i < bucketCount; i++ {
		if index[i] > 0 {
			result = append(result, countingSort(bucket[i], min+i*delta, min+i*delta+delta)...)
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

func seedWithinRange(n int, bottom, up int) (arr []int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < n; i++ {
		arr = append(arr, r1.Intn(up-bottom)+bottom)
	}
	return
}
