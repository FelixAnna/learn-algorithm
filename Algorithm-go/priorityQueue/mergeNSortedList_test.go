package priorityqueue

import "testing"

func TestMergeKLists_Sorted(t *testing.T) {
	lists := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 3}},
		{Val: 2},
	}

	results := MergeKLists(lists)

	if results == nil || results.Val != 1 ||
		results.Next == nil || results.Next.Val != 2 ||
		results.Next.Next == nil || results.Next.Next.Val != 3 {
		t.Fail()
	}
}

func TestMergeKLists_Empty(t *testing.T) {
	lists := []*ListNode{}

	results := MergeKLists(lists)

	if results != nil {
		t.Fail()
	}
}

func TestMergeKLists_PQ_Sorted(t *testing.T) {
	lists := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 3}},
		{Val: 2},
	}

	results := MergeKLists_PriorityQUeue(lists)

	if results == nil || results.Val != 1 ||
		results.Next == nil || results.Next.Val != 2 ||
		results.Next.Next == nil || results.Next.Next.Val != 3 {
		t.Fail()
	}
}

func TestMergeKLists_PQ_Empty(t *testing.T) {
	lists := []*ListNode{}

	results := MergeKLists_PriorityQUeue(lists)

	if results != nil {
		t.Fail()
	}
}
