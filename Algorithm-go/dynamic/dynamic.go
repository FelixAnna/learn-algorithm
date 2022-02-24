package dynamic

import (
	"github.com/felixanna/algorithm-go/sort"
)

type cutSolution struct {
	profit float64

	cut int
	len int

	next *cutSolution
}

/* cut rod problem:
given a cut stratergy with profit, find the cut stratergy with maxinum profit
	1. define a map for the maxinum profit , cut stratergies for a variable x
	2. from k = 1 to n, do following:
		a. f(0) = 0
		b. f(k) = maxinum (f(k-1)+st[1], f(k-2) + st[2], .... , f(k-x)+st[x])  # x is the possible cut ways, k >=x in the loop
	3. return f(n), and path if necessary

*/
func CutRod(n int, st map[int]float64) (map[int]int, float64) {
	bst := make(map[int]*cutSolution)

	//get sorted keys
	keys := make([]int, 0, len(st))
	for k := range st {
		keys = append(keys, k)
	}
	sort.QuickSort(keys)

	bst[0] = &cutSolution{0, 0, 0, nil}

	for i := 1; i <= n; i++ {

		solution := cutSolution{}
		for _, s := range keys {
			if i < s {
				continue
			}

			currentProfit := st[s] + bst[i-s].profit
			if currentProfit > solution.profit {
				solution = cutSolution{profit: currentProfit, len: i, cut: s, next: bst[i-s]}
			}
		}

		bst[i] = &solution
	}

	//combine roadmap
	stratergies := make(map[int]int, 0)
	curBst := bst[n]

	for curBst != nil {
		if curBst.cut > 0 {
			stratergies[curBst.cut] += 1
		}
		curBst = curBst.next
	}

	return stratergies, bst[n].profit
}

/* Climbing stair issue:
You can only climb 1 or 2 stairs one time, this is actually an fibonacci, to climb to n stairs:
	1. you can climb to n-1 then climb just 1 star,
	2. you can also climb to n-2, then just climb 2 stairs
because the last step is not same, even f(n-1) contains f(n-2) with one step more, they are still conains different ways
	so: f(n) = f(n-1) + f(n-2)
*/
func climbingStair(n int) map[int]int {
	results := make(map[int]int, n)

	for i := 1; i <= n; i++ {
		if i <= 2 {
			results[i] = i
			continue
		}

		results[i] = results[i-1] + results[i-2]
	}

	return results
}

/* IsMatch
Test is a string matches regular pattern: contains: .*0-9a-z only
Solution
As we have .* , it might not be the last part, but it can match anything, so a simple loop to check match is not feasible
This is another dp issue, for given string s and pattern p, we can start by check s[i] matches p[j] or not,
and save this in a 2 factor array dp[m+1][n+1], dp[0][0]=true, since empty s matches empty pattern.

dp[i][j] means s[:i] matches p[:j] - 1st i-1 element matches 1st j-1 pattern

1. if s[i-1] == p[j-1] : dp[i][j] = dp[i-1][j-1] && true (since last part same, if previous part match, then new part matches to new pattern)
2. if p[j-1] == '.’ 	：dp[i][j] = dp[i-1][j-1] && true
3. if p[j-1] == '*'
	3.1 if p[j-2] != '.' && if s[i-1] != p[j-2] : dp[i][j] = dp[i][j-2] (consider x* as empty matching)

	3.1 if p[j-2] == '.' || p[j-2] == s[i-1]  //.*  a.*
		dp[i][j] = dp[i-1][j-2]
	3.2

*/
func IsMatch(s string, p string) bool {
	return false
}

/* Maximum Subarray
find contiguous subarry who have maximum value
solution:
always keep the maxinum sum till index i in array,
maximum of the whole array is one of the maximum sum of index i in array
this is so called: Kadane's algorithm
*/
func MaxSubArray(nums []int) int {
	maxTillNow, max := nums[0], nums[0]

	for _, val := range nums[1:] {
		if maxTillNow > 0 {
			maxTillNow = val + maxTillNow
		} else {
			maxTillNow = val
		}
		if maxTillNow > max {
			max = maxTillNow
		}
	}

	return max
}
