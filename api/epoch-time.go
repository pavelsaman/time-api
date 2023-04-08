package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	api_utils "github.com/pavelsaman/time-api/api/utils"
	"github.com/pavelsaman/time-api/service"
	"github.com/pavelsaman/time-api/types"
)

func GetEpochTime(w http.ResponseWriter, r *http.Request) {
	epochTime, err := service.GetEpochTime(mux.Vars(r)["epochType"])
	if err != nil {
		api_utils.SendApiError(w, &types.ApiError{
			ErrorCode:    400,
			ErrorMessage: err.Error(),
		})
	} else {
		json.NewEncoder(w).Encode(types.EpochTimeResponse{
			Type:  epochTime.Type,
			Epoch: fmt.Sprint(epochTime.Epoch),
			Utc:   fmt.Sprint(epochTime.Utc),
		})
	}
}

func GetEpochToUtc(w http.ResponseWriter, r *http.Request) {
	utc, err := service.GetEpochToUtc(mux.Vars(r)["epochValue"])
	if err != nil {
		api_utils.SendApiError(w, &types.ApiError{
			ErrorCode:    400,
			ErrorMessage: "Invalid epoch time in path parameter",
		})
	} else {
		json.NewEncoder(w).Encode(types.EpochToUtcTimeResponse{
			Utc:       utc.Utc,
			EpochType: utc.Type,
		})
	}
}
