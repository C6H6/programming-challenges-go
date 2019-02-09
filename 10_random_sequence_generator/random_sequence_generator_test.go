package main

import (
	"reflect"
	"testing"
)

func TestGetSentenceReturnsString(t *testing.T) {
	result := getSentence()

	if reflect.TypeOf(result).String() != "string" {
		t.Error("Invalid type of returned value")
	}
}
