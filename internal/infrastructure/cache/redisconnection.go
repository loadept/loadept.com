package cache

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	instance redisConnection
	once     sync.Once
)

type redisConnection interface {
	Connect(context.Context) error
	GetNow() (*time.Time, error)
	GetClient() *redis.Client
	Close() error
}

func NewRedisConnection(ctx context.Context) (redisConnection, error) {
	var err error
	once.Do(func() {
		redisClient := &cache{}
		if err = redisClient.Connect(ctx); err == nil {
			instance = redisClient
		}
	})
	return instance, err
}
