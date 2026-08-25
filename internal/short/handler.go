package short

import (
	"errors"
	"log/slog"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

var (
	cacheMu sync.RWMutex
	cache   = make(map[string]string)
)

type Handler struct {
	r *Repo
}

func NewHandler(r *Repo) *Handler {
	return &Handler{r: r}
}

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	shortCode := chi.URLParam(r, "code")

	cacheMu.RLock()
	if cachedURL, ok := cache[shortCode]; ok {
		cacheMu.RUnlock()
		http.Redirect(w, r, cachedURL, http.StatusFound)
		return
	}
	cacheMu.RUnlock()

	originalURL, err := h.r.GetURL(ctx, shortCode)
	if err != nil {
		slog.Error("failed to get URL from database", "err", err)
		if errors.Is(err, ErrShortURLNotFound) {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("500 internal error"))
		return
	}
	cacheMu.Lock()
	cache[shortCode] = originalURL
	cacheMu.Unlock()

	http.Redirect(w, r, originalURL, http.StatusFound)
}
