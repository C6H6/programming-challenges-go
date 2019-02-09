package main

import "testing"

func TestGetNewNumberReturnsProperNumber(t *testing.T) {
	if getNewNumber(1) != 4 {
		t.Error("Invalid new number")
	}

	if getNewNumber(10) != 5 {
		t.Error("Invalid new number")
	}

	if getNewNumber(15) != 46 {
		t.Error("Invalid new number")
	}
}
