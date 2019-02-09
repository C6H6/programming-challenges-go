package main

import (
	"reflect"
	"testing"
)

func TestResultIsInt(t *testing.T) {
	result := getResult("aaa", "bbb")

	if reflect.TypeOf(result).String() != "int" {
		t.Error("Result is not int")
	}
}

func TestResultInRange(t *testing.T) {
	result := getResult("aaa", "bbb")

	if result > 100 || result < 0 {
		t.Error("Result out of range")
	}
}

func TestReturnsSameResultOnDifferentLetterCase(t *testing.T) {
	result1 := getResult("aaa", "bbb")
	result2 := getResult("AAA", "BBB")

	if result1 != result2 {
		t.Error("Different results on the same letters case")
	}
}
