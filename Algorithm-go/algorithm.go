package main

import (
	"fmt"

	"github.com/felixanna/algorithm-go/lib"
)

func main() {
	fmt.Println("start")

	//lib.TestChan(40, 2)
	//lib.TestFib(100, 1)
	//lib.TestSearch(8, 2)
	lib.TestDynamicProgramming(1, 1)
	lib.TestDynamicProgramming(2, 1)
	lib.TestDynamicProgramming(3, 1)
	lib.TestDynamicProgramming(4, 1)
	lib.TestDynamicProgramming(5, 1)
	lib.TestDynamicProgramming(6, 1)
	lib.TestDynamicProgramming(7, 1)
	//fmt.Scanln()
}
