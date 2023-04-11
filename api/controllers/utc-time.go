package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pavelsaman/time-api/api/types"
)

func GetUtcTime(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(types.UtcTimeResponse{
		Type: "utc",
		Utc:  fmt.Sprint(time.Now().UTC()),
	})
}
