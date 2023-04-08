package service_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pavelsaman/time-api/service"
)

func TestServiceGetEpochTime(t *testing.T) {
	actual, err := service.GetEpochTime("epoch")
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
	actual, err := service.GetEpochTime("epochmilli")
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
	actual, err := service.GetEpochTime("epochmicro")
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
	actual, err := service.GetEpochTime("epochnano")
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
	_, err := service.GetEpochTime()
	if err == nil {
		t.Error("expected error because non-existent epoch time")
	}
}

func TestServiceGetEpochTimeCanUseCamelCase(t *testing.T) {
	_, err := service.GetEpochTime("epochMilli")
	if err != nil {
		t.Errorf("cannot use camelCase in epoch type")
	}
}
