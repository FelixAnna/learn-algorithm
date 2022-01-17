package simple

import (
	"strings"
)

/* twoSum: https://leetcode.com/problems/two-sum/
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

solution:
take advantage of existing data structure: map have O(1) access time：
1. loop for one round to keep position of every value
2. also check for the diff part is existing in map or not,
	a. if exists, then we get one answer, just return the index of diff + current
*/
func TwoSum(nums []int, target int) []int {
	var nmap map[int]int = map[int]int{}
	result := make([]int, 2)

	for index, num := range nums {
		diff := target - num
		if _, ok := nmap[diff]; ok {
			result[0] = nmap[diff]
			result[1] = index
			return result
		}

		nmap[num] = index
	}

	return result
}

/*WordPattern
give pattern: abba, words: "dog cat cat dog" is one match
solution:
1. split string s, make sure it have same length as pattern,
2. define 2 maps, for keep: pattern to value , and value to pattern
2. loop rune (char) in pattern
	if rune in map, the value should also in map
	if they both in map, then the value of each mapped item should also as expected
		map[rune] == value && map[value] == rune
	if they both not in map, just add them to map

* make sure diff partten match diff word （so we use 2 maps here)
*/
func WordPattern(pattern string, s string) bool {
	results := strings.Split(s, " ")
	if len(results) != len(pattern) {
		return false
	}

	var maps map[rune]string = map[rune]string{}
	var mapsr map[string]rune = map[string]rune{}
	for i, char := range pattern {
		target := results[i]
		vs, oks := maps[char]
		pr, okr := mapsr[target]
		if oks != okr {
			return false
		}

		if oks {
			if vs != target || char != pr {
				return false
			}

			continue
		} else {
			maps[char] = target
			mapsr[target] = char
		}
	}

	return true
}
