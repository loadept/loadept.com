package redis

import (
	"context"

	domain "github.com/loadept/loadept.com/internal/domain/checkhealth"
	"github.com/redis/go-redis/v9"
)

type CheckHealthRedisRepository struct {
	rdb *redis.Client
}

func NewCheckHealthRedisRepository(rdb *redis.Client) domain.CheckHealthRepository {
	return &CheckHealthRedisRepository{
		rdb: rdb,
	}
}

func (c *CheckHealthRedisRepository) CheckConnection(ctx context.Context) error {
	if _, err := c.rdb.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}
