package throwable

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, response any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"detail": "Internal server error"}`, http.StatusInternalServerError)
	}
}
