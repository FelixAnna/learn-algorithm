package main

import (
	"fmt"

	"github.com/felixanna/algorithm-go/search"
)

func main() {
	orderIds := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	orders := search.GetCanceledPaidOrders(orderIds)

	fmt.Println(orders)
	//example.MathCheck(10, 1, 50, false)
}
