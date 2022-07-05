package internal

import (
	"fmt"
	"math/rand"
	"time"
)

type quiz struct {
	Problems    []Problem
	PlayerScore int
	TotalScore  int
	Timer       *time.Timer
	AnswerCh    chan string
}

// Quiz constructor
func NewQuiz(qf QuizFile, randomizeP bool) *quiz {
	qz := new(quiz)

	qz.Problems, _ = qf.LoadProblem()
	if randomizeP {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(qz.Problems), func(i, j int) {
			qz.Problems[i], qz.Problems[j] = qz.Problems[j], qz.Problems[i]
		})
	}

	qz.PlayerScore = 0
	qz.TotalScore = len(qz.Problems)
	qz.Timer = &time.Timer{C: make(chan time.Time)}
	qz.AnswerCh = make(chan string)

	return qz
}

// QuizHandler handles timer and answer channels.
// If data arrives in the timer channel, displays "time is over" and returns false.
// If data arrives in the answer channel, increment score if answer is equal to problem.a and return true.
func (qz *quiz) QuizHandler(pNum int) bool {
	select {
	case <-qz.Timer.C:
		fmt.Println("#### time is over ####")
		return false
	case answer := <-qz.AnswerCh:
		if answer == qz.Problems[pNum].A {
			qz.PlayerScore++
		}
	}
	return true
}

func (qz *quiz) PrintScore() {
	fmt.Printf("You scored %d out of %d.\n", qz.PlayerScore, qz.TotalScore)
}

// MakeQPTimer receives the duration of the timer in seconds (t).
// If t is non-zero, creates a new timer with duration t
func (qz *quiz) MakeQPTimer(t int) {
	if t != 0 {
		qz.Timer = time.NewTimer(time.Duration(t) * time.Second)
	}
}
