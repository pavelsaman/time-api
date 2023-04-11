package controllers

import (
	"net/http"

	"github.com/pavelsaman/time-api/config"
)

type Endpoint struct {
	Url     string
	Func    func(http.ResponseWriter, *http.Request)
	Methods []string
}

var Endpoints []Endpoint = []Endpoint{
	{
		Url:     "/" + config.ApiVersion() + "/version",
		Func:    GetVersion,
		Methods: []string{"GET"},
	},
	// epoch
	{
		Url:     "/" + config.ApiVersion() + "/unix/{epochType}",
		Func:    GetEpochTime,
		Methods: []string{"GET"},
	},
	{
		Url:     "/" + config.ApiVersion() + "/unix/epoch/{epochValue}",
		Func:    GetEpochToUtc,
		Methods: []string{"GET"},
	},
	// utc
	{
		Url:     "/" + config.ApiVersion() + "/time/utc",
		Func:    GetUtcTime,
		Methods: []string{"GET"},
	},
}
