package api_test

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pavelsaman/time-api/api/controllers"
	"github.com/pavelsaman/time-api/config"
	test_types "github.com/pavelsaman/time-api/tests/api/types"
	test_utils "github.com/pavelsaman/time-api/tests/api/utils"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	testServer = test_utils.StartTestServerAndRegisterHandlers(&test_types.Handlers{
		Handlers: []*test_types.Handler{
			{
				Url:     "/" + config.ApiVersion() + "/unix/{epochType}",
				Func:    controllers.GetEpochTime,
				Methods: []string{"GET"},
			},
			{
				Url:     "/" + config.ApiVersion() + "/unix/epoch/{epochValue}",
				Func:    controllers.GetEpochTime,
				Methods: []string{"GET"},
			},
			{
				Url:     "/" + config.ApiVersion() + "/time/utc",
				Func:    controllers.GetUtcTime,
				Methods: []string{"GET"},
			},
		},
	})
	defer testServer.Close()

	exitCode := m.Run()

	os.Exit(exitCode)
}
