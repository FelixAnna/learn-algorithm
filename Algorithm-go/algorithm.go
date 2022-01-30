package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/felixanna/algorithm-go/homework"
)

func main() {
	mathCheck(5, 1, 10, true)
}

func mathCheck(count int, min, max int, md bool) {
	scanner := bufio.NewScanner(os.Stdin)

	criteria := &homework.MathCriteria{
		Count:                 count,
		InputRange:            []int{min, max},
		IncludeMultiplyDivide: md,
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
