package main

import (
	"fmt"

	"github.com/felixanna/algorithm-go/simple"
)

func main() {
	cmds := []string{
		"C 10 20",
		"N 3 4 6",
		"N 4 5 6",
		"S 3 4 4 5 5 6",
	}
	result := simple.Spreadsheet(cmds)
	fmt.Println(result[5][6])
}
