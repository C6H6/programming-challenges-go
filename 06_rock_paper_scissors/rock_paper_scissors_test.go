package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetResultReturnsProperString(t *testing.T) {
	result := getResult(0, 0)

	fmt.Println(result)

	if result != "Draw! (Rock vs. Rock)" {
		t.Error("Invalid result")
	}

	result = getResult(1, 0)

	fmt.Println(result)

	if result != "Win! (Paper vs. Rock)" {
		t.Error("Invalid result")
	}

	result = getResult(2, 0)

	fmt.Println(result)

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
