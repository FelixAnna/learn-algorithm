package homework

import "testing"

func TestCreateQuestions_PA(t *testing.T) {
	count := 5
	min, max := 1, 10
	md := false
	criteria := MathCriteria{
		Count:                 count,
		InputRange:            []int{min, max},
		IncludeMultiplyDivide: md,
	}

	result := CreateQuestions(&criteria)
	if len(result) != count {
		t.Fail()
	}
}

func TestCreateQuestions_MD(t *testing.T) {
	count := 5
	min, max := 1, 10
	md := true
	criteria := MathCriteria{
		Count:                 count,
		InputRange:            []int{min, max},
		IncludeMultiplyDivide: md,
	}

	result := CreateQuestions(&criteria)
	if len(result) != count {
		t.Fail()
	}
}

func TestCheckAnswers_100(t *testing.T) {
	questions := []Question{
		{Id: 0, Expression: "1+2=", Answer: 3},
		{Id: 0, Expression: "2-1=", Answer: 1},
		{Id: 0, Expression: "1+3=", Answer: 4},
	}

	answers := []int{3, 1, 4}

	score, msg := CheckAnswers(questions, answers)
	if score != 100 || len(msg) != 0 {
		t.Fail()
	}
}

func TestCheckAnswers_667(t *testing.T) {
	questions := []Question{
		{Id: 0, Expression: "1+2=", Answer: 3},
		{Id: 0, Expression: "2-1=", Answer: 1},
		{Id: 0, Expression: "1+3=", Answer: 4},
	}

	answers := []int{3, 1, 5}

	score, msg := CheckAnswers(questions, answers)
	if score == 100 || len(msg) == 0 {
		t.Fail()
	}
}
