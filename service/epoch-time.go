package service

import (
	"fmt"
	"time"
)

func GetEpochTime(epochType ...string) (int64, error) {
	if len(epochType) == 0 {
		return 0, fmt.Errorf("missing epoch time")
	}

	et := epochType[0]
	switch et {
	case "epochmilli":
		return time.Now().UnixMilli(), nil
	case "epochmicro":
		return time.Now().UnixMicro(), nil
	case "epochnano":
		return time.Now().UnixNano(), nil
	case "epoch":
		return time.Now().Unix(), nil
	default:
		return 0, fmt.Errorf("unknown epoch time")
	}
}
