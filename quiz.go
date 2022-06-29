package main

import (
	"fmt"
	"strings"
	"time"
)

type quiz struct {
	problems    []problem
	playerScore int
	totalScore  int
	timer       *time.Timer
	answerCh    chan string
}

// Quiz constructor
func NewQuiz(problems [][]string) *quiz {
	qz := new(quiz)
	for _, pr := range problems {
		qz.problems = append(
			qz.problems,
			problem{
				q: strings.ToLower(strings.TrimSpace(pr[0])),
				a: strings.ToLower(strings.TrimSpace(pr[1])),
			})
	}
	qz.playerScore = 0
	qz.totalScore = len(qz.problems)
	qz.timer = &time.Timer{C: make(chan time.Time)}
	qz.answerCh = make(chan string)

	return qz
}

// QuizHandler handles timer and answer channels.
// If data arrives in the timer channel, displays "time is over" and returns false.
// If data arrives in the answer channel, increment score if answer is equal to problem.a and return true.
func (qz *quiz) QuizHandler(pNum int) bool {
	select {
	case <-qz.timer.C:
		fmt.Println("#### time is over ####")
		return false
	case answer := <-qz.answerCh:
		if answer == qz.problems[pNum].a {
			qz.playerScore++
		}
	}
	return true
}

func (qz *quiz) PrintScore() {
	fmt.Printf("You scored %d out of %d.\n", qz.playerScore, qz.totalScore)
}

// MakeQPTimer receives the duration of the timer in seconds (t).
// If t is non-zero, creates a new timer with duration t
func (qz *quiz) MakeQPTimer(t int) {
	if t != 0 {
		qz.timer = time.NewTimer(time.Duration(t) * time.Second)
	}
}
