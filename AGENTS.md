# AGENTS.md

Personal site loadept.com: Go backend (chi + zombiezen/sqlite) serving an embedded Astro static site plus a URL shortener (`/s/{code}`). README is in Spanish.

## Critical: frontend must be built before Go builds

- The Astro build output lives in `ui/dist` (gitignored; root `dist/` rule also covers it).
- `ui/embed.go` uses `//go:embed all:dist`, so on a fresh clone `go build`/`go test ./...` fails until the frontend is built:
  1. `bun run install:ui` (installs deps in `ui/`)
  2. `bun run build` (Astro -> `ui/dist`)
- All bun scripts live at the root `package.json` and delegate via `--cwd ui`: `bun run dev`, `build`, `preview`, `astro`.

## Layout

- `server.go` (root): the web server binary. Requires env vars `DB_PATH` and `ADDR` or it exits fatally. Routes: `/s/{code}` redirect + static file server for the embedded site.
- `cmd/migrate`: applies the SQLite schema (inline in its `main.go`, no migration files). Run: `go run ./cmd/migrate -db "$DB_PATH" -run`.
- `cmd/shortener`: CLI to insert short URLs: `-name -url [-code]` (base62 code, max 10 chars if custom).
- `internal/`: config (YAML loader), middleware, short handler + its SQLite repository. This is where tests live.
- `.agents/skills/`: agent skills used by AI tooling; versioned in the repo.

## Verify

```sh
gofumpt -l .            # formatter used by CI (format before pushing)
gotestsum --format testdox ./internal/...   # CI runs only ./internal/...
go test ./internal/...
```

Local `docker build` needs `--network=host` (bridge MTU truncates registry tarballs behind VPN; integrity check failures on native packages).

## CI / deploy

Pushing to `master` triggers a chained pipeline: Test Project -> Build & Push Docker image (ghcr.io) -> SSH deploy to production (`docker compose up -d`). Merging to master = production deploy.

## Conventions

- Commit messages: short, lowercase, no conventional-commit prefixes (e.g. `fix go version on dockerfile`).
- Go toolchain pinned to 1.26.x; frontend runs on Bun (lockfile is `ui/bun.lock`, CI Docker build uses `--frozen-lockfile`).
