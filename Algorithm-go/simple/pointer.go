package simple

/* LengthOfLongestSubstring
find the lenght of the longest substring in a given string
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
	mp := map[byte]int{}

	for i, j := 0, 0; j < len(s); j++ {
		value := s[j]
		if index, ok := mp[value]; ok && index >= i { //duplicate in current range [i, j)
			i = index + 1
			mp[value] = j
			continue
		} else {
			mp[value] = j

			curLen := j - i + 1
			if curLen > lgstLen {
				lgstLen = curLen
			}
		}
	}

	return lgstLen
}
