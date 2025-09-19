package service

import (
	"github.com/loadept/loadept.com/internal/repository"
)

const (
	DBStatusUp   = "up"
	DBStatusDown = "down"
)

type CheckHealthDBService struct {
	respository *repository.CheckHealthDBRepository
}

func NewCheckHealthDBService(repository *repository.CheckHealthDBRepository) *CheckHealthDBService {
	return &CheckHealthDBService{
		respository: repository,
	}
}

func (s *CheckHealthDBService) Ping() (string, error) {
	if err := s.respository.Ping(); err != nil {
		return DBStatusDown, err
	}
	return DBStatusUp, nil
}
