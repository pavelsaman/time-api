package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavelsaman/time-api/api"
	"github.com/pavelsaman/time-api/config"
)

func registerHandlers(router *mux.Router) {
	router.HandleFunc("/"+config.ApiVersion()+"/version", api.GetVersion).Methods("GET")
	router.HandleFunc("/"+config.ApiVersion()+"/time/{epochType}", api.GetEpochTime).Methods("GET")
}

func main() {
	router := mux.NewRouter()
	registerHandlers(router)

	http.ListenAndServe(config.ApiPort(), router)
}
