package main

import (
	"testing"
)

func TestGetResultReturnsProperValue(t *testing.T) {
	result := getResult()

	if result != "head" && result != "tail" {
		t.Error("Invalid result")
	}
}
