package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pavelsaman/time-api/types"
)

func GetEpochTime(epochType ...string) (*types.EpochAndUtcTime, error) {
	if len(epochType) == 0 {
		return nil, fmt.Errorf("missing epoch time")
	}

	et := strings.ToLower(epochType[0])
	t := time.Now().UTC()
	switch et {
	case "epoch":
		utc, err := GetEpochToUtc(strconv.FormatInt(t.Unix(), 10))
		if err != nil {
			return nil, err
		}
		return &types.EpochAndUtcTime{
			Type:  et,
			Epoch: t.Unix(),
			Utc:   utc.Utc,
		}, nil
	case "epochmilli":
		utc, err := GetEpochToUtc(strconv.FormatInt(t.UnixMilli(), 10))
		if err != nil {
			return nil, err
		}
		return &types.EpochAndUtcTime{
			Type:  et,
			Epoch: t.UnixMilli(),
			Utc:   utc.Utc,
		}, nil
	case "epochmicro":
		utc, err := GetEpochToUtc(strconv.FormatInt(t.UnixMicro(), 10))
		if err != nil {
			return nil, err
		}
		return &types.EpochAndUtcTime{
			Type:  et,
			Epoch: t.UnixMicro(),
			Utc:   utc.Utc,
		}, nil
	case "epochnano":
		utc, err := GetEpochToUtc(strconv.FormatInt(t.UnixNano(), 10))
		if err != nil {
			return nil, err
		}
		return &types.EpochAndUtcTime{
			Type:  et,
			Epoch: t.UnixNano(),
			Utc:   utc.Utc,
		}, nil
	default:
		return nil, fmt.Errorf("epoch type \"%v\" does not exist", et)
	}
}

func GetEpochToUtc(epochValue string) (*types.EpochToUtcTime, error) {
	const (
		epoch      int = 10
		epochMilli int = 13
		epochMicro int = 16
		epochNano  int = 19
	)

	epochTime, err := strconv.ParseInt(epochValue, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot convert epoch")
	}

	var t time.Time
	var et string
	switch len(epochValue) {
	case epoch:
		t = time.Unix(epochTime, 0)
		et = "epoch"
	case epochMilli:
		t = time.UnixMilli(epochTime)
		et = "epochmilli"
	case epochMicro:
		t = time.UnixMicro(epochTime)
		et = "epochmicro"
	case epochNano:
		t = time.Unix(0, epochTime)
		et = "epochnano"
	default:
		return nil, fmt.Errorf("unknown epoch")
	}

	return &types.EpochToUtcTime{
		Utc:  fmt.Sprint(t.UTC()),
		Type: et,
	}, nil
}
