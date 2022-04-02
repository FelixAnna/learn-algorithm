package search

import (
	"sync"
)

const batchSize = 1000

type OrderStatus string

const (
	New        OrderStatus = "New"
	Processing             = "Processing"
	Done                   = "Done"
	Canceled               = "Canceled"
)

type Order struct {
	Id     string
	Status OrderStatus
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
		status := New
		paid := true

		if i%3 == 0 {
			status = Canceled
		}

		if i%2 == 0 {
			paid = true
		}

		orderItem := Order{Id: orderIds[i], Status: status, Paid: paid}
		results[i] = orderItem
	}

	return results
}

/*** dao***/

var dao = MockOrderDAO{}

/* Given orderIds, the func will query mock api service and then get all orders with status = Canceled, Paid = true
 */
func GetCanceledPaidOrders(orderIds []string) []Order {
	chunkList := chunkBy(orderIds, batchSize)

	var ch chan []Order = make(chan []Order)
	queryOrders(chunkList, ch)

	canceledPaidOrders := filterOrders(ch)
	return canceledPaidOrders
}

func queryOrders(chunkList [][]string, ch chan []Order) {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(chunkList))
	go func() {
		for _, chunk := range chunkList {
			go getOrder(chunk, ch, waitGroup)
		}

		waitGroup.Wait()
		close(ch)
	}()
}

func filterOrders(ch <-chan []Order) []Order {
	allOrders := make([]Order, 0)
	for orders := range ch {
		for _, order := range orders {
			if order.IsCanceledPaid() {
				allOrders = append(allOrders, order)
			}
		}
	}
	return allOrders
}

func (o *Order) IsCanceledPaid() bool {
	return o.Paid && o.Status == Canceled
}

func getOrder(orderIds []string, ch chan<- []Order, waitGroup sync.WaitGroup) {
	orders := dao.GetOrders(orderIds)
	ch <- orders
	waitGroup.Done()
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
