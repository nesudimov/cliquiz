package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	internal "github.com/nesudimov/cliquiz/internal"
)

func main() {
	quizFile, qTime, pTime, randomizeP := parseFlags()

	content, err := os.ReadFile(quizFile)
	if err != nil {
		log.Fatal(err)
	}

	file := new(internal.CsvFile)
	ext := strings.Split(quizFile, ".")
	switch ext[len(ext)-1] {
	case "csv":
		file.R = csv.NewReader(strings.NewReader(string(content)))
	}

	q := internal.NewQuiz(file, randomizeP)
	q.MakeQPTimer(qTime)
problemLoop:
	for n, p := range q.Problems {

		q.MakeQPTimer(pTime)

		fmt.Printf("Problem #%d: %s = ", n+1, p.Q)
		go func() {
			var in string
			fmt.Scanf("%s\n", &in)
			q.AnswerCh <- strings.ToLower(in)
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
	var pTime = flag.Int("ptime", 0, "the time limit for each internal.Problem in second")
	var randomizeP = flag.Bool("randomize", false, "enables random display of tasks")
	flag.Parse()

	return *quizFile, *qTime, *pTime, *randomizeP
}
