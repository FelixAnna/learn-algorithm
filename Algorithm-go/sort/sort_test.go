package sort

import (
	"testing"

	"github.com/felixanna/algorithm-go/lib"
)

func TestQuickSort_Single(t *testing.T) {
	length := 1
	arr := lib.Seed(length)
	QuickSort(arr)

	if err := lib.CheckSorted(arr, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestQuickSort_More(t *testing.T) {
	length := 1000000
	arr := lib.Seed(length)
	QuickSort(arr)

	if err := lib.CheckSorted(arr, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestMergeSort_Single(t *testing.T) {
	length := 1
	arr := lib.Seed(length)
	result := MergeSort(arr)

	if err := lib.CheckSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestMergeSort_More(t *testing.T) {
	length := 1000000
	arr := lib.Seed(length)
	result := MergeSort(arr)

	if err := lib.CheckSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestCountingSort_Single(t *testing.T) {
	length := 1
	min := 1000
	max := 2000
	arr := lib.SeedWithinRange(length, min, max)
	result := CountingSort(arr, min, max)

	if err := lib.CheckSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestCountingSort_More(t *testing.T) {
	length := 1000000
	min := 1000
	max := 2000
	arr := lib.SeedWithinRange(length, min, max)
	result := CountingSort(arr, min, max)

	if err := lib.CheckSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestRadixSort_Single(t *testing.T) {
	length := 1
	arr := lib.Seed(length)
	result := RadixSort(arr)

	if err := lib.CheckSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestRadixSort_More(t *testing.T) {
	length := 1000000
	arr := lib.Seed(length)
	result := RadixSort(arr)

	if err := lib.CheckSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestBucketSort_Single(t *testing.T) {
	length := 1
	arr := lib.Seed(length)
	result := BucketSort(arr)

	if err := lib.CheckSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}

func TestBucketSort_More(t *testing.T) {
	length := 1000000
	arr := lib.Seed(length)
	result := BucketSort(arr)

	if err := lib.CheckSorted(result, length); err != nil {
		t.Fatalf(`Failed to sort, error: %v`, err)
	}
}
