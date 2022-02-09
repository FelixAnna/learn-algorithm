package fib

/* FibRecursive
use recursive way to calculate fibonacci top down,
	f(n) = f(n-1) + f(n-2)
contains many duplicate calculations

return f(n) only
*/
func FibRecursive(n int) int64 {
	if n <= 1 {
		return 1
	}

	return FibRecursive(n-1) + FibRecursive(n-2)
}

/* FibBottomUp
use bottom up way to calculate fib, this is dynamic programming algorithm
	1. define an array to keep result for fib mapping
	2. loop i from 0 to n:
		if i<=1, just set result[i] = 1
		otherwise: result[i] = result[i-1] + result[i-2]
return fib mappings from 0 - n
*/
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
	// if n <= 1 {
	// 	return 1
	// }

	return results[n-1] + results[n-2]
}
