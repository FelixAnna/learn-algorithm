package fib

func FibRecursive(n int) int64 {
	if n <= 1 {
		return 1
	}

	return FibRecursive(n-1) + FibRecursive(n-2)
}

func FibBottomUp(n int) []int64 {
	results := make([]int64, n+1)

	for i := 0; i <= n; i++ {
		if i <= 1 {
			results[i] = 1
			continue
		}

		results[i] = fibBottomUpInternal(int64(i), results)
		//fmt.Println("chan:", i, results[i])
	}

	return results
}

func fibBottomUpInternal(n int64, results []int64) int64 {
	if n <= 1 {
		return 1
	}

	return results[n-1] + results[n-2]
}
