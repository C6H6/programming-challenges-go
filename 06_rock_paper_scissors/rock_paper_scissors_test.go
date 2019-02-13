package main

import (
	"reflect"
	"testing"
)

func TestGetResultReturnsProperString(t *testing.T) {
	result := getResult(0, 0)

	if result != "Draw! (Rock vs. Rock)" {
		t.Error("Invalid result")
	}

	result = getResult(1, 0)

	if result != "Win! (Paper vs. Rock)" {
		t.Error("Invalid result")
	}

	result = getResult(2, 0)

	if result != "Loss! (Scissors vs. Rock)" {
		t.Error("Invalid result")
	}
}

func TestComputerChoiceIsInt(t *testing.T)  {
	number := getComputerChoice()

	if reflect.TypeOf(number).String() != "int" {
		t.Error("Invalid computer choice")
	}
}
