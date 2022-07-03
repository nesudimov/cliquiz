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

	return m.Problems, nil
}
