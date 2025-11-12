package db

import (
	"context"
	"database/sql"

	domain "github.com/loadept/loadept.com/internal/domain/checkhealth"
)

type CheckHealthDBRepository struct {
	db *sql.DB
}

func NewCheckHealthDBRepository(db *sql.DB) domain.CheckHealthRepository {
	return &CheckHealthDBRepository{db: db}
}

func (c *CheckHealthDBRepository) CheckConnection(ctx context.Context) error {
	if err := c.db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}
