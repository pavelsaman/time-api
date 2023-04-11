package utils

import (
	"encoding/json"
	"net/http"

	"github.com/pavelsaman/time-api/api/types"
)

func SendApiError(w *http.ResponseWriter, err *types.ApiErrorResponse) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(err.ErrorCode)
	json.NewEncoder(*w).Encode(err)
}
