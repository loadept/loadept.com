package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/loadept/loadept.com/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	db *sql.DB
}

func (s *sqlite) Connect() error {
	DB_NAME := config.Env.DB_NAME
	if len(DB_NAME) == 0 {
		return fmt.Errorf("DB_NAME variable is not defined")
	}

	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *sqlite) GetNow() (*time.Time, error) {
	var t string
	err := s.db.QueryRow("SELECT DATE('now');").Scan(&t)
	if err != nil {
		return nil, err
	}

	currentTime, err := time.Parse("2006-01-02", t)
	if err != nil {
		return nil, err
	}

	return &currentTime, nil
}

func (s *sqlite) GetDB() *sql.DB {
	return s.db
}

func (s *sqlite) Close() error {
	return s.db.Close()
}
