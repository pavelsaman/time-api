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

var endpointPaths map[string]string = map[string]string{
	"version":               "/" + config.ApiVersion() + "/version",
	"describe":              "/" + config.ApiVersion() + "/describe",
	"unix.epochtype":        "/" + config.ApiVersion() + "/unix/{epochType}",
	"unix.epoch.epochvalue": "/" + config.ApiVersion() + "/unix/epoch/{epochValue}",
	"time.utc":              "/" + config.ApiVersion() + "/time/utc",
}

var Endpoints []Endpoint = []Endpoint{
	{
		Url:     endpointPaths["version"],
		Func:    GetVersion,
		Methods: []string{"GET"},
	},
	{
		Url:     endpointPaths["describe"],
		Func:    Describe,
		Methods: []string{"GET"},
	},
	// epoch
	{
		Url:     endpointPaths["unix.epochtype"],
		Func:    GetEpochTime,
		Methods: []string{"GET"},
	},
	{
		Url:     endpointPaths["unix.epoch.epochvalue"],
		Func:    GetEpochToUtc,
		Methods: []string{"GET"},
	},
	// utc
	{
		Url:     endpointPaths["time.utc"],
		Func:    GetUtcTime,
		Methods: []string{"GET"},
	},
}

func GetAllControllerPaths() []string {
	var paths []string
	for _, v := range endpointPaths {
		paths = append(paths, v)
	}
	return paths
}
