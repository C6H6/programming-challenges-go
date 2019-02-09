package main

import (
	"reflect"
	"testing"
	"time"
)

func TestResultIsFloat(t *testing.T) {
	date, _ := time.Parse("02-01-2006", "11-11-2011")
	result := getResult(date)

	if reflect.TypeOf(result).String() != "float64" {
		t.Errorf("Invalid result")
	}
}
