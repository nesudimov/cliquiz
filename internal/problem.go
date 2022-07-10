package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type Problem struct {
	Q string
	A []string
	V []string
	S int
}

func (p *Problem) PrintProblem() {
	fmt.Printf("Question: %s\n", p.Q)
	if len(p.A) > 1 {
		for i, a := range p.V {
			fmt.Printf("%d. %s\n", i+1, a)
		}
	}
	fmt.Printf("Answer: ")
}

func (p *Problem) SolveProblem(input string) (int, error) {
	answers := strings.Split(input, ",")
	score := 0
	var err error
	var answer string
	for _, a := range answers {

		if len(p.V) > 0 {
			v, err := strconv.Atoi(strings.TrimSpace(a))
			if err != nil {
				return score, err
			}
			answer = p.V[v-1]
		} else {
			answer = a
		}

		if Contains(p.A, answer) {
			score += 1
		}
	}
	return score, err
}

func Contains[any comparable](slice []any, el any) bool {
	for _, i := range slice {
		if el == i {
			return true
		}
	}
	return false
}
