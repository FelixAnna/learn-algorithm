package dynamic

import "fmt"

func TestDynamicProgramming(n int, action int) {
	strategies := make(map[int]float64)
	strategies[1] = 2.5
	strategies[2] = 5.5
	strategies[3] = 8.05
	strategies[5] = 15
	st, profit := CutRod(n, strategies)

	fmt.Println(n, st, profit)
}
