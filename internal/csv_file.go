package internal

import (
	"encoding/csv"
	"strings"
)

type CsvFile struct {
	R *csv.Reader
}

func (cf *CsvFile) LoadProblem() ([]Problem, error) {

	problems, err := cf.R.ReadAll()
	if err != nil {
		return nil, err
	}

	var p []Problem
	for _, pr := range problems {
		p = append(
			p,
			Problem{
				Q: strings.ToLower(strings.TrimSpace(pr[0])),
				A: strings.ToLower(strings.TrimSpace(pr[1])),
			})
	}

	return p, nil
}
