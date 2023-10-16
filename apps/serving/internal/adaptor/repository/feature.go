package repository

import (
	"context"

	"github.com/sdual/mlserving/apps/serving/internal/domain/model"
	"github.com/sdual/mlserving/pkg/gcp/memorystore/redis"
)

type Feature struct {
	client *redis.Client
}

func NewFeature(client *redis.Client) Feature {
	return Feature{
		client: client,
	}
}

func (f Feature) Get(ctx context.Context, key string) (model.Features, error) {
	f.client.Get(ctx, key)
	return model.Features{}, nil
}
