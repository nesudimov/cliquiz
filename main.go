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
	var quizFile = flag.String("file", "problems.csv", "file in the format of 'question,answer'")
	var qTime = flag.Int("qtime", 0, "the time limit for the quiz in second")
	var pTime = flag.Int("ptime", 0, "the time limit for each problem in second")
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

	q := NewQuiz(problems)
	q.runQuiz(*qTime, *pTime)
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
func NewQuiz(problems [][]string) *quiz {
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
	qz.timer = &time.Timer{C: make(chan time.Time)}

	return qz
}

func (qz *quiz) runQuiz(qt int, pt int) {
	qz.MakeQPTimer(qt)
	answerCh := make(chan string)
problemLoop:
	for n, p := range qz.problems {

		qz.MakeQPTimer(pt)

		fmt.Printf("Problem #%d: %s = ", n+1, p.q)
		go func() {
			var in string
			fmt.Scanf("%s\n", &in)
			answerCh <- in
		}()

		select {
		case <-qz.timer.C:
			fmt.Println("#### time is over ####")
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

// MakeQPTimer receives the duration of the timer in seconds (t).
// If t is non-zero, creates a new timer with duration t
func (qz *quiz) MakeQPTimer(t int) {
	if t != 0 {
		qz.timer = time.NewTimer(time.Duration(t) * time.Second)
	}
}
