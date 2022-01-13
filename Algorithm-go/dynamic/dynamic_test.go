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
