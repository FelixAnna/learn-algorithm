package sort

import (
	"errors"
	"fmt"
	"testing"

	"github.com/felixanna/algorithm-go/lib"
)

func TestQuickSort_Single(t *testing.T) {
	length := 1
	arr := lib.Seed(length)
	QuickSort(arr)

	if err := checkSorted(arr, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestQuickSort_More(t *testing.T) {
	length := 1000000
	arr := lib.Seed(length)
	QuickSort(arr)

	if err := checkSorted(arr, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestMergeSort_Single(t *testing.T) {
	length := 1
	arr := lib.Seed(length)
	result := MergeSort(arr)

	if err := checkSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestMergeSort_More(t *testing.T) {
	length := 1000000
	arr := lib.Seed(length)
	result := MergeSort(arr)

	if err := checkSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestCountingSort_Single(t *testing.T) {
	length := 1
	min := 1000
	max := 2000
	arr := lib.SeedWithinRange(length, min, max)
	result := CountingSort(arr, min, max)

	if err := checkSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestCountingSort_More(t *testing.T) {
	length := 1000000
	min := 1000
	max := 2000
	arr := lib.SeedWithinRange(length, min, max)
	result := CountingSort(arr, min, max)

	if err := checkSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestRadixSort_Single(t *testing.T) {
	length := 1
	arr := lib.Seed(length)
	result := RadixSort(arr)

	if err := checkSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestRadixSort_More(t *testing.T) {
	length := 1000000
	arr := lib.Seed(length)
	result := RadixSort(arr)

	if err := checkSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestBucketSort_Single(t *testing.T) {
	length := 1
	arr := lib.Seed(length)
	result := BucketSort(arr)

	if err := checkSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestBucketSort_More(t *testing.T) {
	length := 1000000
	arr := lib.Seed(length)
	result := BucketSort(arr)

	if err := checkSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func checkSorted(arr []int, n int) error {
	//fmt.Println("sorted, checking")
	if len(arr) != n {
		err := fmt.Sprintln("Length incorrect:", len(arr), n)
		return errors.New(err)
	}

	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			err := fmt.Sprintln("Failed:", arr[i-1], arr[i])
			return errors.New(err)
		}
	}

	return nil
}
