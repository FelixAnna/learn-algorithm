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
package channel

import (
	"fmt"
	"math/rand"
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

func waitGroupTest(wg *sync.WaitGroup, n int) {
	defer wg.Done()

	time.Sleep(time.Second * time.Duration(n))
}

//wait for goroutine with chan: goroutine - chan<-val, main routine - <- chan, main routine will wait for chan output
func goRoutineChan(n int) {
	done := make(chan bool)
	go chanTest(n, done)

	<-done
}

func chanTest(n int, done chan<- bool) {

	time.Sleep(time.Second * time.Duration(n))

	done <- true
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

		if !y1 && !y3 {
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

//fanIn 2 channel with timeout example
func fanIn(input1, input2 <-chan string, timeOutInSec int) <-chan string {
	output := make(chan string)
	var timeout = 0
	var totalTimeTimeout = time.After(1 * time.Second)
	go func() {
		for {
			select {
			case s1 := <-input1:
				output <- s1
			case s2 := <-input2:
				output <- s2
			case <-time.After(time.Duration(timeOutInSec) * time.Millisecond):
				output <- fmt.Sprintf("Timeout %v", timeout)
			case <-totalTimeTimeout:
				output <- fmt.Sprintf("Total Timeout %v", timeout)
				//return
			}
		}
	}()

	return output
}

func makeChan(name string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%v %v", name, i)
			time.Sleep(time.Duration(time.Duration(rand.Intn(900)) * time.Millisecond))
		}
	}()

	return ch
}

func ProcessChans() {
	f, a := makeChan("felix"), makeChan("anna")

	ch := fanIn(f, a, 300)
	fmt.Println("Start printing...")
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Finised printing.")
}
