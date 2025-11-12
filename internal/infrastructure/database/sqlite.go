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

// Connect establishes a connection to the SQLite database.
//
// This method should only be called internally, as
//
//	NewConnection()
//
// ensures a single instance.
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
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

// GetNow returns the current date from the database as type
//
//	*time.Time
func (s *sqlite) getNow() (*time.Time, error) {
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

// GetDB returns the
//
//	db
//
// field of type "sqlite" that contains an underlying
//
//	*sql.DB
//
// instance for direct database access.
// Use this method if raw SQL queries or transactions are required.
func (s *sqlite) GetDB() *sql.DB {
	return s.db
}

// Close closes the connection created to the database
func (s *sqlite) Close() error {
	return s.db.Close()
}
