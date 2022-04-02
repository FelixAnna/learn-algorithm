package backTracking

func GenerateParenthesis(n int) []string {
	result := make([]string, 0)

	result = backTracking(result, "", 0, 0, n)

	return result
}

func backTracking(result []string, p string, left, right, max int) []string {
	if left == max && right == max {
		result = append(result, p)
		return result
	}

	if left < max {
		result = backTracking(result, p+"(", left+1, right, max)
	}

	if right < left {
		result = backTracking(result, p+")", left, right+1, max)
	}

	return result
}

func Permute(nums []int) [][]int {
	results := [][]int{}
	backTrackingNums(nums, []int{}, &results)

	return results
}

func backTrackingNums(remainings []int, current []int, results *[][]int) {

	if len(remainings) == 0 {
		*results = append(*results, current)
	}

	for i, val := range remainings {
		backTrackingNums(
			append(append([]int{}, remainings[:i]...), remainings[i+1:]...), //append to new slice, avoid change existing slice
			append(current, val),
			results)
	}
}

/* n queen issue
 */
func NQueen(n int) [][]int {
	results := make([][]int, 0)

	findSolution([]int{}, n, 0, &results)
	return results
}

func findSolution(current []int, n int, i int, results *[][]int) {
	if i == n {
		*results = append(*results, current)
	}

	for j := 0; j < n; j++ {
		newcurrent := append(append([]int{}, current...), j)
		if !boundChecknQ(newcurrent, len(newcurrent)-1) {
			continue
		}

		findSolution(newcurrent, n, i+1, results)
	}
}

//check if last element is valid in current array
func boundChecknQ(current []int, c int) bool {
	for i := 0; i < len(current)-1; i++ { //compare to last element only, suppose current[0:n-1] already valid
		if current[i] == current[c] || //same column
			current[i]-current[c] == i-c || current[i]-current[c] == c-i { //diagonal
			return false
		}
	}

	return true
}

func SolveNQueens(n int) [][]string {
	results := make([][]string, 0)

	findSolutions(make([]int, 0), n, 0, &results)

	return results
}

func findSolutions(current []int, n, i int, results *[][]string) {
	if i == n {
		*results = append(*results, converToString(current))
		return
	}

	for j := 0; j < n; j++ {

		if !boundCheck(current, j) {
			continue
		}

		findSolutions(append(append([]int{}, current...), j), n, i+1, results)
	}
}

func boundCheck(current []int, k int) bool {
	for i := 0; i < len(current); i++ {
		if current[i] == k || //same column
			current[i]-k == len(current)-i || current[i]-k == i-len(current) {
			return false
		}
	}

	return true
}

func converToString(c []int) []string {
	results := make([]string, 0)
	for i := 0; i < len(c); i++ {
		line := ""
		for j := 0; j < len(c); j++ {
			if j == c[i] {
				line += "Q"
			} else {
				line += "."
			}
		}

		results = append(results, line)
	}

	return results
}

func Subsets(nums []int) [][]int {
	results := make([][]int, 0)
	results = append(results, []int{})

	findSubset(nums, []int{}, &results)
	return results
}

func findSubset(nums, current []int, results *[][]int) {
	if len(nums) == 0 {
		return
	}

	for i := 0; i < len(nums); i++ {
		newcurrent := append(append([]int{}, current...), nums[i])
		*results = append(*results, newcurrent)
		if i+1 <= len(nums) {
			findSubset(nums[i+1:], newcurrent, results)
		}
	}
}

//backtracking for this issue have performace issue
func uniquePaths(m int, n int) int {
	quantity := 0
	if m < n {
		m, n = n, m
	}

	if m == 1 && n == 1 {
		return 1
	}

	findPath(m-1, n-1, &quantity)
	return quantity
}

func findPath(i, j int, quantity *int) {
	if i == 0 && j == 0 {
		*quantity += 1
	}

	if i < 0 || j < 0 {
		return
	}

	findPath(i-1, j, quantity)
	findPath(i, j-1, quantity)
}
