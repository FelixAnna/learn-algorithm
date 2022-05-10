package dynamic

import (
	"testing"
)

var strategies map[int]float64

func init() {
	strategies = make(map[int]float64)
	strategies[1] = 2.5
	strategies[2] = 5.5
	strategies[3] = 8.05
	strategies[5] = 15
}

func TestCutRod_One(t *testing.T) {
	length := 1
	_, profit := CutRod(length, strategies)

	if profit != strategies[1] {
		t.Fail()
	}
}

func TestCutRod_Five(t *testing.T) {
	length := 5
	_, profit := CutRod(length, strategies)

	if profit != strategies[5] {
		t.Fail()
	}
}

func TestClimbStair_One(t *testing.T) {
	length := 1
	nmap := climbingStair(length)

	if nmap[length] != 1 {
		t.Fail()
	}
}

func TestClimbStair_Five(t *testing.T) {
	length := 5
	nmap := climbingStair(length)

	if nmap[length] != 8 {
		t.Fail()
	}
}

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ret := MaxSubArray(nums)

	if ret != 6 {
		t.Fail()
	}
}

func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	ret := CanJump(nums)

	if !ret {
		t.Fail()
	}
}

func TestUniquePaths(t *testing.T) {
	ret := UniquePaths(3, 7)
	if ret != 28 {
		t.Fail()
	}
}
func TestNumDecodings_NoWay(t *testing.T) {
	ret := NumDecodings("06")
	if ret != 0 {
		t.Fail()
	}
}

func TestNumDecodings_HaveWay(t *testing.T) {
	ret := NumDecodings("12")
	if ret != 2 {
		t.Fail()
	}
}

