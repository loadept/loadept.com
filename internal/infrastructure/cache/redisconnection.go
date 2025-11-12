package cache

import (
	"context"
	"log"
	"sync"
	"time"
)

var (
	instance *cache
	once     sync.Once
)

func NewRedisConnection(ctx context.Context) (*cache, error) {
	var err error

	once.Do(func() {
		instance = &cache{}
		if err = instance.Connect(ctx); err == nil {
			var now *time.Time

			if now, err = instance.getNow(); err == nil {
				log.Printf("Redis connection established, current date: %s\n", now.Format("2006-01-02"))
			}
		}
	})
	return instance, err
}
