package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CheckHealthRedisRepository struct {
	rdb *redis.Client
}

func NewCheckHealthRedisRepository(rdb *redis.Client) *CheckHealthRedisRepository {
	return &CheckHealthRedisRepository{
		rdb: rdb,
	}
}

func (c *CheckHealthRedisRepository) Ping(ctx context.Context) error {
	if _, err := c.rdb.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}
