package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavelsaman/time-api/api/services"
	"github.com/pavelsaman/time-api/api/types"
	"github.com/pavelsaman/time-api/api/utils"
)

func GetEpochTime(w http.ResponseWriter, r *http.Request) {
	epochTime, err := services.GetEpochTime(mux.Vars(r)["epochType"])
	if err != nil {
		utils.SendApiError(&w, &types.ApiErrorResponse{
			ErrorCode:    400,
			ErrorMessage: err.Error(),
		})
	} else {
		json.NewEncoder(w).Encode(types.EpochAndUtcTimeResponse{
			Type:  epochTime.Type,
			Epoch: fmt.Sprint(epochTime.Epoch),
			Utc:   fmt.Sprint(epochTime.Utc),
		})
	}
}

func GetEpochToUtc(w http.ResponseWriter, r *http.Request) {
	utc, err := services.GetEpochToUtc(mux.Vars(r)["epochValue"])
	if err != nil {
		utils.SendApiError(&w, &types.ApiErrorResponse{
			ErrorCode:    400,
			ErrorMessage: err.Error(),
		})
	} else {
		json.NewEncoder(w).Encode(types.EpochToUtcTimeResponse{
			Utc:       utc.Utc,
			EpochType: utc.Type,
		})
	}
}
