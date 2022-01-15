package main

import (
	"fmt"

	"github.com/felixanna/algorithm-go/dynamic"
)

func main() {
	fmt.Println("start")

	//lib.TestChan(40, 2)
	//lib.TestFib(100, 1)
	//lib.TestSearch(8, 2)
	dynamic.TestDynamicProgramming(1, 1)
	dynamic.TestDynamicProgramming(2, 1)
	dynamic.TestDynamicProgramming(3, 1)
	dynamic.TestDynamicProgramming(4, 1)
	dynamic.TestDynamicProgramming(5, 1)
	dynamic.TestDynamicProgramming(6, 1)
	dynamic.TestDynamicProgramming(700, 1)
	//fmt.Scanln()
}
