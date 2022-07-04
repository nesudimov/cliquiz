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

	return m.Problems, nil
}
