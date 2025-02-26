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
