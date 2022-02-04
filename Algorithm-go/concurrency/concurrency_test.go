package concurrency

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPrintInOrder_Reverse(t *testing.T) {
	po := PrintInOrderGo{
		c1: make(chan bool),
		c2: make(chan bool),
		c3: make(chan bool),
	}

	f := func(i int) string {
		result := strconv.FormatInt(int64(i), 10)
		return result
	}

	go po.Third(f)
	go po.Second(f)
	go po.First(f)

	<-po.c3

	if po.Output != "123" {
		fmt.Println(po.Output)
		t.Fail()
	}
}

func TestPrintInOrder_AnyOrder(t *testing.T) {
	po := PrintInOrderGo{
		c1: make(chan bool),
		c2: make(chan bool),
		c3: make(chan bool),
	}

	f := func(i int) string {
		result := strconv.FormatInt(int64(i), 10)
		return result
	}

	go po.Second(f)
	go po.Third(f)
	go po.First(f)

	<-po.c3

	if po.Output != "123" {
		fmt.Println(po.Output)
		t.Fail()
	}
}
