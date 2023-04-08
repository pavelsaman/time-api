package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavelsaman/time-api/service"
	"github.com/pavelsaman/time-api/types"
)

func GetEpochTime(w http.ResponseWriter, r *http.Request) {
	epochTime, err := service.GetEpochTime(mux.Vars(r)["epochType"])
	if err != nil {
		http.NotFound(w, r)
	} else {
		json.NewEncoder(w).Encode(types.EpochTimeResponse{
			Type:  mux.Vars(r)["epochType"],
			Epoch: fmt.Sprint(epochTime),
		})
	}
}
