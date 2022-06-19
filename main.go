package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

type quiz struct {
	problems    []problem
	playerScore int
	totalScore  int
}

func (qz *quiz) MakeQuiz(problems [][]string) {
	for _, pr := range problems {
		qz.problems = append(
			qz.problems,
			problem{
				q: strings.TrimSpace(pr[0]),
				a: strings.TrimSpace(pr[1]),
			})
	}
	qz.playerScore = 0
	qz.totalScore = len(qz.problems)
}

func (qz *quiz) runQuiz() {
	var in string
	for n, p := range qz.problems {
		in = ""
		fmt.Printf("Problem #%d: %s = ", n+1, p.q)
		fmt.Scanf("%s\n", &in)
		if in == p.a {
			qz.playerScore++
		}
	}
}

func main() {
	var quizFile = flag.String("f", "problems.csv", "file in the format of 'question,answer'")
	flag.Parse()

	content, err := os.ReadFile(*quizFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(content)))

	problems, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var q quiz
	q.MakeQuiz(problems)
	q.runQuiz()

	fmt.Printf("You scored %d out of %d.\n", q.playerScore, q.totalScore)
}
