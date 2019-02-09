package main

import (
	"testing"
)

func TestCrypt(t *testing.T) {
	result := crypt("test")

	if result != "4%34" {
		t.Errorf("Invalid crypt output: test -> %s", result)
	}
}
