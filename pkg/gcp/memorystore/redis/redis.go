package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	Client struct {
		client *redis.Client
	}

	Option func(*redisOptions)

	redisOptions struct {
		password string
		db       int
	}
)

func NewClient(addr string, opts ...Option) *Client {
	redisOpts := &redisOptions{
		password: "",
		db:       0,
	}
	for _, opt := range opts {
		opt(redisOpts)
	}

	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisOpts.password,
		DB:       redisOpts.db,
	})
	return &Client{
		client: c,
	}
}

func WithDB(db int) Option {
	return func(c *redisOptions) {
		c.db = db
	}
}

func WithPassword(password string) Option {
	return func(c *redisOptions) {
		c.password = password
	}
}

func (c Client) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (c Client) Set(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	err := c.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
