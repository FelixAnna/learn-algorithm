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

	ApplicationQuestions bool
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
			qe = createPlusMinusQuestions(i, num1, num2, questions, c.ApplicationQuestions)
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

func createPlusMinusQuestions(i, num1, num2 int, questions []Question, application bool) *Question {
	plus := func(i, n1, n2 int) *Question {
		var expression string
		if application {
			expression = createdPlusMinusCase(num1, num2, true)
		} else {
			expression = fmt.Sprintf("%v + %v =", n1, n2)
		}
		qe := Question{Id: i, Expression: expression, Answer: n1 + n2}
		if checkExistance(qe.Expression, questions) {
			return nil
		}

		return &qe
	}
	minus := func(i, n1, n2 int) *Question {
		var expression string
		if application {
			expression = createdPlusMinusCase(num1, num2, false)
		} else {
			expression = fmt.Sprintf("%v - %v =", n1, n2)
		}
		qe := Question{Id: i, Expression: expression, Answer: n1 - n2}
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

func createdPlusMinusCase(i, j int, plus bool) string {
	users := []string{"小明", "小张", "小周", "小吴", "Kitty", "Alice", "Tom", "Eddie", "Anna", "Luca"}
	goods := []string{"Apple", "Orange", "Pear", "巧克力糖果"}
	minus_templates := []string{"%v买了%v个%v,弄丢了%v个,还剩几个?",
		"%v有%v个%v,吃掉了%v个,还有几个?",
		"%v昨天获得了%v个%v的奖励,已经吃掉了%v个,还剩几个?",
		"%v有%v元钱, 他想买一个%v,苹果的价格是%v元,买完之后他还有多少钱?",
		"%v有%v个%v,吃掉了一些,还有%v个, 吃掉了几个?",
	}

	plus_templates := []string{"%v刚刚买了%v个%v,然后又发现冰箱里还有%v个,现在总共有几个?",
		"%v昨天获得了%v个%v的奖励,前台也获得了%v个,一个获得了几个?"}

	user := users[getRandomInt(0, len(users)-1)]
	good := goods[getRandomInt(0, len(goods)-1)]
	var template string
	if plus {
		template = plus_templates[getRandomInt(0, len(plus_templates)-1)]
	} else {
		template = minus_templates[getRandomInt(0, len(minus_templates)-1)]
	}

	result := fmt.Sprintf(template, user, i, good, j)
	return result
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
