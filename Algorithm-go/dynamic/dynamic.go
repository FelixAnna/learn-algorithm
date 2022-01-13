package dynamic

import (
	"github.com/felixanna/algorithm-go/sort"
)

/* cut rod problem:
given a cut stratergy with profit, find the cut stratergy with maxinum profit
	1. define a map for the maxinum profit , cut stratergies for a varible x
	2. from k = 1 to n, do following:
		a. f(0) = 0
		b. f(k) = maxinum (f(k-1)+st[1], f(k-2) + st[2], .... , f(k-x)+st[x])  # x is the possible cut ways, k >=x in the loop
	3. return f(n), and path if necessary

*/
type cutSolution struct {
	profit    float64
	stratergy int
}

func CutRod(n int, st map[int]float64) ([]int, float64) {
	bst := make(map[int]*cutSolution)

	//get sorted keys
	keys := make([]int, 0, len(st))
	for k := range st {
		keys = append(keys, k)
	}
	sort.Quicksort(keys, 0, len(keys)-1)

	bst[0] = &cutSolution{0, 0}

	for i := 1; i <= n; i++ {

		solution := cutSolution{}
		for _, s := range keys {
			if i < s {
				continue
			}

			currentProfit := st[s] + bst[i-s].profit
			if currentProfit > solution.profit {
				solution = cutSolution{profit: currentProfit, stratergy: s}
			}
		}

		bst[i] = &solution
	}

	//combine stratergy
	stratergies := make([]int, 0)
	i := n
	for i > 0 {
		stratergies = append(stratergies, bst[i].stratergy)
		i -= bst[i].stratergy
	}

	return stratergies, bst[n].profit
}
