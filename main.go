package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavelsaman/time-api/api"
	"github.com/pavelsaman/time-api/config"
	"github.com/pavelsaman/time-api/middleware"
)

func registerHandlers(router *mux.Router) {
	router.HandleFunc("/"+config.ApiVersion()+"/version", api.GetVersion).Methods("GET")
	router.HandleFunc("/"+config.ApiVersion()+"/time/{epochType}", api.GetEpochTime).Methods("GET")
	router.HandleFunc("/"+config.ApiVersion()+"/time/epoch/{epochValue}", api.GetEpochToUtc).Methods("GET")
}

func main() {
	router := mux.NewRouter()
	router.Use(middleware.LogRequest)
	registerHandlers(router)

	fmt.Printf("Api version: %v, starting listening on port %v\n", config.ApiVersion(), config.ApiPort())
	log.Fatal(http.ListenAndServe(config.ApiPort(), router))
}