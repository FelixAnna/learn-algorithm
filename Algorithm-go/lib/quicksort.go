package lib

import (
	"fmt"
	"math/rand"
	"time"
)

func TestQuickSort(n int) {
	arr := seed(n)
	//fmt.Println(arr)
	fmt.Println(arr[100:115])

	start := time.Now()
	sort(arr, 0, n-1)
	cost := time.Now().UnixMilli() - start.UnixMilli()

	fmt.Println(arr[100:115])

	fmt.Println(cost)
	fmt.Println("sorted, checking")

	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			fmt.Println(arr[i-1], arr[i])
			break
		}
	}

	fmt.Println("done")
}

/* quick sort：
if start>=end, no need to sort(at most 1 element)
	1. random find a pivot in the slice;
	2. swap pivot element with start;  ##[pivot, e2, e3, ... ,e1, ... , ei]
	3. from start + 1 , to end (include), compare to start, set partitionIndex =  start :  ## [pivot (partition index), e2, e3, ... , ei]
		a. if x <= pivot, partitionIndex +=1; then swap partitionIndex with x;   e3<=pivot =>  ## [pivot, e3 (old partition index), e2 (new partition index), ... , ei]
		b. if bigger, do nothing	 e3>pivot  => ## [pivot, e2  (old / new partition index), e3, ... , ei]
	4. swap start with partitionIndex
	5. sort( arr, start, partition index - 1), sort(arr, partition index + 1, end)
*/
func sort(arr []int, start, end int) {

	//only have one element or no element
	if start >= end {
		return
	}

	pivotIndex := rand.Intn(end-start) + start

	swap(arr, start, pivotIndex)
	partitionIndex := start

	for i := start + 1; i <= end; i++ {
		if arr[i] <= arr[start] {
			partitionIndex += 1
			swap(arr, i, partitionIndex)
		}
	}

	swap(arr, start, partitionIndex)

	sort(arr, start, partitionIndex-1)
	sort(arr, partitionIndex+1, end)
}

func swap(arr []int, start, target int) {
	//fmt.Println(arr[start], arr[target])
	arr[start], arr[target] = arr[target], arr[start]

	//fmt.Println(arr)
}

func seed(n int) (arr []int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < n; i++ {
		arr = append(arr, r1.Intn(i+n))
	}

	return
}
