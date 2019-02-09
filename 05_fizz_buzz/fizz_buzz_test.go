package main

import (
	"testing"
)

func TestValues(t *testing.T) {

	result := getValue(1)

	if result != "1" {
		t.Errorf("Invalid value: 1 -> %s", result)
	}

	result = getValue(3)

	if result != "Fizz" {
		t.Errorf("Invalid value: 3 -> %s", result)
	}

	result = getValue(5)

	if result != "Buzz" {
		t.Errorf("Invalid value: 5 -> %s", result)
	}
	result = getValue(15)

	if result != "Fizz Buzz" {
		t.Errorf("Invalid value: 15 -> %s", result)
	}
}
