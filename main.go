package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavelsaman/time-api/api/controllers"
	"github.com/pavelsaman/time-api/config"
	"github.com/pavelsaman/time-api/middleware"
)

func registerHandlers(router *mux.Router) {
	router.HandleFunc("/"+config.ApiVersion()+"/version", controllers.GetVersion).Methods("GET")
	// epoch
	router.HandleFunc("/"+config.ApiVersion()+"/unix/{epochType}", controllers.GetEpochTime).Methods("GET")
	router.HandleFunc("/"+config.ApiVersion()+"/unix/epoch/{epochValue}", controllers.GetEpochToUtc).Methods("GET")
	// utc
	router.HandleFunc("/"+config.ApiVersion()+"/time/utc", controllers.GetUtcTime).Methods("GET")
}

func main() {
	router := mux.NewRouter()
	router.Use(middleware.LogRequest)
	registerHandlers(router)

	fmt.Printf("Api version: %v, starting listening on port %v\n", config.ApiVersion(), config.ApiPort())
	log.Fatal(http.ListenAndServe(config.ApiPort(), router))
}
