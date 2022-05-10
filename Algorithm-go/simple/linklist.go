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

func MergeTwoLists(h1 *ListNode, h2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	result := head

	val := 0
	for h1 != nil || h2 != nil {
		if h1 != nil && h2 != nil {
			if h1.Val < h2.Val {
				val = h1.Val
				h1 = h1.Next
			} else {
				val = h2.Val
				h2 = h2.Next
			}
		} else if h1 != nil {
			val = h1.Val
			h1 = h1.Next
		} else if h2 != nil {
			val = h2.Val
			h2 = h2.Next
		} else {
			break
		}

		result.Next = &ListNode{val, nil}
		result = result.Next
	}

	return head.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	f1, f2 := head, head
	index := n

	for f1.Next != nil { //last element
		f1 = f1.Next
		if index == 0 {
			f2 = f2.Next

		} else {
			index--
		}
	}

	if index == 1 { //same as length
		head = head.Next
	} else {
		f2.Next = f2.Next.Next
	}

	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	return isSymmetricNode(root.Left, root.Right)
}

func isSymmetricNode(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricNode(left.Left, right.Right) &&
		isSymmetricNode(left.Right, right.Left)
}
