package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetSubNameReturnsString(t *testing.T) {
	result := getSubName([]string{"aaa"})

	if reflect.TypeOf(result).String() != "string" {
		t.Error("Result is not a string")
	}
}

func TestGetSubNameMakeUppercase(t *testing.T) {
	result := getSubName([]string{"aaa"})

	if result[0] != 'A' {
		t.Error("Result doesn't start with uppercase")
	}
}

func TestGetNameAddsSpace(t *testing.T) {
	result := getName()

	if !strings.Contains(result, " ") {
		t.Error("Result doesn't contain space")
	}
}
