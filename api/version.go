package api

import (
	"encoding/json"
	"net/http"

	"github.com/pavelsaman/time-api/config"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		Version string `json:"version"`
	}{
		config.ApiVersion(),
	})
}
