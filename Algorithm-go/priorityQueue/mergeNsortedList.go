package priorityqueue

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

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
2. use priority queue
*/
