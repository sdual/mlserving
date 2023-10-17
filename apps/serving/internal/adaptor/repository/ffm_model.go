package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/sdual/mlserving/apps/serving/internal/domain/model"
	"github.com/simonfrey/jsonl"
)

type (
	FFMModelParam struct {
		modelParams map[string]*model.CacheFFMParameter // いったんキャッシュしておく。
	}

	RawFFMParameter struct {
		Field         string               `json:"key"`
		LatentVectors map[string][]float64 `json:"value"`
	}
)

func NewFFMModelParam() FFMModelParam {
	modelBytes, err := os.ReadFile("/model.jsonl")
	if err != nil {
		log.Fatal().Msgf("failed to read model file: %v", err)
	}
	params := make([]*RawFFMParameter, 0)
	r := jsonl.NewReader(bytes.NewReader(modelBytes))

	err = r.ReadLines(func(data []byte) error {
		param := &RawFFMParameter{}
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

	paramMap := make(map[string]*model.CacheFFMParameter, len(params))

	for _, param := range params {
		latentVectors := make(map[int][]float64, len(param.LatentVectors))
		for k, v := range param.LatentVectors {
			intKey, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			latentVectors[intKey] = v
		}
		paramMap[param.Field] = &model.CacheFFMParameter{
			Field:         param.Field,
			LatentVectors: latentVectors,
		}
	}

	fmt.Println(paramMap["criteo_131071"])
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
