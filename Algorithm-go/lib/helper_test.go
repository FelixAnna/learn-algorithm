package lib

import (
	"testing"
)

func TestSeed(t *testing.T) {
	n := 100
	var result interface{} = Seed(n)

	re, ok := result.([]int)
	if !ok || len(re) != n {
		t.Fail()
	}
}

func TestSeedWithinRange(t *testing.T) {
	n := 100
	bottom, up := 1000, 2000
	result := SeedWithinRange(n, bottom, up)

	if len(result) != n {
		t.Fatalf("Length not correct: %v, expected: %v", len(result), n)
	}

	for i := 0; i < n; i++ {
		if result[i] < bottom || result[i] > up {
			t.Fatalf("Value not correct: %v, expected: [%v, %v]", result[i], bottom, up)
		}
	}
}

func TestCheckSorted_OK(t *testing.T) {
	n := 3
	arr := []int{1, 2, 3}
	err := CheckSorted(arr, n)

	if err != nil {
		t.Fail()
	}
}

func TestCheckSorted_WhenNotSameLength_ThenError(t *testing.T) {
	n := 3
	arr := []int{1, 2, 3, 2}
	err := CheckSorted(arr, n)

	if err == nil {
		t.Fail()
	}
}

func TestCheckSorted_WhenUnSorted_ThenError(t *testing.T) {
	n := 4
	arr := []int{1, 2, 3, 2}
	err := CheckSorted(arr, n)

	if err == nil {
		t.Fail()
	}
}
