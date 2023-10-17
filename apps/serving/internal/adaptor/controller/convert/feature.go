package convert

import (
	"github.com/sdual/mlserving/apps/serving/internal/usecase/input"
	pb "github.com/sdual/mlserving/proto/grpc/serving/predict"
)

func ConvertToFeatures(req *pb.PredictRequest) input.FeatureInput {
	features := make([]*input.Feature, len(req.Features))
	for i, reqFeature := range req.Features {
		numericals := make([]*input.Numerical, len(reqFeature.NumericalFeatures))
		for j, reqNumerical := range reqFeature.NumericalFeatures {
			numericals[j] = &input.Numerical{
				Name:  reqNumerical.Name,
				Value: reqNumerical.Value,
			}
		}

		categoricals := make([]*input.Categorical, len(reqFeature.CategoricalFeatures))
		for k, reqCat := range reqFeature.CategoricalFeatures {
			categoricals[k] = &input.Categorical{
				Name:  reqCat.Name,
				Value: reqCat.Value,
			}
		}

		features[i] = &input.Feature{
			Numericals:   numericals,
			Categoricals: categoricals,
		}
	}
	return input.FeatureInput{
		Features: features,
	}
}
