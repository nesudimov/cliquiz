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

func (qz *quiz) MakeQuiz(problems [][]string, timeLimit int) {
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
	qz.timer = time.NewTimer(time.Duration(timeLimit) * time.Second)
}

func (qz *quiz) runQuiz() {
problemLoop:
	for n, p := range qz.problems {
		fmt.Printf("Problem #%d: %s = ", n+1, p.q)
		answerCh := make(chan string)
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

func main() {
	var quizFile = flag.String("f", "problems.csv", "file in the format of 'question,answer'")
	var timeLimit = flag.Int("t", 30, "the time limit for the quiz in second")
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
	q.MakeQuiz(problems, *timeLimit)
	q.runQuiz()
}
