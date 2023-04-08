package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/pavelsaman/time-api/api"
	"github.com/pavelsaman/time-api/config"
	api_test_utils "github.com/pavelsaman/time-api/tests/api/utils"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	testServer = api_test_utils.StartTestServerAndRegisterHandlers(&api_test_utils.Handlers{
		Handlers: []*api_test_utils.Handler{
			{
				Url:     "/" + config.ApiVersion() + "/time/{epochType}",
				Func:    api.GetEpochTime,
				Methods: []string{"GET"},
			},
		},
	})
	defer testServer.Close()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestApiGetEpochTime(t *testing.T) {
	epochTypes := []string{
		"epoch",
		"Epoch",
		"epochmilli",
		"epochmicro",
		"epochMICRO",
		"epochnano",
		"epochNano",
	}
	for _, epochType := range epochTypes {
		req, err := http.NewRequest("GET", testServer.URL+"/"+config.ApiVersion()+"/time/"+epochType, nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer resp.Body.Close()

		// Checks
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", resp.Status)
		}

		var data map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			t.Errorf("error decoding the response body")
		}

		if data["type"] != strings.ToLower(epochType) {
			t.Errorf("no value %v in type key in response body", epochType)
		}
		if data["epoch"] == nil {
			t.Error("epoch key not present in the response body")
		}
	}
}

func TestApiGetEpochTimeBadRequest(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL+"/"+config.ApiVersion()+"/time/noEpoch", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("could not send request: %v", err)
	}
	defer resp.Body.Close()

	// Checks
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400; got %v", resp.Status)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		t.Errorf("error decoding the response body")
	}

	if int(data["errorCode"].(float64)) != http.StatusBadRequest {
		t.Errorf("Response body contains error %v, expected code 400", data["errorCode"])
	}
	if data["errorMessage"] == nil {
		t.Error("no errorMessage property in the response")
	}
}
