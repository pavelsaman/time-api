package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/pavelsaman/time-api/api"
	"github.com/pavelsaman/time-api/config"
)

func TestApiGetEpochTime(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/"+config.ApiVersion()+"/time/{epochType}", api.GetEpochTime).Methods("GET")

	ts := httptest.NewServer(router)
	defer ts.Close()

	epochTypes := []string{
		"epoch",
		"epochmilli",
		"epochmicro",
		"epochnano",
	}
	for _, epochType := range epochTypes {
		// create and send request
		req, err := http.NewRequest("GET", ts.URL+"/"+config.ApiVersion()+"/time/"+epochType, nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer resp.Body.Close()

		// Verify the response
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", resp.Status)
		}

		var data map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			t.Errorf("error decoding the response body")
		}

		if data["type"] != epochType {
			t.Errorf("no value %v in type key in response body", epochType)
		}
		if data["epoch"] == nil {
			t.Error("epoch key not present in the response body")
		}
	}
}
