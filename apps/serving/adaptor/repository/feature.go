package repository

import "github.com/sdual/mlserving/gcp/memorystore/redis"

type Feature struct {
	client *redis.Client
}

func NewFeature(client *redis.Client) Feature {
	return Feature{
		client: client,
	}
}

func (f Feature) Get() {
	f.client.Get()
}
