package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	quizFile, qTime, pTime, randomizeP := parseFlags()

	content, err := os.ReadFile(quizFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(content)))

	problems, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if randomizeP {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	q := NewQuiz(problems)

	q.MakeQPTimer(qTime)
problemLoop:
	for n, p := range q.problems {

		q.MakeQPTimer(pTime)

		fmt.Printf("Problem #%d: %s = ", n+1, p.q)
		go func() {
			var in string
			fmt.Scanf("%s\n", &in)
			q.answerCh <- in
		}()

		if !q.QuizHandler(n) {
			break problemLoop
		}
	}
	q.PrintScore()

}

// parseFlags parses the cli flags from os.Args[1:] and return their values
func parseFlags() (string, int, int, bool) {
	var quizFile = flag.String("file", "problems.csv", "file in the format of 'question,answer'")
	var qTime = flag.Int("qtime", 0, "the time limit for the quiz in second")
	var pTime = flag.Int("ptime", 0, "the time limit for each problem in second")
	var randomizeP = flag.Bool("randomize", false, "enables random display of tasks")
	flag.Parse()

	return *quizFile, *qTime, *pTime, *randomizeP
}
