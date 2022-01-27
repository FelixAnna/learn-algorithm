package simple

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/felixanna/algorithm-go/sort"
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
	1. create mappings for one and another
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

/* ZigZag
https://leetcode.com/problems/zigzag-conversion/
*/
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

/*
Reverse integer
abc to cba, out of range returns 0
*/
func Reverse(x int) int {
	result := 0

	sign := true
	abx := x
	if x < 0 {
		abx = -1 * x
		sign = false
	}

	for abx > 0 {
		inc := abx % 10

		if sign && (math.MaxInt32-inc)/10 < result {
			return 0
		} else if !sign && (math.MinInt32+inc)/-10 < result {
			return 0
		} else {
			result = result*10 + inc
			abx = abx / 10
		}
	}

	if !sign {
		result = -1 * result
	}

	return result
}

func MyAtoi(s string) int {
	result := 0
	i := 0
	for s[i] == 32 {
		s = s[i+1:] //remove space
	}

	sign := true
	if s[i] == 43 {
		s = s[i+1:] //positive
	}

	if s[i] == 45 {
		sign = false
		s = s[i+1:] //negative
	}

	fmt.Println(s)
	//32 space,  43+, 45-, 48~57
	for i < len(s) {
		if s[i] >= 48 && s[i] <= 57 {
			inc := int(s[i] - 48)

			if sign && (math.MaxInt32-inc)/10 < result {
				return math.MaxInt32
			} else if !sign && (math.MinInt32+inc)/-10 < result {
				return math.MinInt32
			} else {
				result = result*10 + inc
			}
			i++
		} else {
			return -1
		}
	}

	if !sign {
		result *= -1
	}

	fmt.Println(result)
	return result
}

/*maxArea
give height like [1,8,6,2,5,4,8,3,7], found the max area in the lines from the height and index ()
solution:
the point of the solution in O(n) time is the area = width * mininum height,
so for given two line, if one of them is shorter than another,
we only move the shorter one so it is possible to get a longer line next time
we hope the width is as large as possible, so loop from i=0, j=len-1
*/
func MaxArea(height []int) int {
	maxA := 0
	length := len(height)

	i, j := 0, length-1
	for i < j {
		distance := j - i
		h := height[i]
		if height[j] < h {
			h = height[j]
			j--
		} else {
			i++
		}

		area := distance * h
		if area > maxA {
			maxA = area
		}
	}

	return maxA
}

/* ThreeSum
find 3 different numbers (index diff), make sure they sum result is 0
solution:
1. sort first:
2. loop and found any element = 0 - current elements
3. loop for j, k := i+1, len-1, found 2 sum = elements
	a. to remove duplicate, we need skip when arr[j] = arr[j+1] or arr[k] == arr[k-1]
	b. if 2sum > element, then k--, to decrease
	c. if 2sum < element, then j++, to increase
	c. if they same, then we found the element pair, add to list, continue with j++, k--
*/
func ThreeSum(nums []int) [][]int {
	sortedNums := sort.BucketSort(nums)

	length := len(sortedNums)
	results := make([][]int, 0)

	for i, vx := range sortedNums {
		vyz := 0 - vx
		j, k := i+1, length-1
		for j < k {
			sumjk := sortedNums[j] + sortedNums[k]
			switch {
			case vyz == sumjk:
				{
					results = append(results, []int{i, j, k})
					for j < k && sortedNums[j] == sortedNums[j+1] {
						j++
					}

					for j < k && sortedNums[k] == sortedNums[k-1] {
						k--
					}

					j++
					k--
				}
			case vyz > sumjk:
				k--
			case vyz < sumjk:
				j++
			}
		}
	}

	return results
}

/* IsPalindrome
the ideal is we can reverse the number anc compare,
but a more effcient way is to just reverse half of the number,
and only compare the 2 half should be same or diff in 10 times
*/
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0
	remain := x
	for remain > reverse {
		reverse = reverse*10 + remain%10
		remain = remain / 10
	}

	fmt.Println(reverse, remain)
	return (reverse == remain) || (remain == reverse/10)
}

func ValidMountainArray(arr []int) bool {
	if len(arr) <= 2 || arr[0] >= arr[1] {
		return false
	}

	turn := 0
loop:
	for i := 1; i < len(arr); i++ {
		switch {
		case arr[i] == arr[i-1]:
			return false

		case arr[i] > arr[i-1] && turn == 0: //increase
			continue loop
		case arr[i] < arr[i-1] && turn == 0: //start decrease
			turn += 1
			continue loop
		case arr[i] < arr[i-1] && turn == 1: //decrease
			continue loop
		default:
			return false
		}
	}

	return turn == 1
}

func Spreadsheet() {
	var cmd string
	var sph [][]int
	scanner := bufio.NewScanner(os.Stdin)

	help := func() {
		fmt.Printf("Help Info: %v", cmd)
	}

	fmt.Println("Please input a command:")
	var err error
	for {
		scanner.Scan()
		cmd = scanner.Text()

		c := cmd[0]
		switch c {
		case byte('C'):
			fmt.Println("create spreadsheet", cmd[1:])
			sph, err = createSpread(cmd[1:])
			fmt.Println("created spreadsheet")
		case byte('N'):
			fmt.Println("Insert spreadsheet")
			err = insertSpread(cmd[1:], sph)
		case byte('S'):
			fmt.Println("Sum spreadsheet")
			err = sumSpread(cmd[1:], sph)
		case byte('Q'):
			fmt.Println(sph)
			return
		default:
			help()
		}

		if err != nil {
			help()
		}
	}
}

func createSpread(s string) ([][]int, error) {
	instructions := strings.Split(strings.TrimSpace(s), " ")
	if len(instructions) != 2 {
		fmt.Println(instructions)
		return nil, errors.New("invalid arguments")
	}

	row, errRow := strconv.ParseInt(instructions[0], 10, 32)
	col, errColumn := strconv.ParseInt(instructions[1], 10, 32)
	if errRow != nil || errColumn != nil {
		return nil, errors.New("invalid arguments")
	}

	sph := make([][]int, row)
	for i := 0; i < int(row); i++ {
		sph[i] = make([]int, col)
	}

	return sph, nil
}

func insertSpread(s string, sph [][]int) error {
	instructions := strings.Split(strings.TrimSpace(s), " ")
	if len(instructions) != 3 {
		return errors.New("invalid arguments")
	}

	row, errRow := strconv.ParseInt(instructions[0], 10, 32)
	col, errColumn := strconv.ParseInt(instructions[1], 10, 32)
	val, errVal := strconv.ParseInt(instructions[2], 10, 32)
	if errRow != nil || errColumn != nil || errVal != nil {
		return errors.New("invalid arguments")
	}

	if int(row) >= len(sph) || int(col) >= len(sph[0]) {
		return errors.New("invalid arguments")
	}

	sph[row][col] = int(val)

	return nil
}

func sumSpread(s string, sph [][]int) error {
	instructions := strings.Split(strings.TrimSpace(s), " ")
	if len(instructions)%2 != 0 || len(instructions) < 6 {
		return errors.New("invalid arguments")
	}

	sum := 0
	for i := 0; i < len(instructions)-2; i += 2 {
		row, errRow := strconv.ParseInt(instructions[i], 10, 32)
		col, errColumn := strconv.ParseInt(instructions[i+1], 10, 32)
		if errRow != nil || errColumn != nil {
			return errors.New("invalid arguments")
		}

		if int(row) >= len(sph) || int(col) >= len(sph[0]) {
			return errors.New("invalid arguments")
		}

		val := sph[row][col]
		sum += val
	}

	row, errRow := strconv.ParseInt(instructions[len(instructions)-2], 10, 32)
	col, errColumn := strconv.ParseInt(instructions[len(instructions)-1], 10, 32)
	if errRow != nil || errColumn != nil {
		return errors.New("invalid arguments")
	}

	if int(row) >= len(sph) || int(col) >= len(sph[0]) {
		return errors.New("invalid arguments")
	}

	sph[row][col] = sum

	return nil
}
