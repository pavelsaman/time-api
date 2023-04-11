package api_test

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pavelsaman/time-api/api/controllers"
	test_types "github.com/pavelsaman/time-api/tests/api/types"
	test_utils "github.com/pavelsaman/time-api/tests/api/utils"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	var handlers test_types.Handlers
	for _, endpoint := range controllers.Endpoints {
		handlers.Handlers = append(handlers.Handlers, &test_types.Handler{
			Url:     endpoint.Url,
			Func:    endpoint.Func,
			Methods: endpoint.Methods,
		})
	}
	testServer = test_utils.StartTestServerAndRegisterHandlers(&handlers)
	defer testServer.Close()

	exitCode := m.Run()

	os.Exit(exitCode)
}
