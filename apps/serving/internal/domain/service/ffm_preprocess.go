package service

import (
	"math"
	"strconv"

	"github.com/sdual/mlserving/apps/serving/internal/domain/model"
)

type FFMPreprocessor struct {
	vecDim int64
}

const prefix = "criteo_"

func NewFFMPreprocessor() FFMPreprocessor {
	vecDim := math.Pow(2.0, 17) // TODO: ここで管理しない方がよさそう。定数として適切な場所に置く。
	return FFMPreprocessor{
		vecDim: int64(vecDim),
	}
}

func (fp FFMPreprocessor) BatchPreprocess(featureList []*model.Features) [][]*model.FFMModelField {
	ffmFields := make([][]*model.FFMModelField, len(featureList))
	for i, features := range featureList {
		ffmFields[i] = fp.Preprocess(features)
	}
	return ffmFields
}

func (fp FFMPreprocessor) Preprocess(features *model.Features) []*model.FFMModelField {
	ffmFields := make([]*model.FFMModelField, len(features.Numericals)+len(features.Categoricals))
	index := 0
	for _, num := range features.Numericals {
		ffmFields[index] = fp.preprocessNumerical(*num)
		index++
	}

	for _, cat := range features.Categoricals {
		ffmFields[index] = fp.preprocessCat(*cat)
		index++
	}
	return ffmFields
}

func (fp FFMPreprocessor) preprocessCat(cat model.Categorical) *model.FFMModelField {
	index := cat.HashingTrick()
	field := prefix + strconv.Itoa(index)
	return &model.FFMModelField{
		FieldIndex:  model.CriteoClickFeatureIndex[cat.Name],
		TargetField: field,
	}
}

func (fp FFMPreprocessor) preprocessNumerical(num model.Numerical) *model.FFMModelField {
	index := model.CriteoClickFeatureIndex[num.Name]
	field := prefix + strconv.Itoa(index)
	return &model.FFMModelField{
		FieldIndex:  index,
		TargetField: field,
	}
}
