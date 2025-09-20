package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type migration struct {
	conn         *sql.DB
	migrationDir string
	sqlFiles     []os.DirEntry
}

func NewMigration(conn *sql.DB) (*migration, error) {
	if conn == nil {
		return nil, fmt.Errorf("database connection not initialized")
	}

	migrationDir := "migrations/"

	dirEntry, err := os.ReadDir(migrationDir)
	if err != nil {
		return nil, err
	}

	return &migration{
		conn:         conn,
		migrationDir: migrationDir,
		sqlFiles:     dirEntry,
	}, nil
}

func (m *migration) RunMigrations() error {
	tx, err := m.conn.Begin()
	if err != nil {
		return err
	}

	for _, file := range m.sqlFiles {
		if file.IsDir() || filepath.Ext(file.Name()) != ".sql" {
			continue
		}

		filePath := filepath.Join(m.migrationDir, file.Name())
		sqlFile, err := os.ReadFile(filePath)
		if err != nil {
			tx.Rollback()
			return err
		}

		query := string(sqlFile)
		log.Printf("\033[33mPrepare %s migration\033[0m\n", file.Name())

		if _, err := tx.Exec(query); err != nil {
			tx.Rollback()
			return fmt.Errorf("error running migration %s: %w", file.Name(), err)
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
