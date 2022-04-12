package search

import "sync"

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
	allOrders := make([]Order, 0)
	for _, ch := range channels {
		orders := <-ch
		for _, order := range orders {
			if order.IsCanceledPaid() {
				allOrders = append(allOrders, order)
			}
		}
		close(ch)
	}

	return allOrders
}

func (o *Order) IsCanceledPaid() bool {
	return o.Paid && o.Status == Canceled
}

/* Given orderIds, the func will query mock api service and then get all orders with status = Canceled, Paid = true
 */
func GetCanceledPaidOrders2(orderIds []string) []Order {
	chunkList := chunkBy(orderIds, batchSize)
	group := sync.WaitGroup{}
	ch := make(chan []Order)

	defer close(ch)

	for i := 0; i < len(chunkList); i++ {
		group.Add(1)
		go func() {
			orders := dao.GetOrders(orderIds)
			result := make([]Order, 0)
			for _, order := range orders {
				if order.IsCanceledPaid() {
					result = append(result, order)
				}
			}

			ch <- result
			group.Done()
		}()
	}

	orders := make([]Order, 0)
	for i := 0; i < len(chunkList); i++ {
		orders = append(orders, <-ch...)
	}

	group.Wait()

	return orders
}
