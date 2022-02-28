package backTracking

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	input := []int{1, 2, 3}
	result := Permute(input)

	if len(result) != 6 || result[len(result)-1][0] != 3 || result[len(result)-1][1] != 2 || result[len(result)-1][2] != 1 {
		fmt.Println(result)
		t.Fail()
	}
}

func TestNQueen(t *testing.T) {
	n := 4
	res := NQueen(n)

	if len(res) != 2 {
		t.Fail()
	}
}

func TestNQueen_Eight(t *testing.T) {
	n := 8
	res := NQueen(n)

	if len(res) != 92 {
		t.Fail()
	}
}

func TestSolveNQueens(t *testing.T) {
	n := 4
	res := SolveNQueens(n)

	if len(res) != 2 {
		t.Fail()
	}
}

func TestSolveNQueens_Eight(t *testing.T) {
	n := 8
	res := SolveNQueens(n)

	if len(res) != 92 {
		t.Fail()
	}
}

func TestGenerateParenthesis(t *testing.T) {
	result := GenerateParenthesis(1)

	if result[0] != "()" {
		t.Fail()
	}
}

func TestGenerateParenthesis_More(t *testing.T) {
	result := GenerateParenthesis(4)

	if len(result) != 14 {
		t.Fail()
	}
}

func TestSubset(t *testing.T) {
	nums := []int{1, 3, 7}
	results := Subsets(nums)
	if len(results) != 8 {
		t.Fail()
	}
}
