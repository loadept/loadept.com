package main

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/loadept/pirca"
	"github.com/loadept/website/internal/middleware"
	"github.com/loadept/website/internal/short"
	"github.com/loadept/website/internal/storage"
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

//go:embed all:web/static
var staticFiles embed.FS

func main() {
	log.SetFlags(0)

	db := os.Getenv("DB_PATH")
	if db == "" {
		log.Fatalf("env var DB_PATH is required")
	}
	addr := os.Getenv("ADDR")
	if addr == "" {
		log.Fatalf("env var ADDR is required")
	}
	webhookSecret := os.Getenv("WEBHOOK_SECRET")
	if webhookSecret == "" {
		log.Fatalf("env var WEBHOOK_SECRET is required")
	}
	webhook := os.Getenv("WEBHOOK")
	if webhook == "" {
		log.Fatalf("env var WEBHOOK is required")
	}

	pool, err := sqlitex.NewPool(db, sqlitex.PoolOptions{
		PoolSize: 3,
		PrepareConn: func(conn *sqlite.Conn) error {
			return sqlitex.ExecuteScript(conn, `
				PRAGMA foreign_keys = ON;
				PRAGMA busy_timeout = 5000;
				PRAGMA journal_mode = WAL;
			`, nil)
		},
	})
	fatalIfErr(err)
	defer pool.Close()

	mux := http.NewServeMux()
	visit := middleware.NewVisitsMiddleware(
		&http.Client{Timeout: 10 * time.Second},
		webhookSecret,
		webhook,
	)

	s, err := storage.NewShortURLStorage(pool)
	fatalIfErr(err)

	shortHandler := short.NewShortHandler(s)
	subFS, err := fs.Sub(staticFiles, "web/static")
	fatalIfErr(err)

	mux.Handle("GET /", http.FileServerFS(&neuteredFS{fs: subFS}))
	mux.Handle("GET /s/{code}", visit(http.HandlerFunc(shortHandler.RedirectURL)))

	server := http.Server{
		Addr:         addr,
		Handler:      pirca.New()(middleware.Logger(mux)),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		log.Println("server listen on", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("shutting down server...")
	fatalIfErr(server.Shutdown(shutdownCtx))
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// FileServer Wrapper
type neuteredFS struct{ fs fs.FS }

func (n *neuteredFS) Open(name string) (fs.File, error) {
	name = strings.TrimPrefix(name, "/")
	if name == "" {
		name = "."
	}

	f, err := n.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, errors.Join(err, f.Close())
	}

	if stat.IsDir() {
		_, err := fs.Stat(n.fs, path.Join(name, "index.html"))
		if err != nil {
			return nil, errors.Join(fs.ErrNotExist, f.Close())
		}
	}

	return f, nil
}
