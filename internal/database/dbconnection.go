package database

import (
	"database/sql"
	"sync"
	"time"
)

var (
	instance dbconnection
	once     sync.Once
)

// dbconnection defines methods needed to create a connection to sqlite.
type dbconnection interface {
	Connect() error
	GetNow() (*time.Time, error)
	GetDB() *sql.DB
	Close() error
}

// NewConnection returns a new SQLite connection as a dbconnection instance.
//
// This function ensures a single instance and should not call Connect() again.
func NewConnection() (dbconnection, error) {
	var err error
	once.Do(func() {
		sqliteDB := &sqlite{}
		if err = sqliteDB.Connect(); err == nil {
			instance = sqliteDB
		}
	})
	return instance, err
}
