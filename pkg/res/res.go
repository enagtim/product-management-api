package res

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&data)
}
