package database

import (
	"database/sql"
	"log"
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
	getNow() (*time.Time, error)
	GetDB() *sql.DB
	Close() error
}

// NewConnection returns a new SQLite connection as a dbconnection instance.
//
// This function ensures a single instance and should not call Connect() again.
func NewConnection() (dbconnection, error) {
	var err error
	once.Do(func() {
		instance = &sqlite{}
		if err = instance.Connect(); err == nil {
			var now *time.Time

			if now, err = instance.getNow(); err == nil {
				log.Printf("DB connection established, current date %s\n", now.Format("2006-01-02"))
			}
		}
	})
	return instance, err
}
