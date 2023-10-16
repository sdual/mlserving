package service

import (
	"math"

	"github.com/sdual/mlserving/apps/serving/internal/domain/model"
)

type FFMPreprocessor struct {
	vecDim int64
}

func NewFFMPreprocessor() FFMPreprocessor {
	vecDim := math.Pow(2.0, 17) // TODO: ここで管理しない方がよさそう。定数として適切な場所に置く。
	return FFMPreprocessor{
		vecDim: int64(vecDim),
	}
}

func (fp FFMPreprocessor) Preprocess(features *model.Features) {
	_ = features.Categorical.HashingTrick()

}
