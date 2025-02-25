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

type dbconnection interface {
	Connect() error
	GetNow() (*time.Time, error)
	GetDB() *sql.DB
	Close() error
}

func NewConnection() dbconnection {
	once.Do(func() {
		instance = &sqlite{}
	})
	return instance
}
