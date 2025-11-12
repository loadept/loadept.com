package respond

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Map map[string]any

func JSON(w http.ResponseWriter, response any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"detail":"Internal server error"}`)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(buf.Bytes())
}
