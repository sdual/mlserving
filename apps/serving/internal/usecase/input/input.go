package input

import "github.com/sdual/mlserving/apps/serving/internal/domain/model"

type (
	FeatureInput struct {
		Features []*Feature
	}

	Feature struct {
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

func (fi FeatureInput) Convert() []*model.Features {
	features := make([]*model.Features, len(fi.Features))
	for i, f := range fi.Features {
		numericals := make([]*model.Numerical, len(f.Numericals))
		for j, n := range f.Numericals {
			numericals[j] = &model.Numerical{
				Name:  n.Name,
				Value: n.Value,
			}
		}

		categoricals := make([]*model.Categorical, len(f.Categoricals))
		for k, c := range f.Categoricals {
			categoricals[k] = &model.Categorical{
				Name:  c.Name,
				Value: c.Value,
			}
		}

		features[i] = &model.Features{
			Numericals:   numericals,
			Categoricals: categoricals,
		}
	}
	return features
}
