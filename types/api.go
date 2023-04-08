package types

type EpochAndUtcTimeResponse struct {
	Type  string `json:"type"`
	Epoch string `json:"epoch"`
	Utc   string `json:"utc"`
}

type EpochToUtcTimeResponse struct {
	Utc       string `json:"utc"`
	EpochType string `json:"epochType"`
}

type ApiErrorResponse struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
