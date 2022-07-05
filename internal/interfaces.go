package internal

type QuizFile interface {
	LoadProblem() ([]Problem, error)
}
