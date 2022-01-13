package fib

import "github.com/felixanna/algorithm-go/lib"

func Recursive(n int64) int64 {
	if n <= 1 {
		return 1
	}

	return lib.Fibonacci(n-1) + lib.Fibonacci(n-2)
}

func BottomUp(n int64, results []int64) int64 {
	if n <= 1 {
		return 1
	}

	return results[n-1] + results[n-2]
}
