package simple

import (
	"testing"
)

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

func TestLongestCommonString_Symmetry(t *testing.T) {
	str1 := "abcbde"
	str2 := "edbcba"
	result := LongestCommonString(str1, str2)
	if result != "bcb" {
		t.Fail()
	}
}

func TestLongestCommonString_UNSymmetry(t *testing.T) {
	str1 := "aacabdkacaa"
	str2 := "kdacaaaa"
	result := LongestCommonString(str1, str2)
	if len(result) != 4 {
		t.Fail()
	}
}

func TestLongestCommonString_One(t *testing.T) {
	str1 := "aacabdkacaa"
	str2 := "1a1"
	result := LongestCommonString(str1, str2)
	if len(result) != 1 {
		t.Fail()
	}
}

func TestLongestCommonString_Empty(t *testing.T) {
	str1 := "aacabdkacaa"
	str2 := ""
	result := LongestCommonString(str1, str2)
	if len(result) != 0 {
		t.Fail()
	}
}

func TestLongestPalindrome_Normal(t *testing.T) {
	val := "sadab1asd"
	result := LongestPalindrome(val)
	if result != "ada" {
		t.Fail()
	}
}

func TestLongestPalindrome_One(t *testing.T) {
	val := "s"
	result := LongestPalindrome(val)
	if result != "s" {
		t.Fail()
	}
}

func TestConvert(t *testing.T) {
	val :=
		"PAYPALISHIRING"
	result := Convert(val, 3)
	if result !=
		"PAHNAPLSIIGYIR" {
		t.Fail()
	}
}

func TestConvert_Two(t *testing.T) {
	val :=
		"ABCD"
	result := Convert(val, 2)
	if result !=
		"ACBD" {
		t.Fail()
	}
}

/*P  A   H   N
AP L S I I G
Y  I   R*/
