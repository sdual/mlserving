package repository

import "github.com/sdual/mlserving/apps/serving/internal/domain/model"

type (
	FFMModelParamRepository interface {
		Find(fields []*model.FFMModelField) ([]*model.FFMParameter, error)
	}
)
