package test_utils

import (
	"net/http/httptest"

	"github.com/gorilla/mux"
	test_types "github.com/pavelsaman/time-api/tests/api/types"
)

func StartTestServerAndRegisterHandlers(handlers *test_types.Handlers) *httptest.Server {
	router := mux.NewRouter()

	for _, handler := range handlers.Handlers {
		router.HandleFunc(handler.Url, handler.Func).Methods(handler.Methods...)
	}

	return httptest.NewServer(router)
}
