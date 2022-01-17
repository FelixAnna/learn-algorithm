package simple

type ListNode struct {
	Val  int
	Next *ListNode
}

/** Add two numbers in a linklist: https://leetcode.com/problems/add-two-numbers/
numbers in linked list are in reverse order
1. add them from 2 list from begin  to end
2. if have carry after all done, need add additional node for carry
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	header := ListNode{Val: 0}

	pnode := &header
	carry := 0
	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		val := val1 + val2 + carry

		curVal := val % 10
		carry = val / 10

		pnode.Next = &ListNode{Val: curVal}
		pnode = pnode.Next
	}

	if carry > 0 {
		pnode.Next = &ListNode{Val: carry}
	}

	return header.Next
}
