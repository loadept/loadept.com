package cache

import (
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	instance redisConnection
	once sync.Once
)

type redisConnection interface {
	Connect() error
	GetNow() (*time.Time, error)
	GetClient() *redis.Client
	Close() error
}

func NewRedisConnection() (redisConnection, error) {
	var err error
	once.Do(func() {
		redisClient := &cache{}
		if err = redisClient.Connect(); err == nil {
			instance = redisClient
		}
	})
	return instance, err
}
