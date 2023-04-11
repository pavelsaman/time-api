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
	for _, endpoint := range controllers.Endpoints {
		router.HandleFunc(endpoint.Url, endpoint.Func).Methods(endpoint.Methods...)	
	}
}

func main() {
	router := mux.NewRouter()
	router.Use(middleware.LogRequest)
	registerHandlers(router)

	fmt.Printf("Api version: %v, starting listening on port %v\n", config.ApiVersion(), config.ApiPort())
	log.Fatal(http.ListenAndServe(config.ApiPort(), router))
}
