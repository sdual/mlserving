package model

type (
	Features struct {
		numerical   NumericalFeatures
		categorical CategoricalFeatures
	}

	NumericalFeatures   []float64
	CategoricalFeatures []string
)
