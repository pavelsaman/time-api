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
}

func TestServiceGetEpochTimeMilli(t *testing.T) {
	actual, err := service.GetEpochTime("epochmilli")
	expected := len(fmt.Sprint(time.Now().UnixMilli()))
	if len(fmt.Sprint(actual)) != expected {
		t.Errorf("Milli - Actual: %v, expected %v", len(fmt.Sprint(actual)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epochmilli type")
	}
}

func TestServiceGetEpochTimeMicro(t *testing.T) {
	actual, err := service.GetEpochTime("epochmilli")
	expected := len(fmt.Sprint(time.Now().UnixMilli()))
	if len(fmt.Sprint(actual)) != expected {
		t.Errorf("Micro - Actual: %v, expected %v", len(fmt.Sprint(actual)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epochmilli type")
	}
}

func TestServiceGetEpochTimeNano(t *testing.T) {
	actual, err := service.GetEpochTime("epochmilli")
	expected := len(fmt.Sprint(time.Now().UnixMilli()))
	if len(fmt.Sprint(actual)) != expected {
		t.Errorf("Nano - Actual: %v, expected %v", len(fmt.Sprint(actual)), expected)
	}
	if err != nil {
		t.Errorf("Error present for epochmilli type")
	}
}

func TestServiceGetEpochNone(t *testing.T) {
	_, err := service.GetEpochTime()
	if err == nil {
		t.Error("expected error because non-existent epoch time")
	}
}
