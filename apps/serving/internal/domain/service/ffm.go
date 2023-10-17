package service

import (
	"math"

	"github.com/sdual/mlserving/apps/serving/internal/domain/model"
	"github.com/sdual/mlserving/apps/serving/internal/domain/repository"
)

type FFMPredictor struct {
	paramRepository repository.FFMModelParamRepository
}

func NewFFMPredictor(paramRepo repository.FFMModelParamRepository) FFMPredictor {
	return FFMPredictor{
		paramRepository: paramRepo,
	}
}

func (fp FFMPredictor) BatchPredict(modelFieldsList [][]*model.FFMModelField) ([]float64, error) {
	preds := make([]float64, len(modelFieldsList))
	for i, modelFields := range modelFieldsList {
		pred, err := fp.Predict(modelFields)
		if err != nil {
			return nil, err
		}
		preds[i] = pred
	}
	return preds, nil
}

func (fp FFMPredictor) Predict(modelFields []*model.FFMModelField) (float64, error) {
	params, err := fp.paramRepository.Find(modelFields)
	if err != nil {
		return 0.0, err
	}

	exponent := 0.0
	for i := 0; i < len(params); i++ {
		for j := i + 1; j < len(params); j++ {
			p1 := params[i]
			p2 := params[j]
			exponent += fp.innerProduct(p1.LatentVectors[p2.Field], p2.LatentVectors[p1.Field])
		}
	}
	paramLen := len(params)
	exponent = exponent / float64(paramLen)
	return sigmoid(exponent), nil
}

func (fp FFMPredictor) innerProduct(v1, v2 []float64) float64 {
	result := 0.0
	for i := 0; i < len(v1); i++ {
		result += v1[i] * v2[i]
	}
	return result
}

func sigmoid(exponent float64) float64 {
	return 1.0 / (1.0 + math.Exp(-exponent))
}
