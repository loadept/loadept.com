package service

import (
	"context"

	"github.com/loadept/loadept.com/internal/repository/redis"
)

const (
	RedisStatusUp   = "up"
	RedisStatusDown = "down"
)

type CheckHealthRedisService struct {
	respository *redis.CheckHealthRedisRepository
}

func NewCheckHealthRedisService(repository *redis.CheckHealthRedisRepository) *CheckHealthRedisService {
	return &CheckHealthRedisService{
		respository: repository,
	}
}

func (s *CheckHealthRedisService) Ping(ctx context.Context) (string, error) {
	if err := s.respository.Ping(ctx); err != nil {
		return RedisStatusDown, err
	}
	return RedisStatusUp, nil
}
