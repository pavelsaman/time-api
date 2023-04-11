package test_utils

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

type Handler struct {
	Url     string
	Func    func(http.ResponseWriter, *http.Request)
	Methods []string
}

type Handlers struct {
	Handlers []*Handler
}

func StartTestServerAndRegisterHandlers(handlers *Handlers) *httptest.Server {
	router := mux.NewRouter()

	for _, handler := range handlers.Handlers {
		router.HandleFunc(handler.Url, handler.Func).Methods(handler.Methods...)
	}

	return httptest.NewServer(router)
}
