package types

import "net/http"

type EpochAndUtcTimeResponse struct {
	Type  string `json:"type"`
	Epoch string `json:"epoch"`
	Utc   string `json:"utc"`
}

type EpochToUtcTimeResponse struct {
	Utc       string `json:"utc"`
	EpochType string `json:"epochType"`
}

type UtcTimeResponse struct {
	Type string `json:"type"`
	Utc  string `json:"utc"`
}

type ApiErrorResponse struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type Handler struct {
	Url     string
	Func    func(http.ResponseWriter, *http.Request)
	Methods []string
}

type Handlers struct {
	Handlers []*Handler
}
