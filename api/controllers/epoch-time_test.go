package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/pavelsaman/time-api/config"
)

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
		req, err := http.NewRequest("GET", testServer.URL+"/"+config.ApiVersion()+"/unix/"+epochType, nil)
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
	req, err := http.NewRequest("GET", testServer.URL+"/"+config.ApiVersion()+"/unix/noEpoch", nil)
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
