package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetPasswordWithGivenLength(t *testing.T) {
	password := getPassword(12)

	if len(password) != 12 {
		t.Error("Invalid length")
	}
}

func TestGetGroupsReturnsExpectedGroups(t *testing.T) {
	result := getGroups(12, 3)
	expected := []int{4, 4, 4}

	fmt.Println(result)

	if !reflect.DeepEqual(result, expected) {
		t.Error("Invalid groups")
	}
}
