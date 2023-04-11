package api_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/pavelsaman/time-api/config"
)

func TestApiGetUtcTime(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL+"/"+config.ApiVersion()+"/time/utc", nil)
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
		t.Errorf("Expected %v; got %v", http.StatusOK, resp.Status)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		t.Errorf("error decoding the response body")
	}

	if data["type"] != "utc" {
		t.Error("Expected type \"utc\"")
	}
	if data["utc"] == nil {
		t.Error("no utc key in the response")
	}
}
