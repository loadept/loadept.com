package short

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/loadept/pirca"
	"github.com/loadept/website/internal/middleware"
	"github.com/loadept/website/internal/storage"
)

// local cache vars
var (
	cacheMu sync.RWMutex
	cache   = make(map[string]string)
)

type shortHandler struct {
	s *storage.ShortURLStorage
}

func NewShortHandler(s *storage.ShortURLStorage) *shortHandler {
	return &shortHandler{s: s}
}

func (h *shortHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	ctx := pirca.Ctx(r)

	val, _ := ctx.Get("logentry")
	le := val.(*middleware.LogEntry)

	shortCode := ctx.Param("code")

	cacheMu.RLock()
	if cachedURL, ok := cache[shortCode]; ok {
		cacheMu.RUnlock()
		le.CacheHit = true
		ctx.Redirect(http.StatusFound, cachedURL)
		return
	}
	cacheMu.RUnlock()

	originalURL, err := h.s.GetURL(ctx, shortCode)
	if err != nil {
		le.Error = fmt.Sprintf("failed to get URL from database: %v", err)
		if errors.Is(err, storage.ErrShortURLNotFound) {
			http.NotFound(w, r)
			return
		}
		ctx.String(http.StatusInternalServerError, "500 internal error")
		return
	}
	cacheMu.Lock()
	cache[shortCode] = originalURL
	cacheMu.Unlock()

	ctx.Redirect(http.StatusFound, originalURL)
}
