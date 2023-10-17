package repository

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/sdual/mlserving/apps/serving/internal/domain/model"
	"github.com/simonfrey/jsonl"
)

type FFMModelParam struct {
	modelParams map[string]*model.RawFFMParameter // いったんキャッシュしておく。
}

func NewFFMModelParam() FFMModelParam {
	modelBytes, err := os.ReadFile("model.jsonl")
	if err != nil {
		log.Fatal().Msgf("failed to read model file: %v", err)
	}
	params := make([]*model.RawFFMParameter, 0)
	r := jsonl.NewReader(bytes.NewReader(modelBytes))

	err = r.ReadLines(func(data []byte) error {
		param := &model.RawFFMParameter{}
		err := json.Unmarshal(data, param)
		if err != nil {
			return err
		}
		params = append(params, param)
		return nil
	})
	if err != nil {
		log.Fatal().Msgf("failed to read model file: %v", err)
	}

	paramMap := make(map[string]*model.RawFFMParameter, len(params))

	for _, param := range params {
		paramMap[param.Field] = param
	}
	return FFMModelParam{
		modelParams: paramMap,
	}
}

func (f FFMModelParam) Find(fields []*model.FFMModelField) ([]*model.FFMParameter, error) {
	params := make([]*model.FFMParameter, len(fields))
	for i, field := range fields {
		if p, ok := f.modelParams[field.TargetField]; ok {
			params[i] = &model.FFMParameter{
				FieldIndex:    field.FieldIndex,
				LatentVectors: p.LatentVectors,
			}
		}
	}
	return params, nil
}
