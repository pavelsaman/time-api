package types

type EpochTimeResponse struct {
	Type  string `json:"type"`
	Epoch string `json:"epoch"`
	Utc   string `json:"utc"`
}

type EpochToUtcTimeResponse struct {
	Utc       string `json:"utc"`
	EpochType string `json:"epochType"`
}

type ApiError struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
