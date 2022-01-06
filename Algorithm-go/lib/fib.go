package lib

import (
	"fmt"
	"time"
)

//action: 1 - recursive calculate fibonacci, normally not bigger than 50,  2 - bottom up calculate fibonacci, very efficient
func TestFib(n int, action int) {
	start := time.Now()
	switch action {
	case 1:
		{
			done := make(chan bool)
			go func() {
				a := recursive(int64(n))
				fmt.Println(a)
				done <- true
			}()

			<-done

		}
	case 2:
		{
			results := make([]int64, n)
			for i := 0; i < n; i++ {
				if i <= 1 {
					results[i] = 1
					continue
				}

				results[i] = bottomUp(int64(i), results)
			}

			fmt.Println(results[n-1] + results[n-2])
		}
	default:
		fmt.Println("not support")
	}

	cost := time.Now().UnixMilli() - start.UnixMilli()
	fmt.Println(cost)
}

func recursive(n int64) int64 {
	if n <= 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func bottomUp(n int64, results []int64) int64 {
	if n <= 1 {
		return 1
	}

	return results[n-1] + results[n-2]
}
