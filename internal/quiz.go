package internal

import (
	"fmt"
	"log"
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
	qz.LoadTotalScore()
	qz.Timer = &time.Timer{C: make(chan time.Time)}
	qz.AnswerCh = make(chan string)

	fmt.Println("Debug:", qz.Problems)
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
		score, err := qz.Problems[pNum].SolveProblem(answer)
		if err != nil {
			log.Fatal(err)
		}
		qz.PlayerScore += score
	}
	return true
}

// LoadTotalScore adds the score for each problem.
// Writes the results of the addition to qz.TotalScore.
func (qz *quiz) LoadTotalScore() {
	for _, p := range qz.Problems {
		qz.TotalScore += p.S
	}
}

// PrintScore printed quiz score
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
