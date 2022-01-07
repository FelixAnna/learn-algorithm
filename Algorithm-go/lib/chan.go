/*
channel: chan is a bidirectional pipeline used in goroutine for comminication or exchange data.
goroutine: goroutine is a concurreny technology in go, it build an abstract layer on os level thread,
user use goroutine not directly manage thread, user manage goroutine,
go runtime would scheduler goroutine run on system thread damymically, it more efficient than os thread
1. main routine would not wait for goroutine by default
2. sync.WaitGroup can be used to wait goroutine finish
3. main routine can wait for goroutine by wait a chan's output, sub routine would fill the chan when finished
4. buffered chan: buffer x items in chan before the chan is blocked for consum in other goroutine (include main goroutine)
5. chan with select: if wait for multiple chan in different select branches, would run either chan is ready, usually in for loop
6. chan with chan
*/
package lib

import (
	"fmt"
	"sync"
	"time"
)

//action: 1 - sync + chan, 2 - buffered chan, 3 - select on multiple chan, 4 - chan & chan
func TestChan(n int, action int) {
	fmt.Println("TestChan:", n, action)
	switch action {
	case 1:
		{
			done := make(chan bool)
			go func() {
				goRoutineWaitGroup(n)
				done <- true
			}()

			go func() {
				goRoutineChan(n)
				done <- true
			}()

			<-done
			<-done
		}
	case 2:
		bufferedChan(n)
	case 3:
		selectChan(n)
	case 4:
		chanWithChan(n)
	default:
		fmt.Println("not support")
	}

	fmt.Println("done")
}

//wait for goroutine with sync.WaitGroup: Add(x), Done(), Wait()
func goRoutineWaitGroup(n int) {
	wg := sync.WaitGroup{}
	wg.Add(1) //wait for 1 goroutine

	go waitGroupTest(&wg, n)
	wg.Wait()
}

func waitGroupTest(wg *sync.WaitGroup, n int) []int64 {
	defer wg.Done()

	results := make([]int64, n)
	for i := 0; i < n; i++ {
		results[i] = fibonacci(int64(i))
		fmt.Println("waitGroup:", i, results[i])
	}

	return results
}

func fibonacci(n int64) int64 {
	if n <= 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

//wait for goroutine with chan: goroutine - chan<-val, main routine - <- chan, main routine will wait for chan output
func goRoutineChan(n int) {
	done := make(chan bool)
	go chanTest(n, done)

	<-done
}

func chanTest(n int, done chan<- bool) {

	if n <= 1 {
		fmt.Println("chan:", 0, 1)
		fmt.Println("chan:", 1, 1)
		return
	}

	results := make([]int64, n)

	for i := 0; i < n; i++ {
		if i <= 1 {
			results[i] = 1
			continue
		}

		results[i] = fibonacciBottomUp(int64(i), results)
		fmt.Println("chan:", i, results[i])
	}

	done <- true
}

func fibonacciBottomUp(n int64, results []int64) int64 {
	if n <= 1 {
		return 1
	}

	return results[n-1] + results[n-2]
}

func bufferedChan(n int) {
	fmt.Println("buffered chan start")
	bch := make(chan int, 2)

	go func() {
		i := 0
		for i < n {

			bch <- i
			bch <- i * i
			bch <- i * i * i
			i++
			time.Sleep(time.Millisecond * 100)
		}

		close(bch)
	}()

	for i := range bch {
		println(i)
	}

	fmt.Println("buffered chan done")
}

func selectChan(n int) {
	sl := make([]int, n)
	i := 0
	for i < n {
		sl[i] = i
		i++
	}

	ch1, ch3 := make(chan int), make(chan int)
	go func() {
		j := 0
		for j < n {
			fmt.Println("3x:", sl[j])
			ch3 <- sl[j]
			j += 3
			time.Sleep(time.Millisecond * 100)
		}

		close(ch3)
	}()

	go func() {
		j := 0
		for j < n {
			fmt.Println("x:", sl[j])
			ch1 <- sl[j]
			j += 1
			time.Sleep(time.Millisecond * 100)
		}

		close(ch1)
	}()

	for {
		x1, y1, x3, y3 := 0, false, 0, false
		select {
		case x1, y1 = <-ch1:
			fmt.Println(x1, y1)
		case x3, y3 = <-ch3:
			fmt.Println(x3, y3)
		}

		if y1 == false && y3 == false {
			fmt.Println("Done")
			break
		}
	}
}

func chanWithChan(n int) {
	inch, outch := make(chan int), make(chan int)

	go func() {
		i := 0
		for i < n {
			inch <- i
			i++
		}

		close(inch)
	}()

	go swapChan(inch, outch)

	for a := range outch {
		fmt.Println(a)
	}

	fmt.Println("done")
}

//chan<- only allow write into chan,  <-chan only allow read from chan
func swapChan(inch <-chan int, outch chan<- int) {
	for a := range inch {
		outch <- a * a
	}

	close(outch)
}
