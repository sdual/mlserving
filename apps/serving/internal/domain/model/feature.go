package model

import (
	"hash/fnv"

	"github.com/james-bowman/sparse"
)

type (
	Features struct {
		Numerical   NumericalFeatures
		Categorical CategoricalFeatures
	}

	NumericalFeatures   [][]float64
	CategoricalFeatures [][]string

	HashedFeatures struct {
		NumericalLen int
		Numerical    NumericalFeatures
		Hashed       *sparse.DOK
	}
)

func (f HashedFeatures) At(i, j int) float64 {
	if j < f.NumericalLen {
		return f.Numerical[i][j]
	} else {
		colIndex := j - f.NumericalLen
		return f.Hashed.At(i, colIndex)
	}
}

const (
	hashedVecDim int = 131072 // 2^17
)

func (cs CategoricalFeatures) HashingTrick() *sparse.DOK {
	rowNum := len(cs)
	colNum := len(cs[0])
	dokMatrix := sparse.NewDOK(rowNum, colNum)

	for i, row := range cs {
		for _, feature := range row {
			hash := fnv.New32a()
			hash.Write([]byte(feature))
			colIndex := int(hash.Sum32()) % hashedVecDim
			dokMatrix.Set(i, colIndex, 1.0)
		}
	}
	return dokMatrix
}
