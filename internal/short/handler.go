package short

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/singleflight"
)

var (
	cacheMu sync.RWMutex
	sfGroup singleflight.Group
	cache   = make(map[string]string)
)

type Handler struct {
	r *Repo
}

func NewHandler(r *Repo) *Handler {
	return &Handler{r: r}
}

func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "code")

	cacheMu.RLock()
	cachedURL, ok := cache[shortCode]
	cacheMu.RUnlock()
	if ok {
		http.Redirect(w, r, cachedURL, http.StatusFound)
		return
	}

	v, err, _ := sfGroup.Do(shortCode, func() (any, error) {
		cacheMu.RLock()
		cachedURL, ok := cache[shortCode]
		cacheMu.RUnlock()
		if ok {
			return cachedURL, nil
		}

		originalURL, err := h.r.GetURL(context.Background(), shortCode)
		if err != nil {
			return nil, err
		}

		cacheMu.Lock()
		defer cacheMu.Unlock()
		cache[shortCode] = originalURL

		return originalURL, nil
	})
	if err != nil {
		slog.Error("failed to get URL from database", "err", err)
		if errors.Is(err, ErrShortURLNotFound) {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("500 internal error"))
		return
	}

	http.Redirect(w, r, v.(string), http.StatusFound)
}
