package main

import (
	"reflect"
	"testing"
)

func TestGetAnswerReturnsString(t *testing.T) {
	result := getAnswer()

	if reflect.TypeOf(result).String() != "string" {
		t.Error("Invalid type of returned value")
	}
}
