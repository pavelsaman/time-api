package controllers

import (
	"encoding/json"
	"net/http"
)

func Describe(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: GetAllControllerPaths(),
	})
}
