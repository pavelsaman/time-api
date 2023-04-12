package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/pavelsaman/time-api/api/types"
)

func TestServiceGetEpochTime(t *testing.T) {
	actual, err := GetEpochTime("epoch")
	expected := len(fmt.Sprint(time.Now().UTC().Unix()))
	if len(fmt.Sprint(actual.Epoch)) != expected {
		t.Errorf("Actual: %v, expected %v", len(fmt.Sprint(actual.Epoch)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epoch type")
	}
	if fmt.Sprint(time.Unix(actual.Epoch, 0).UTC()) != actual.Utc {
		t.Errorf("Epoch %v is not equal to utc %v", actual.Epoch, actual.Utc)
	}
}

func TestServiceGetEpochTimeMilli(t *testing.T) {
	actual, err := GetEpochTime("epochmilli")
	expected := len(fmt.Sprint(time.Now().UnixMilli()))
	if len(fmt.Sprint(actual.Epoch)) != expected {
		t.Errorf("Milli - Actual: %v, expected %v", len(fmt.Sprint(actual.Epoch)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epochmilli type")
	}
	if fmt.Sprint(time.UnixMilli(actual.Epoch).UTC()) != actual.Utc {
		t.Errorf("Epoch %v is not equal to utc %v", actual.Epoch, actual.Utc)
	}
}

func TestServiceGetEpochTimeMicro(t *testing.T) {
	actual, err := GetEpochTime("epochmicro")
	expected := len(fmt.Sprint(time.Now().UnixMicro()))
	if len(fmt.Sprint(actual.Epoch)) != expected {
		t.Errorf("Micro - Actual: %v, expected %v", len(fmt.Sprint(actual.Epoch)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epochmicro type")
	}
	if fmt.Sprint(time.UnixMicro(actual.Epoch).UTC()) != actual.Utc {
		t.Errorf("Epoch %v is not equal to utc %v", actual.Epoch, actual.Utc)
	}
}

func TestServiceGetEpochTimeNano(t *testing.T) {
	actual, err := GetEpochTime("epochnano")
	expected := len(fmt.Sprint(time.Now().UnixNano()))
	if len(fmt.Sprint(actual.Epoch)) != expected {
		t.Errorf("Nano - Actual: %v, expected %v", len(fmt.Sprint(actual.Epoch)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epochnano type")
	}
	if fmt.Sprint(time.Unix(0, actual.Epoch).UTC()) != actual.Utc {
		t.Errorf("Epoch %v is not equal to utc %v", actual.Epoch, actual.Utc)
	}
}

func TestServiceGetEpochNone(t *testing.T) {
	_, err := GetEpochTime()
	if err == nil {
		t.Error("expected error because non-existent epoch time")
	}
}

func TestServiceGetEpochTimeCanUseCamelCase(t *testing.T) {
	_, err := GetEpochTime("epochMilli")
	if err != nil {
		t.Errorf("cannot use camelCase in epoch type")
	}
}

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
		actual, err := GetEpochToUtc(fmt.Sprint(epoch))
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
	actual, err := GetEpochToUtc("0")
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
	actual, err := GetEpochToUtc("-12")
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
	actual, err := GetEpochToUtc("9223372036854775808")
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
	actual, err := GetEpochToUtc("a")
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
