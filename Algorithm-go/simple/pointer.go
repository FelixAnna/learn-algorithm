package simple

/* LengthOfLongestSubstring (without repeat)
find the lenght of the longest substring without repeating char in a given string
solution:
1. define 2 pointer index (start, end) for keep boundary of latest no duplicate string area;
2. define a map for keep every char's last index;
3. initial the 2 index to 0;
3. loop the char in the string
	a. if the char is not present in map (between start to end),
		then add the char to map;
		get the length of no duplicate area, if greater, then save to max
	b. if the char is present in map  (between start to end),
		then set the start = last present index + 1 (new no deplicate area);
		update the index of the char in map to current position;
		continue;

*/
func LengthOfLongestSubstring(s string) int {
	lgstLen := 0
	charIndexMap := map[byte]int{}

	for i, j := 0, 0; j < len(s); j++ {
		char := s[j]
		if index, ok := charIndexMap[char]; ok && index >= i { //duplicate in current range [i, j)
			i = index + 1
			charIndexMap[char] = j
			continue
		} else {
			charIndexMap[char] = j

			curLen := j - i + 1
			if curLen > lgstLen {
				lgstLen = curLen
			}
		}
	}

	return lgstLen
}

/*
set red index to the first elements,
set blue index to the last elements,

then for i=0; i<blue; i++ {
	if nums[i] == red {
		swap(i, red)
		red++
		i++
	} else if nums[i] == blue {
		swap(i, blue)
		blue--
		continue
	} else {
		i++
	}
}

return nums
*/
func SortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	red, blue := 0, len(nums)-1

	i := red
	for i <= blue {
		if nums[i] == 2 {
			nums[i], nums[blue] = nums[blue], nums[i]
			blue--
			continue
		}
		if nums[i] == 0 {
			nums[i], nums[red] = nums[red], nums[i]
			red++
		}

		i++
	}

}
