package internal

import "encoding/json"

type JsonFile struct {
	D *json.Decoder
}

func (jf *JsonFile) LoadProblem() ([]Problem, error) {
	var m struct {
		Problems []Problem
	}

	err := jf.D.Decode(&m)
	if err != nil {
		return nil, err
	}

	for i, p := range m.Problems {
		m.Problems[i].S = len(p.A)
	}

	return m.Problems, nil
}
