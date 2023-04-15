package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pavelsaman/time-api/api/types"
)

var (
	errUnknownEpochType   = errors.New("epoch type does not exist")
	errUnknownEpoch       = errors.New("unknown epoch")
	errCannotConvertEpoch = errors.New("cannot convert epoch")
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
		return nil, errUnknownEpochType
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
		return nil, errCannotConvertEpoch
	}

	var t time.Time
	var et string
	switch {
	case len(epochValue) < epochMilli:
		t = time.Unix(epochTime, 0)
		et = "epoch"
	case len(epochValue) < epochMicro:
		t = time.UnixMilli(epochTime)
		et = "epochmilli"
	case len(epochValue) < epochNano:
		t = time.UnixMicro(epochTime)
		et = "epochmicro"
	case len(epochValue) == epochNano:
		t = time.Unix(0, epochTime)
		et = "epochnano"
	default:
		return nil, errUnknownEpoch
	}

	return &types.EpochToUtcTime{
		Utc:  fmt.Sprint(t.UTC()),
		Type: et,
	}, nil
}
