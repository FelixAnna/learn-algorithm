package lib

import (
	"math/rand"
	"time"
)

func Seed(n int) (arr []int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < n; i++ {
		arr = append(arr, r1.Intn(i+n))
	}

	return
}

func SeedWithinRange(n int, bottom, up int) (arr []int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < n; i++ {
		arr = append(arr, r1.Intn(up-bottom)+bottom)
	}
	return
}

func Fibonacci(n int64) int64 {
	if n <= 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
