package usecase

import (
	"github.com/sdual/mlserving/apps/serving/internal/domain/service"
	"github.com/sdual/mlserving/apps/serving/internal/usecase/input"
)

type Prediction struct {
	predictor    service.FFMPredictor
	preprocessor service.FFMPreprocessor
}

func NewPediction(ffm service.FFMPredictor) Prediction {
	return Prediction{
		predictor: ffm,
	}
}

func (p Prediction) Predict(inputFeatures input.FeatureInput) ([]float64, error) {
	features := inputFeatures.Convert()
	preprocessed := p.preprocessor.BatchPreprocess(features)
	return p.predictor.BatchPredict(preprocessed)
}
