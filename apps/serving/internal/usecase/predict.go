package usecase

import "github.com/sdual/mlserving/apps/serving/internal/domain/service"

type Prediction struct {
	predictor service.FFMPredictor
}

func NewPediction(ffm service.FFMPredictor) Prediction {
	return Prediction{
		predictor: ffm,
	}
}

func (p Prediction) Predict() []float64 {
	return p.predictor.Predict()
}
