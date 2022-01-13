package fib

import (
	"testing"
)

func TestFibRecursive_One(t *testing.T) {
	number := 1
	result := FibRecursive(number)

	if result != 1 {
		t.Fail()
	}
}

func TestFibRecursive_Five(t *testing.T) {
	number := 5
	result := FibRecursive(number)

	if result != 8 {
		t.Fail()
	}
}

func TestFibBottomUp_One(t *testing.T) {
	number := 1
	result := FibBottomUp(number)

	if result[number] != 1 {
		t.Fail()
	}
}

func TestFibBottomUp_Five(t *testing.T) {
	number := 5
	result := FibBottomUp(number)

	if result[number] != 8 {
		t.Fail()
	}
}

func TestFibBottomUp_More(t *testing.T) {
	number := 100
	result := FibBottomUp(number)

	if result[number] != 1298777728820984005 {
		t.Fail()
	}
}
