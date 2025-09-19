package repository

import "database/sql"

type CheckHealthDBRepository struct {
	db *sql.DB
}

func NewCheckHealthDBRepository(db *sql.DB) *CheckHealthDBRepository{
	return &CheckHealthDBRepository{db: db}
}

func (c *CheckHealthDBRepository) Ping() error {
	if err := c.db.Ping(); err != nil {
		return err
	}

	return nil
}
