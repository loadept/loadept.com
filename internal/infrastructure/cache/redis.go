package cache

import (
	"fmt"
	"time"

	"github.com/loadept/loadept.com/internal/config"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type cache struct {
	client *redis.Client
	ctx    context.Context
}

func (c *cache) Connect(ctx context.Context) error {
	c.ctx = ctx

	uri := fmt.Sprintf("%s:%s", config.Env.REDIS_ADDR, config.Env.REDIS_PORT)
	options := &redis.Options{
		Addr:     uri,
		Password: config.Env.REDIS_PASSWORD,
		DB:       0,
		Protocol: 2,
	}

	client := redis.NewClient(options)

	if _, err := client.Ping(c.ctx).Result(); err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *cache) GetNow() (*time.Time, error) {
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
