package handler

import (
	"fmt"
	"net/http"
	"time"

	application "github.com/loadept/loadept.com/internal/application/checkhealth"
	"github.com/loadept/loadept.com/internal/config"
	domain "github.com/loadept/loadept.com/internal/domain/checkhealth"
	"github.com/loadept/loadept.com/pkg/respond"
)

type CheckHealthHandler struct {
	checkhealthService *application.CheckHealthService
}

func NewHealthHandler(checkhealthService *application.CheckHealthService) *CheckHealthHandler {
	return &CheckHealthHandler{
		checkhealthService: checkhealthService,
	}
}

func (h *CheckHealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	response := domain.CheckHealth{
		Timestamp: time.Now().Format(time.RFC3339),
		Timezone:  time.Now().Location().String(),
	}
	response.Status = "ok"

	rdbPing, err := h.checkhealthService.CheckCacheConnection(r.Context())
	if err != nil {
		response.Services.Redis = rdbPing
		response.Status = "bad"
	}
	dbPing, err := h.checkhealthService.CheckDBConnection(r.Context())
	if err != nil {
		response.Services.Database = dbPing
		response.Status = "bad"
	}

	response.Services.Redis = rdbPing
	response.Services.Database = dbPing
	response.UptimeS = time.Since(config.ApplicationUptime).String()

	response.ResponseTimeMs = fmt.Sprintf("%dms", time.Since(start).Milliseconds())
	respond.JSON(w, response, http.StatusOK)
}
