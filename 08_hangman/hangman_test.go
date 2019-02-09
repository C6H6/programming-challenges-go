package main

import "testing"

func TestProcessReturnsProperValues(t *testing.T) {
	masked := []byte{'*', '*', '*', '*'}
	masked, success := process("test", masked, "t")

	if !success {
		t.Error("Invalid success value")
	}

	if string(masked) != "t**t" {
		t.Error("Invalid masked value")
	}

	masked, success = process("test", masked, "x")

	if success {
		t.Error("Invalid success value")
	}

	if string(masked) != "t**t" {
		t.Error("Invalid masked value")
	}
}
