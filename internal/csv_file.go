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
		var v []string
		if len(pr[2]) != 0 {
			v = strings.Split(strings.ToLower(pr[2]), ",")
		}

		a := strings.Split(strings.ToLower(pr[1]), ",")
		p = append(
			p,
			Problem{
				Q: strings.ToLower(strings.TrimSpace(pr[0])),
				A: a,
				V: v,
				S: len(a),
			})
	}

	return p, nil
}
