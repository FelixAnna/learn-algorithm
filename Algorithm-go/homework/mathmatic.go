/* homework
Generate plus / minus /multiply / divide expressions and answers pairs
*/
package homework

import (
	"fmt"
	"math/rand"
	"time"
)

type MathCriteria struct {
	Count      int
	InputRange []int

	IncludeMultiplyDivide bool
}

type Question struct {
	Id         int
	Expression string
	Answer     int
}

var r1 *rand.Rand

func init() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
}

func CheckAnswers(questions []Question, answers []int) (int, []string) {
	correct := 0
	failed := make([]string, 0)
	for i, val := range questions {
		if answers[i] == val.Answer {
			correct++
		} else {
			failed = append(failed, fmt.Sprintf("%v. %v %v， but your answer is: %v \r\n", val.Id+1, val.Expression, val.Answer, answers[i]))
		}
	}

	return correct * 100 / len(questions), failed
}

func CreateQuestions(c *MathCriteria) (questions []Question) {
	if c.Count <= 0 || len(c.InputRange) != 2 || c.InputRange[0] < 1 || c.InputRange[1] < c.InputRange[0] {
		fmt.Println("Invalid arguments.")
		return
	}

	i, loop := 0, 0
	for i < c.Count {
		loop++
		if loop > 1000 {
			fmt.Println("Too many duplicates with current range and count", c.InputRange, c.Count)
			break
		}

		num1 := getRandomInt(c.InputRange[0], c.InputRange[1])
		num2 := getRandomInt(c.InputRange[0], c.InputRange[1])

		var qe *Question
		if !c.IncludeMultiplyDivide {
			qe = createPlusMinusQuestions(i, num1, num2, questions)
		} else {
			qe = createMultiplyDivideQuestions(i, num1, num2, questions)
		}

		if nil != qe {
			questions = append(questions, *qe)
			i++
		}
	}

	return
}

func createPlusMinusQuestions(i, num1, num2 int, questions []Question) *Question {
	plus := func(i, n1, n2 int) *Question {
		qe := Question{Id: i, Expression: fmt.Sprintf("%v + %v =", n1, n2), Answer: n1 + n2}
		if checkExistance(qe.Expression, questions) {
			return nil
		}

		return &qe
	}
	minus := func(i, n1, n2 int) *Question {
		qe := Question{Id: i, Expression: fmt.Sprintf("%v - %v =", n1, n2), Answer: n1 - n2}
		if checkExistance(qe.Expression, questions) {
			return nil
		}

		return &qe
	}

	var created *Question
	switch {
	case num1 < num2:
		created = plus(i, num1, num2)
	case num1 > num2:
		created = minus(i, num1, num2)
	case time.Now().UnixMicro()%2 == 0:
		created = minus(i, num1, num2)
	default:
		created = plus(i, num1, num2)
	}

	return created
}

func createMultiplyDivideQuestions(i, num1, num2 int, questions []Question) *Question {
	multiply := func(i, n1, n2 int) *Question {
		qe := Question{Id: i, Expression: fmt.Sprintf("%v * %v =", n1, n2), Answer: n1 * n2}
		if checkExistance(qe.Expression, questions) {
			return nil
		}

		return &qe
	}
	divide := func(i, n1, n2 int) *Question {
		qe := Question{Id: i, Expression: fmt.Sprintf("%v / %v =", n1, n2), Answer: n1 / n2}
		if checkExistance(qe.Expression, questions) {
			return nil
		}

		return &qe
	}

	var created *Question
	switch {
	case num1 < num2:
		created = multiply(i, num1, num2)
	case num1 > num2 && num1%num2 == 0:
		created = divide(i, num1, num2)
	default:
		created = multiply(i, num1, num2)
	}

	return created
}

func checkExistance(exp string, questions []Question) bool {
	for _, val := range questions {
		if val.Expression == exp {
			return true
		}
	}

	return false
}

func getRandomInt(bottom, up int) int {
	return r1.Intn(up-bottom) + bottom
}
