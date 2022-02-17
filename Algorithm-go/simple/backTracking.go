package simple

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
	result := backTrackingNums(nums, []int{})

	return result
}

func backTrackingNums(remainings []int, current []int) [][]int {

	if len(remainings) == 0 {
		return [][]int{current}
	}

	results := make([][]int, 0)
	for i, val := range remainings {
		newcurrent := append(current, val)
		remainings[0], remainings[i] = remainings[i], remainings[0] //changing the current slice

		remainingNums := make([]int, 0) // create new slice for avoid overwrite by other recursive: otherwise
		//the recursive func will change the same slice - remainingNums := remainings[1:]
		remainingNums = append(remainingNums, remainings[1:]...)

		results = append(results, backTrackingNums(remainingNums, newcurrent)...)
	}

	return results
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
		if !boundCheck(newcurrent, len(newcurrent)-1) {
			continue
		}

		findSolution(newcurrent, n, i+1, results)
	}
}

//check if last element is valid in current array
func boundCheck(current []int, c int) bool {
	for i := 0; i < len(current)-1; i++ { //compare to last element only, suppose current[0:n-1] already valid
		if current[i] == current[c] || //same column
			current[i]-current[c] == i-c || current[i]-current[c] == c-i { //diagonal
			return false
		}
	}

	return true
}
