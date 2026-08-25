package short

import (
	"context"
	"errors"

	"zombiezen.com/go/sqlite/sqlitex"
)

type Repo struct {
	pool *sqlitex.Pool
}

func NewRepo(pool *sqlitex.Pool) *Repo { return &Repo{pool: pool} }

var ErrShortURLNotFound = errors.New("short url does not exist or invalid short code")

const (
	URLStatusActive  = "active"
	URLStatusDeleted = "deleted"
)

func (r *Repo) GetURL(ctx context.Context, shortCode string) (string, error) {
	conn, err := r.pool.Take(ctx)
	if err != nil {
		return "", err
	}
	defer r.pool.Put(conn)

	stmt, err := conn.Prepare(`
		SELECT original_url
		FROM short_urls
		WHERE short_code = $1
			AND status = $2
	`)
	if err != nil {
		return "", err
	}
	defer stmt.Reset()

	stmt.SetText("$1", shortCode)
	stmt.SetText("$2", URLStatusActive)

	if hasRow, err := stmt.Step(); err != nil {
		return "", err
	} else if !hasRow {
		return "", ErrShortURLNotFound
	}
	return stmt.GetText("original_url"), nil
}
