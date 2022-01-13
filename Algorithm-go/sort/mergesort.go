package sort

/* MergeSort
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
func MergeSort(arr []int) []int {
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
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

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
