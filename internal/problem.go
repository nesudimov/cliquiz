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
	answer := strings.Split(input, ",")
	score := 0
	var err error
	for _, a := range answer {
		v, err := strconv.Atoi(strings.TrimSpace(a))
		if err != nil {
			return score, err
		}

		if Contains(p.A, p.V[v-1]) {
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
