package sort

import "math"

/* CountingSort
suppose all elements in an array are in a relative small range 0~k， then we can use counting sort to sort the array in O(n+k)
steps:
	1. define counting array[k+1] : ca[k+1]     				example: [0, 0, 0]
	2. scan all elements in input array, set: ca[input[i]]++    example: [1, 3, 2]
	3. scan ca from 0 - k: make ca[i]=ca[i-1] + ca[i] 			example: [1, 4, 6]
	4. create a new array as the input array len: result array;
	5. scan input array (from last to first), result[ca[input[i]] - 1] = input[i]， set: ca[input[i]] = ca[input[i]] - 1
*/
func CountingSort(arr []int, bottom, up int) []int {
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

/* RadixSort
when element are positive numbers and distributed a little evenly in a large range, we can use radix sort:
step:
	1. define the max bit in numbers, and the base bit we use for split
	2. get the base digit we used for counting: baseDigit = 1 << baseBit
	3. loop for largest maxBit/baseBit times:
		a. sort element on k th digit with counting sort
		b. copy sorted array to original array as new input for next round loop
*/
func RadixSort(arr []int) []int {
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

		for i := len(arr) - 1; i >= 0; i-- {
			countIndex := getCount(arr, i, int(shift), baseDigit)
			result[count[countIndex]-1] = arr[i]
			count[countIndex]--
		}

		for i := 0; i < len(arr); i++ {
			arr[i] = result[i]
		}
	}

	return result
}

func getCount(arr []int, i, shift, baseDigit int) int {
	return arr[i] / int(shift) % baseDigit
}

/* BucketSort
when all elements are in a not very large range, we can split them in some sub-range - subset of the range, and then sort and merge
	1. find the max and min element in buckets;
	2. define bucket number, and calculate bucket range for every bucket;
	3. loop array, find the bucket for every element;
	4. loop buckets, sort for every buckets;
	5. merge all bucket and get the sorted results.
*/
func BucketSort(arr []int) []int {
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
			result = append(result, CountingSort(bucket[i], min+i*delta, min+i*delta+delta)...)
		}
	}

	return result
}
