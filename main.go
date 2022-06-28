package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	quizFile, qTime, pTime := parseFlags()

	content, err := os.ReadFile(quizFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(content)))

	problems, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
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
func parseFlags() (string, int, int) {
	var quizFile = flag.String("file", "problems.csv", "file in the format of 'question,answer'")
	var qTime = flag.Int("qtime", 0, "the time limit for the quiz in second")
	var pTime = flag.Int("ptime", 0, "the time limit for each problem in second")
	flag.Parse()

	return *quizFile, *qTime, *pTime
}
