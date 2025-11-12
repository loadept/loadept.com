package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/loadept/loadept.com/internal/config"
	"github.com/redis/go-redis/v9"
)

type cache struct {
	client *redis.Client
	ctx    context.Context
}

func (c *cache) Connect(ctx context.Context) error {
	c.ctx = ctx

	var redisScheme string
	if config.Env.DEBUG == "true" {
		redisScheme = "redis"
	} else {
		redisScheme = "rediss"
	}

	redisURL := fmt.Sprintf("%s://%s:%s@%s:%s",
		redisScheme,
		config.Env.REDIS_USER,
		config.Env.REDIS_PASSWORD,
		config.Env.REDIS_HOST,
		config.Env.REDIS_PORT,
	)

	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return err
	}

	client := redis.NewClient(options)

	if _, err := client.Ping(c.ctx).Result(); err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *cache) getNow() (*time.Time, error) {
	timeResult, err := c.client.Time(c.ctx).Result()
	if err != nil {
		return nil, err
	}

	rdTime := timeResult.UTC()

	return &rdTime, nil
}

func (c *cache) GetClient() *redis.Client {
	return c.client
}

func (c *cache) Close() error {
	return c.client.Close()
}
