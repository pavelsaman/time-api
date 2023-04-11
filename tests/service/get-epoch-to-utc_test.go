package service_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pavelsaman/time-api/api/services"
	"github.com/pavelsaman/time-api/api/types"
)

func TestServiceGetUtcTime(t *testing.T) {
	epochToUtc := map[int64]types.EpochToUtcTime{
		1681038302: {
			Type: "epoch",
			Utc:  fmt.Sprint(time.Unix(1681038302, 0).UTC()),
		},
		1681038781158: {
			Type: "epochmilli",
			Utc:  fmt.Sprint(time.UnixMilli(1681038781158).UTC()),
		},
		1681038860897835: {
			Type: "epochmicro",
			Utc:  fmt.Sprint(time.UnixMicro(1681038860897835).UTC()),
		},
		1681038887043863923: {
			Type: "epochnano",
			Utc:  fmt.Sprint(time.Unix(0, 1681038887043863923).UTC()),
		},
	}

	for epoch, utc := range epochToUtc {
		actual, err := services.GetEpochToUtc(fmt.Sprint(epoch))
		if err != nil {
			t.Errorf("Epoch %v cannot be converted to utc", epoch)
		}
		expected := utc
		if actual.Utc != expected.Utc {
			t.Errorf("Actual: %v, expected: %v", actual.Utc, expected.Utc)
		}
		if actual.Type != expected.Type {
			t.Errorf("Actual: %v, expected: %v", actual.Type, expected.Type)
		}
	}
}

func TestServiceCorrectTimeForEpochZero(t *testing.T) {
	actual, err := services.GetEpochToUtc("0")
	if err != nil {
		t.Errorf("Epoch %v cannot be converted to utc", 0)
	}

	expected := fmt.Sprint(time.Unix(0, 0).UTC())
	if actual.Utc != expected {
		t.Errorf("Actual: %v, expected: %v", actual.Utc, expected)
	}
	if actual.Type != "epoch" {
		t.Error("The result epoch type is not \"epoch\"")
	}
}

func TestServiceNegativeEpoch(t *testing.T) {
	actual, err := services.GetEpochToUtc("-12")
	if err != nil {
		t.Errorf("Epoch %v cannot be converted to utc", 0)
	}

	expected := fmt.Sprint(time.Unix(-12, 0).UTC())
	if actual.Utc != expected {
		t.Errorf("Actual: %v, expected: %v", actual.Utc, expected)
	}
	if actual.Type != "epoch" {
		t.Error("The result epoch type is not \"epoch\"")
	}
}

func TestServiceOverflowSignedInt64Bit(t *testing.T) {
	actual, err := services.GetEpochToUtc("9223372036854775808")
	if err == nil {
		t.Error("No error when one expected")
	}
	if err.Error() != "cannot convert epoch" {
		t.Error("Not expected error")
	}
	if actual != nil {
		t.Error("Result given when none expected")
	}
}

func TestServiceTextualInput(t *testing.T) {
	actual, err := services.GetEpochToUtc("a")
	if err == nil {
		t.Error("No error when one expected")
	}
	if err.Error() != "cannot convert epoch" {
		t.Error("Not expected error")
	}
	if actual != nil {
		t.Error("Result given when none expected")
	}
}
