package model

import (
	"hash/fnv"
)

type (
	Features struct {
		Numericals   []*Numerical
		Categoricals []*Categorical
	}

	Numerical struct {
		Name  string
		Value float64
	}

	Categorical struct {
		Name  string
		Value string
	}
)

const (
	hashedVecDim int = 131072 // 2^17
)

func (c Categorical) HashingTrick() int {
	hash := fnv.New32a()
	hash.Write([]byte(c.Value))
	return int(hash.Sum32()) % hashedVecDim
}
