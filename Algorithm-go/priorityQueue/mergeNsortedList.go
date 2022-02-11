package priorityqueue

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
import (
	"container/heap"
	_ "container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1. divide and conqur
*/
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	ll := MergeKLists(lists[:mid])
	rl := MergeKLists(lists[mid:])

	return merge2SortedList(ll, rl)
}

func merge2SortedList(ll, rl *ListNode) *ListNode {
	lh, rh := ll, rl

	head := &ListNode{}
	current := head

	for lh != nil && rh != nil {
		if lh.Val < rh.Val {
			current.Next = &ListNode{Val: lh.Val}
			lh = lh.Next
		} else {
			current.Next = &ListNode{Val: rh.Val}
			rh = rh.Next
		}

		current = current.Next
	}

	if lh != nil {
		current.Next = lh
	}

	if rh != nil {
		current.Next = rh
	}

	return head.Next
}

/*
2. priority queue based on heap (in array)
*/
type IntHeap []int

/*
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}
func (h IntHeap) Len() int
func (h IntHeap) Less(i, j int) bool
func (h IntHeap) Swap(i, j int)
*/
func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h

	length := old.Len()
	last := old[length-1]
	*h = old[0 : length-1]

	return last
}

func MergeKLists_PriorityQUeue(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	//build IntHeap
	hp := &IntHeap{}
	heap.Init(hp)

	for _, list := range lists {
		current := list
		for current != nil {
			heap.Push(hp, current.Val)
			current = current.Next
		}
	}

	//fmt.Println(hp.Len(), (*hp))

	head := &ListNode{}
	current := head
	for hp.Len() > 0 {
		value := heap.Pop(hp)
		current.Next = &ListNode{Val: value.(int)}
		current = current.Next
	}
	return head.Next
}
