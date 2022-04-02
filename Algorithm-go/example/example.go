package example

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/felixanna/algorithm-go/dynamic"
	"github.com/felixanna/algorithm-go/fib"
	"github.com/felixanna/algorithm-go/homework"
	"github.com/felixanna/algorithm-go/lib"
	"github.com/felixanna/algorithm-go/search"
	"github.com/felixanna/algorithm-go/sort"
)

func MathCheck(count int, min, max int, md bool) {
	scanner := bufio.NewScanner(os.Stdin)

	criteria := &homework.MathCriteria{
		Count:                 count,
		InputRange:            []int{min, max},
		IncludeMultiplyDivide: md,
		ApplicationQuestions:  true,
	}

	questions := homework.CreateQuestions(criteria)

	fmt.Printf("Please input yes[y] if you want to print questions only:")
	scanner.Scan()
	next := strings.ToLower(scanner.Text())
	if next == "y" || next == "yes" {
		ql, qla := "", ""
		for _, qu := range questions {
			ql += fmt.Sprintf("#%v. %v\r\n", qu.Id+1, qu.Expression)
			qla += fmt.Sprintf("#%v. %v %v\r\n", qu.Id+1, qu.Expression, qu.Answer)
		}

		fmt.Println(ql)
		fmt.Println(qla)
		return
	} else {

		fmt.Printf("Please answer below %v questions:\r\n", criteria.Count)
		answers := make([]int, len(questions))

		i := 0
		for i < len(questions) {
			fmt.Print(questions[i].Expression)
			scanner.Scan()
			cmd := scanner.Text()

			ans, errRow := strconv.ParseInt(cmd, 10, 32)
			if errRow != nil {
				continue
			}

			answers[i] = int(ans)
			i++
		}

		score, failures := homework.CheckAnswers(questions, answers)

		fmt.Printf("You score: %v, ", score)
		if score == 100 {
			fmt.Printf("Congratulations, you got an A+ !")
		} else {
			fmt.Println(failures)
		}
	}
}

/* sort
ways to sort:
	1. quick sort - O(nlogn) , sort in original array, less memory cost
	2. merge sort - O(nlogn) , need more memory
	3. counting sort - O(n+k) , for elements all in [m, n],  m and n are close numbers
	4. radix sort - O(n*k) , when all elements are positive, k normally not very large (like 32 / 4)
	5. bucket sort - O(n*k), split element in n bucket, then use counting sort to sort each bucket, finally append all of them
*/
func TestSort(n int, action int) {
	fmt.Println("TestSort:", n, action)
	arr := lib.Seed(n)

	start := time.Now()
	switch action {
	case 1:
		{
			sort.QuickSort(arr)
		}
	case 2:
		{
			arr = sort.MergeSort(arr)
		}
	case 3:
		{
			bottom, up := 10000, 35000
			arr = lib.SeedWithinRange(n, bottom, up)
			//fmt.Println(arr)
			arr = sort.CountingSort(arr, bottom, up)
			//fmt.Println(arr)
		}
	case 4:
		{
			arr = sort.RadixSort(arr)
		}
	case 5:
		{
			arr = sort.BucketSort(arr)
		}
	default:
		fmt.Println("not support")
	}

	cost := time.Now().UnixMilli() - start.UnixMilli()
	fmt.Println("cost: ", cost)

	fmt.Println("sorted, checking")
	if err := lib.CheckSorted(arr, n); err != nil {
		fmt.Println(err)
	}

	fmt.Println("done")
}

/* sort
ways to search:
	1. bin search in a sorted array
	2. find xth max / min element in an unsorted array
*/
func TestSearch(n int, action int) {
	fmt.Println("TestSearch:", n, action)
	arr := lib.Seed(n)
	var element int
	var index int = -1

	start := time.Now()
	switch action {
	case 1:
		{
			element = arr[n/2]
			arr = sort.RadixSort(arr)
			index = search.BinarySearch(arr, element)

		}
	case 2:
		{
			index, element = search.FindNthMin(arr, 5)
		}
	case 3:
		{
			index, element = search.FindMin(arr)
		}
	case 4:
		{
			index, element = search.FindMax(arr)
		}
	default:
		fmt.Println("not support")
	}

	cost := time.Now().UnixMilli() - start.UnixMilli()
	fmt.Println("cost: ", cost)

	fmt.Println("result: ", arr, index, element)

	fmt.Println("search, checking")
	if index >= 0 && arr[index] != element {
		fmt.Println("Not same incorrect:", index, arr[index], element)
		return
	}

	fmt.Println("done")
}

/* fibonacci:
2 ways to get fib:
	1 - bottom up calculate fibonacci, very efficient
	2 - recursive calculate fibonacci, normally not bigger than 50
*/
func TestFib(n int, action int) {
	fmt.Println("TestFib:", n, action)
	start := time.Now()
	switch action {
	case 1:
		{
			results := fib.FibBottomUp(n)
			fmt.Println(results[n-1] + results[n-2])
		}
	case 2:
		{
			done := make(chan bool)
			go func() {
				a := fib.FibRecursive(n) // 0 ~ n : n+1 level
				fmt.Println(a)
				done <- true
			}()

			<-done

		}
	default:
		fmt.Println("not support")
	}

	cost := time.Now().UnixMilli() - start.UnixMilli()
	fmt.Println("cost: ", cost)

	fmt.Println("done")
}

func TestDynamicProgramming(n int, action int) {
	strategies := map[int]float64{
		1: 1,
		2: 5,
		3: 8,
		4: 9,
		5: 10,
		6: 17,
		7: 17,
		8: 20,
	}

	st, profit := dynamic.CutRod(n, strategies)

	fmt.Println(n, st, profit)
}
