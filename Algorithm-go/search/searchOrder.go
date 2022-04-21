package search

import (
	"fmt"
	"sync"
)

const batchSize = 1000

const (
	Created int = iota
	Processing
	Done
	Canceled
)

type Order struct {
	Id     string
	Status int
	Paid   bool
}

/*** dao***/
type OrderDAO interface {
	GetOrders([]string) []Order
}

type MockOrderDAO struct {
}

func (dao *MockOrderDAO) GetOrders(orderIds []string) []Order {
	if len(orderIds) == 0 {
		return []Order{}
	}

	results := make([]Order, len(orderIds))
	for i := 0; i < len(orderIds); i++ {
		status := Created
		paid := false

		if i%3 == 0 {
			status = Canceled
		}

		if i%4 == 0 {
			paid = true
		}

		orderItem := Order{Id: orderIds[i], Status: status, Paid: paid}
		results[i] = orderItem
	}

	return results
}

/*** dao***/

var dao OrderDAO

func init() {
	dao = &MockOrderDAO{}
}

/* Given orderIds, the func will query mock api service and then get all orders with status = Canceled, Paid = true
 */
func GetCanceledPaidOrders(orderIds []string) []Order {
	chunkList := chunkBy(orderIds, batchSize)
	var channels []chan []Order = make([]chan []Order, len(chunkList))

	queryOrders(chunkList, channels)

	return filterOrders(channels)
}

func chunkBy(items []string, size int) (results [][]string) {
	for i := 0; i < len(items); i = i + size {

		end := i + size
		if len(items) < i+size {
			end = len(items)
		}

		results = append(results, items[i:end])
	}

	return
}

func queryOrders(chunkList [][]string, channels []chan []Order) {
	for i := 0; i < len(chunkList); i++ {
		channels[i] = make(chan []Order)
		go getOrder(chunkList[i], channels[i])
	}
}

func getOrder(orderIds []string, ch chan<- []Order) {
	orders := dao.GetOrders(orderIds)
	ch <- orders
}

func filterOrders(channels []chan []Order) []Order {
	targetOrders := make([]Order, 0)
	for _, ch := range channels {
		orders := <-ch
		for _, order := range orders {
			if order.IsCanceledPaid() {
				targetOrders = append(targetOrders, order)
			}
		}
		close(ch)
	}

	return targetOrders
}

func (o *Order) IsCanceledPaid() bool {
	return o.Paid && o.Status == Canceled
}

//can i reuse one channel for diff goroutines

func GetCanceledPaidOrders2(orderIds []string) []Order {
	chunkList := chunkBy(orderIds, batchSize)

	wg := sync.WaitGroup{}
	wg.Add(len(chunkList))

	var och chan []Order = make(chan []Order)
	defer close(och)

	for i := 0; i < len(chunkList); i++ {
		go func(chunkOrders []string, ch chan<- []Order) {
			orders := dao.GetOrders(orderIds)
			fmt.Println("Original:", orders)
			results := make([]Order, 0)
			for _, o := range orders {
				if o.IsCanceledPaid() {
					results = append(results, o)
				}
			}
			ch <- results
			wg.Done()

		}(chunkList[i], och)
	}

	canceledOrders := make([]Order, 0)
	for i := 0; i < len(chunkList); i++ {
		cos := <-och
		fmt.Println("Canceled:", cos)
		canceledOrders = append(canceledOrders, cos...)
	}

	wg.Wait()
	return canceledOrders
}
