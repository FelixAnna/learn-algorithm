package search

import (
	"fmt"
	"time"

	"github.com/felixanna/algorithm-go/lib"
	"github.com/felixanna/algorithm-go/sort"
)

/* sort
ways to search:
	1. bin search in a sorted array
	2. find xth max / min element in an unsorted array
*/
func TestSearch(n int, action int) {
	fmt.Println("TestSearch:", n, action)
	arr := lib.Seed(n)
	var element int
	var index int = -1

	start := time.Now()
	switch action {
	case 1:
		{
			element = arr[n/2]
			arr = sort.RadixSort(arr)
			index = BinarySearch(arr, element)

		}
	case 2:
		{
			index, element = FindNthMin(arr, 5)
		}
	case 3:
		{
			index, element = FindMin(arr)
		}
	case 4:
		{
			index, element = FindMax(arr)
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
