package types

type EpochAndUtcTime struct {
	Type  string
	Epoch int64
	Utc   string
}

type EpochToUtcTime struct {
	Type string
	Utc  string
}
