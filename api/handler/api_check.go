package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/service"
	"github.com/loadept/loadept.com/pkg/respond"
)

type HealthHandler struct {
	rdbService *service.CheckHealthRedisService
	dbService  *service.CheckHealthDBService
}

func NewHealthHandler(rdbService *service.CheckHealthRedisService, dbService *service.CheckHealthDBService) *HealthHandler {
	return &HealthHandler{
		rdbService: rdbService,
		dbService:  dbService,
	}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	response := model.HealthModel{
		Timestamp: time.Now().Format(time.RFC3339),
		Timezone:  time.Now().Location().String(),
	}
	response.Status = "ok"

	rdbPing, err := h.rdbService.Ping(r.Context())
	if err != nil {
		response.Services.Redis = rdbPing
		response.Status = "bad"
	}
	dbPing, err := h.dbService.Ping()
	if err != nil {
		response.Services.Database = dbPing
		response.Status = "bad"
	}

	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	response.Services.Redis = rdbPing
	response.Services.Database = dbPing
	response.UptimeS = time.Since(config.ApplicationUptime).String()

	response.ResponseTimeMs = fmt.Sprintf("%dms", time.Since(start).Milliseconds())
	respond.JSON(w, response, http.StatusOK)
}
