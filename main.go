package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type QuizFile interface {
	LoadProblem() ([]problem, error)
}

type CsvFile struct {
	r *csv.Reader
}

func (cf *CsvFile) LoadProblem() ([]problem, error) {

	problems, err := cf.r.ReadAll()
	if err != nil {
		return nil, err
	}

	var p []problem
	for _, pr := range problems {
		p = append(
			p,
			problem{
				q: strings.ToLower(strings.TrimSpace(pr[0])),
				a: strings.ToLower(strings.TrimSpace(pr[1])),
			})
	}

	return p, nil
}

func main() {
	quizFile, qTime, pTime, randomizeP := parseFlags()

	content, err := os.ReadFile(quizFile)
	if err != nil {
		log.Fatal(err)
	}

	file := new(CsvFile)
	ext := strings.Split(quizFile, ".")
	switch ext[len(ext)-1] {
	case "csv":
		file.r = csv.NewReader(strings.NewReader(string(content)))
	}

	q := NewQuiz(file, randomizeP)
	q.MakeQPTimer(qTime)
problemLoop:
	for n, p := range q.problems {

		q.MakeQPTimer(pTime)

		fmt.Printf("Problem #%d: %s = ", n+1, p.q)
		go func() {
			var in string
			fmt.Scanf("%s\n", &in)
			q.answerCh <- strings.ToLower(in)
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
