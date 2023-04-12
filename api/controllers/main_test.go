package controllers

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pavelsaman/time-api/api/types"
	"github.com/pavelsaman/time-api/api/utils"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	var handlers types.Handlers
	for _, endpoint := range Endpoints {
		handlers.Handlers = append(handlers.Handlers, &types.Handler{
			Url:     endpoint.Url,
			Func:    endpoint.Func,
			Methods: endpoint.Methods,
		})
	}
	testServer = utils.StartTestServerAndRegisterHandlers(&handlers)
	defer testServer.Close()

	exitCode := m.Run()

	os.Exit(exitCode)
}
