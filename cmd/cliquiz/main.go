package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	i "github.com/nesudimov/cliquiz/internal"
	"gopkg.in/yaml.v3"
)

func main() {
	quizFile, qTime, pTime, randomizeP := parseFlags()

	file, err := defineQuizFile(quizFile)
	if err != nil {
		log.Fatal(err)
	}

	q := i.NewQuiz(file, randomizeP)
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

// defineQuizFile gets the path to the quiz file.
// Determines the file extension and returns the required QuizFile
func defineQuizFile(filePath string) (i.QuizFile, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	ext := strings.Split(filePath, ".")

	switch ext[len(ext)-1] {
	case "json":
		return &i.JsonFile{
			D: json.NewDecoder(strings.NewReader(string(content))),
		}, nil
	case "yml":
		return &i.YmlFile{
			D: yaml.NewDecoder(strings.NewReader(string(content))),
		}, nil
	case "csv":
		return &i.CsvFile{
			R: csv.NewReader(strings.NewReader(string(content))),
		}, nil
	default:
		return nil, err
	}
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
