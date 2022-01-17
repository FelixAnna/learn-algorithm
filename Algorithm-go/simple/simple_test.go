package simple

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 17}
	target := 9
	results := TwoSum(nums, target)

	if len(results) != 2 || nums[results[0]]+nums[results[1]] != 9 {
		t.Fail()
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{2, nil}}
	l2 := &ListNode{Val: 9, Next: &ListNode{7, nil}}
	results := AddTwoNumbers(l1, l2)

	if results.Val != 0 || results.Next.Val != 0 || results.Next.Next.Val != 1 {
		t.Fail()
	}
}

func TestWordPattern_OK(t *testing.T) {
	pattern := "abba"
	value := "cat dog dog cat"
	result := WordPattern(pattern, value)
	if !result {
		t.Fail()
	}
}

func TestWordPattern_Error(t *testing.T) {
	pattern := "abba"
	value := "cat dog dog cat1"
	result := WordPattern(pattern, value)
	if result {
		t.Fail()
	}
}

func TestLengthOfLongestSubstring_Normal(t *testing.T) {
	str := "pwwkew"
	expectLen := 3
	result := LengthOfLongestSubstring(str)

	if expectLen != result {
		t.Errorf("Expected: %v, actual: %v", expectLen, result)
	}
}

func TestLengthOfLongestSubstring_Space(t *testing.T) {
	str := " "
	expectLen := 1
	result := LengthOfLongestSubstring(str)

	if expectLen != result {
		t.Errorf("Expected: %v, actual: %v", expectLen, result)
	}
}
