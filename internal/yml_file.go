package internal

import yaml "gopkg.in/yaml.v3"

type YmlFile struct {
	D *yaml.Decoder
}

func (yf *YmlFile) LoadProblem() ([]Problem, error) {
	var m struct {
		Problems []Problem
	}

	err := yf.D.Decode(&m)
	if err != nil {
		return nil, err
	}

	for i, p := range m.Problems {
		m.Problems[i].S = len(p.A)
	}

	return m.Problems, nil
}
