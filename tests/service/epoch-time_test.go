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
	if len(fmt.Sprint(actual)) != expected {
		t.Errorf("Actual: %v, expected %v", len(fmt.Sprint(actual)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epoch type")
	}

	actual, err = service.GetEpochTime("epochmilli")
	expected = len(fmt.Sprint(time.Now().UnixMilli()))
	if len(fmt.Sprint(actual)) != expected {
		t.Errorf("Milli - Actual: %v, expected %v", len(fmt.Sprint(actual)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epochmilli type")
	}

	actual, err = service.GetEpochTime("epochmicro")
	expected = len(fmt.Sprint(time.Now().UnixMicro()))
	if len(fmt.Sprint(actual)) != expected {
		t.Errorf("Micro - Actual: %v, expected %v", len(fmt.Sprint(actual)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epochmicro type")
	}

	actual, err = service.GetEpochTime("epochnano")
	expected = len(fmt.Sprint(time.Now().UnixNano()))
	if len(fmt.Sprint(actual)) != expected {
		t.Errorf("Nano - Actual: %v, expected %v", len(fmt.Sprint(actual)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epochnano type")
	}

	_, err = service.GetEpochTime("nano")
	if err == nil {
		t.Errorf("Error present for nano type")
	}
}
