package simple

import (
	"fmt"
	"strings"
)

/* LongestPalindrome
Find Palindrome (aba, abba) in a given string
	1. create mappings for s and reverse s
	2. sum in mappings: if mp[x][y] = 1, then map[x][y]=map[x-1][y-1]+1, else: mp[x][y]= max(mp[x-1][y], mp[x][y-1])
	3. from mp[m][n] to map[0][0]: if mp[x-1][y] == mp[x][y], x=x-1, else if mp[x][y-1] == mp[x][y],y=y-1, else found the position: print till map[i][j]=1
*/
func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	longestPalindrom := ""
	for i := 0; i < len(s); i++ {
		pal1 := extend(s, i, i)
		pal2 := extend(s, i, i+1)

		if len(longestPalindrom) < len(pal1) {
			longestPalindrom = pal1
		}

		if len(longestPalindrom) < len(pal2) {
			longestPalindrom = pal2
		}
	}

	fmt.Println(longestPalindrom)

	return longestPalindrom
}

func extend(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return s[i+1 : j]
}

/* LongestCommonString
Find longest common substring in 2 given string
	1. create mappings for s and reverse s
	2. sum in mappings: if mp[x][y] = 1, then map[x][y]=map[x-1][y-1]+1, else: mp[x][y]= max(mp[x-1][y], mp[x][y-1])
	3. from mp[m][n] to map[0][0]: if mp[x-1][y] == mp[x][y], x=x-1, else if mp[x][y-1] == mp[x][y],y=y-1, else found the position: print till map[i][j]=1
*/
func LongestCommonString(one, another string) string {
	lengthOne := len(one)
	lengthTwo := len(another)
	if lengthOne <= 0 || lengthTwo <= 0 {
		return ""
	}
	mp := make([][]int, lengthTwo+1)
	for j := 0; j < lengthTwo; j++ {
		mp[j] = make([]int, lengthOne+1)
		for i := 0; i < lengthOne; i++ {
			if one[i] == another[j] {
				mp[j][i] = 1
			}
		}
	}

	mp[lengthTwo] = make([]int, lengthOne+1)
	/*for i := 0; i < len(mp); i++ {
		fmt.Println(mp[i])
	}*/

	commonLen := 0
	maxI, maxJ := 0, 0
	for i := 1; i < lengthTwo+1; i++ {
		for j := 1; j < lengthOne+1; j++ {
			if mp[i][j] == 1 {
				mp[i][j] = mp[i][j] + mp[i-1][j-1]
				if mp[i][j] > commonLen {
					maxI, maxJ = i, j
					commonLen = mp[i][j]
				}
			}
		}
	}
	/*for i := 0; i < len(mp); i++ {
		fmt.Println(mp[i])
	}*/

	result := one[maxJ-commonLen+1 : maxJ+1]

	fmt.Println(result, maxI, maxJ, commonLen)
	return result
}

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	zigzag := make(map[int][]int, len(s))

	row, col := 0, 0
	for i := 0; i < len(s); i++ {
		row, col = getPosition(s, i, numRows)
		zigzag[i] = []int{row, col}
	}

	resList := make([]string, numRows)
	for i := 0; i < len(s); i++ {
		row = zigzag[i][0]
		resList[row] += s[i : i+1]
	}

	fmt.Println(resList)
	return strings.Join(resList, "")
}

func getPosition(s string, i, numRows int) (int, int) {
	unitCapacity := numRows + numRows - 2
	unitNumbers := i / unitCapacity
	unitIndex := i % unitCapacity
	unitCols := 1 + numRows/2

	row, col := 0, 0
	if unitIndex < numRows {
		row = unitIndex
		col = 0
	} else {
		row = numRows - (unitIndex + 1 - numRows) - 1
		col = unitIndex - numRows + 1
	}

	col += unitNumbers * unitCols
	return row, col
}
