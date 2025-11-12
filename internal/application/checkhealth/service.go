package checkhealth

import (
	"context"

	domain "github.com/loadept/loadept.com/internal/domain/checkhealth"
)

const (
	DBStatusUp      = "up"
	DBStatusDown    = "down"
	RedisStatusUp   = "up"
	RedisStatusDown = "down"
)

type CheckHealthService struct {
	respository      domain.CheckHealthRepository
	respositoryCache domain.CheckHealthRepository
}

func NewCheckHealthService(
	repository domain.CheckHealthRepository,
	repositoryCache domain.CheckHealthRepository,
) *CheckHealthService {
	return &CheckHealthService{
		respository:      repository,
		respositoryCache: repositoryCache,
	}
}

func (s *CheckHealthService) CheckDBConnection(ctx context.Context) (string, error) {
	if err := s.respository.CheckConnection(ctx); err != nil {
		return DBStatusDown, err
	}

	return DBStatusUp, nil
}

func (s *CheckHealthService) CheckCacheConnection(ctx context.Context) (string, error) {
	if err := s.respositoryCache.CheckConnection(ctx); err != nil {
		return RedisStatusDown, err
	}

	return RedisStatusUp, nil
}
