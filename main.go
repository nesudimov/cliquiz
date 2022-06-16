package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

type quiz struct {
	problems []problem
	score    int
}

func (qz *quiz) MakeQuiz(content string) {
	rows := strings.Split(content, "\n")
	for _, row := range rows {
		qz.problems = append(
			qz.problems,
			problem{
				q: strings.Split(row, ",")[0],
				a: strings.Split(row, ",")[1],
			})
	}
}

func main() {
	var quizFile = flag.String("f", "problems.csv", "file in the format of 'question,answer'")
	flag.Parse()

	content, err := os.ReadFile(*quizFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Split(string(content), "\n")[2])
}
