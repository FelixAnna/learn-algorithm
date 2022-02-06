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
