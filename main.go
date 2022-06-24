package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var quizFile = flag.String("f", "problems.csv", "file in the format of 'question,answer'")
	var timeLimit = flag.Int("t", 0, "the time limit for the quiz in second")
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

	q := MakeQuiz(problems, *timeLimit)
	q.runQuiz()
}

type problem struct {
	q string
	a string
}

type quiz struct {
	problems    []problem
	playerScore int
	totalScore  int
	timer       *time.Timer
}

// Quiz constructor
func MakeQuiz(problems [][]string, timeLimit int) *quiz {
	qz := new(quiz)
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

	if timeLimit != 0 {
		qz.timer = time.NewTimer(time.Duration(timeLimit) * time.Second)
	} else {
		c := make(chan time.Time)
		qz.timer = &time.Timer{C: c}
	}

	return qz
}

func (qz *quiz) runQuiz() {
	answerCh := make(chan string)
problemLoop:
	for n, p := range qz.problems {
		fmt.Printf("Problem #%d: %s = ", n+1, p.q)
		go func() {
			var in string
			fmt.Scanf("%s\n", &in)
			answerCh <- in
		}()

		select {
		case <-qz.timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				qz.playerScore++
			}
		}
	}
	qz.printScore()
}

func (qz *quiz) printScore() {
	fmt.Printf("You scored %d out of %d.\n", qz.playerScore, qz.totalScore)
}
