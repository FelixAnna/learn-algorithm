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

func TestReverse_Positive(t *testing.T) {
	val := 123
	result := Reverse(val)

	if result != 321 {
		t.Fail()
	}
}

func TestReverse_Negative(t *testing.T) {
	val := -123
	result := Reverse(val)

	if result != -321 {
		t.Fail()
	}
}

func TestReverse_Edge(t *testing.T) {
	val := -2147483648
	result := Reverse(val)

	if result != 0 {
		t.Fail()
	}
}

func TestIsPalindrome_Pos(t *testing.T) {
	num := 121
	result := IsPalindrome(num)
	if !result {
		t.Fail()
	}
}

func TestIsPalindrome_NonPalin(t *testing.T) {
	num := 123
	result := IsPalindrome(num)
	if result {
		t.Fail()
	}
}

func TestIsPalindrome_Neg(t *testing.T) {
	num := -121
	result := IsPalindrome(num)
	if result {
		t.Fail()
	}
}

func TestMaxArea_Multiple(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := MaxArea(height)
	if result != 49 {
		t.Fail()
	}
}

func TestMaxArea_One(t *testing.T) {
	height := []int{1, 1}
	result := MaxArea(height)
	if result != 1 {
		t.Fail()
	}
}
func Test3Sum_Ok(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := ThreeSum(nums)

	if len(result) != 2 {
		t.Fail()
	}
}

func TestLetterCombinations_Ok(t *testing.T) {
	digits := "23"
	results := LetterCombinations(digits)

	if len(results) != 9 {
		t.Fail()
	}
}

func TestLetterCombinations_Empty(t *testing.T) {
	digits := ""
	results := LetterCombinations(digits)

	if len(results) != 0 {
		t.Fail()
	}
}

func TestValidMountainArray_OK(t *testing.T) {
	arr := []int{1, 2, 3, 2}
	result := ValidMountainArray(arr)
	if !result {
		t.Fail()
	}
}

func TestValidMountainArray_Neg(t *testing.T) {
	arr := []int{1, 2, 3, 3}
	result := ValidMountainArray(arr)
	if result {
		t.Fail()
	}
}

func TestMergeSortedList_Ok(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	result := MergeTwoLists(l1, l2)
	if result.Val != 1 || result.Next.Val != 1 || result.Next.Next.Val != 2 || result.Next.Next.Next.Val != 3 || result.Next.Next.Next.Next.Val != 4 || result.Next.Next.Next.Next.Next.Val != 4 {
		t.Fail()
	}
}

func TestRemoveNthFromEnd_Middle(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	result := RemoveNthFromEnd(l1, 2)
	if result.Val != 1 || result.Next.Val != 4 || result.Next.Next != nil {
		t.Fail()
	}
}

func TestRemoveNthFromEnd_First(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	result := RemoveNthFromEnd(l1, 3)
	if result.Val != 2 || result.Next.Val != 4 || result.Next.Next != nil {
		t.Fail()
	}
}

func TestRemoveNthFromEnd_Last(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	result := RemoveNthFromEnd(l1, 1)
	if result.Val != 1 || result.Next.Val != 2 || result.Next.Next != nil {
		t.Fail()
	}
}

func TestMyAtoi_Positive(t *testing.T) {
	input := "123"
	result := MyAtoi(input)
	if result != 123 {
		t.Fail()
	}
}

func TestMyAtoi_Negative(t *testing.T) {
	input := "-123"
	result := MyAtoi(input)
	if result != -123 {
		t.Fail()
	}
}

func TestMyAtoi_Zero(t *testing.T) {
	input := "+0"
	result := MyAtoi(input)
	if result != 0 {
		t.Fail()
	}
}

func TestSpreadSheet(t *testing.T) {
	cmds := []string{
		"C 10 20",
		"N 3 4 6",
		"N 4 5 6",
		"S 3 4 4 5 5 6",
	}
	result := Spreadsheet(cmds)
	if len(result) != 10 || len(result[5]) != 20 || result[5][6] != 12 {
		t.Fail()
	}
}

func TestValidParentheses_Ok(t *testing.T) {
	result := ValidParentheses("()[{}]")
	if result != true {
		t.Fail()
	}
}

func TestValidParentheses_Error(t *testing.T) {
	result := ValidParentheses("()[{)]")
	if result == true {
		t.Fail()
	}
}

func TestLongestCommonPrefix_Decrease(t *testing.T) {
	strs := []string{
		"ab", "a",
	}
	result := LongestCommonPrefix(strs)

	if result != "a" {
		t.Fail()
	}
}

func TestLongestCommonPrefix_Common(t *testing.T) {
	strs := []string{
		"farway", "family", "first",
	}
	result := LongestCommonPrefix(strs)

	if result != "f" {
		t.Fail()
	}
}

func TestRomanToInt_IncreaseOnly(t *testing.T) {
	input := "DCXXI"
	result := RomanToInt(input)
	if result != 621 {
		t.Fail()
	}
}

func TestRomanToInt_Decrease(t *testing.T) {
	input := "IV"
	result := RomanToInt(input)
	if result != 4 {
		t.Fail()
	}
}

func TestIsValidSudoku(t *testing.T) {
	input := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	result := IsValidSudoku(input)
	if !result {
		t.Fail()
	}
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := GroupAnagrams(strs)

	if len(res) != 3 {
		t.Fail()
	}
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ret := SpiralOrder(matrix)
	if ret[0] != 1 || ret[8] != 5 || len(ret) != 9 {
		t.Fail()
	}
}

func TestMergeIntervals(t *testing.T) {
	intervals := [][]int{{3, 7}, {6, 8}, {10, 15}}

	ret := MergeIntervals(intervals)
	if len(ret) != 2 {
		t.Fail()
	}
}

func TestPlusOne(t *testing.T) {
	digits := []int{9, 9, 9}
	ret := PlusOne(digits)

	if ret[0] != 1 || len(ret) != 4 {
		t.Fail()
	}
}

func TestMySqrt(t *testing.T) {
	ret := MySqrt(399)
	if ret != 19 {
		t.Fail()
	}
}

func TestMySqrt2(t *testing.T) {
	ret := MySqrt2(399)
	if ret != 19 {
		t.Fail()
	}
}

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	SetZeroes(matrix)
	if matrix[0][1] != 0 || matrix[0][2] != 0 || matrix[1][0] != 0 || matrix[1][3] != 0 {
		t.Fail()
	}
}
