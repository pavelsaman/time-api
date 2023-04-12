package utils

import (
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/pavelsaman/time-api/api/types"
)

func StartTestServerAndRegisterHandlers(handlers *types.Handlers) *httptest.Server {
	router := mux.NewRouter()

	for _, handler := range handlers.Handlers {
		router.HandleFunc(handler.Url, handler.Func).Methods(handler.Methods...)
	}

	return httptest.NewServer(router)
}
