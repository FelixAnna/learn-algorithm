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
