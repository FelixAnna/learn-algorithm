package dynamic

import "fmt"

func TestDynamicProgramming(n int, action int) {
	strategies := map[int]float64{
		1: 1,
		2: 5,
		3: 8,
		4: 9,
		5: 10,
		6: 17,
		7: 17,
		8: 20,
	}

	st, profit := CutRod(n, strategies)

	fmt.Println(n, st, profit)
}
