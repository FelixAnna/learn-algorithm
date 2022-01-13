package search

import (
	"testing"

	"github.com/felixanna/algorithm-go/lib"
	"github.com/felixanna/algorithm-go/sort"
)

func TestBinarySearch_Exists(t *testing.T) {

	arr := lib.Seed(100)
	element := arr[3]
	sort.QuickSort(arr)

	index := -1
	if index = BinarySearch(arr, element); index < 0 {
		t.Fail()
	}
}

func TestBinarySearch_NotExists(t *testing.T) {

	arr := lib.Seed(100)
	element := -100
	sort.QuickSort(arr)

	index := -1
	if index = BinarySearch(arr, element); index >= 0 {
		t.Fail()
	}
}

func TestFindMin_One(t *testing.T) {

	length := 1
	arr := lib.Seed(length)

	_, element := FindMin(arr)

	if element != arr[0] {
		t.Fail()
	}
}

func TestFindMin_More(t *testing.T) {
	length := 100
	arr := lib.Seed(length)
	_, element := FindMin(arr)

	sort.QuickSort(arr)

	if element != arr[0] {
		t.Fail()
	}
}

func TestFindMax_One(t *testing.T) {

	length := 1
	arr := lib.Seed(length)

	_, element := FindMax(arr)

	if element != arr[0] {
		t.Fail()
	}
}

func TestFindMax_More(t *testing.T) {
	length := 100
	arr := lib.Seed(length)
	_, element := FindMax(arr)

	sort.QuickSort(arr)

	if element != arr[length-1] {
		t.Fail()
	}
}

func TestFindNthMin_One(t *testing.T) {

	length := 1
	arr := lib.Seed(length)

	_, element := FindNthMin(arr, 1)

	if element != arr[0] {
		t.Fail()
	}
}

func TestFindNthMin_More(t *testing.T) {
	length := 100
	nth := 5
	arr := lib.Seed(length)
	_, element := FindNthMin(arr, nth)

	sort.QuickSort(arr)

	if element != arr[nth-1] {
		t.Fail()
	}
}

func TestFindNthMin_Exception(t *testing.T) {
	length := 100
	nth := 500
	arr := lib.Seed(length)
	index, _ := FindNthMin(arr, nth)

	if index >= 0 {
		t.Fail()
	}
}
