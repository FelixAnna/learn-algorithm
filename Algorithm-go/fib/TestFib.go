package fib

import (
	"fmt"
	"time"
)

/* fibonacci:
2 ways to get fib:
	1 - bottom up calculate fibonacci, very efficient
	2 - recursive calculate fibonacci, normally not bigger than 50
*/
func TestFib(n int, action int) {
	fmt.Println("TestFib:", n, action)
	start := time.Now()
	switch action {
	case 1:
		{
			results := FibBottomUp(n)
			fmt.Println(results[n-1] + results[n-2])
		}
	case 2:
		{
			done := make(chan bool)
			go func() {
				a := FibRecursive(n) // 0 ~ n : n+1 level
				fmt.Println(a)
				done <- true
			}()

			<-done

		}
	default:
		fmt.Println("not support")
	}

	cost := time.Now().UnixMilli() - start.UnixMilli()
	fmt.Println("cost: ", cost)

	fmt.Println("done")
}
