package main

import (
	"testing"
)

func TestCountReturnsRightNumberOfWords(t *testing.T) {
	result := countWordsInString("one two three")

	if result != 3 {
		t.Errorf("Wrong number of words")
	}
}

func TestCountReturns0OnEmptyString(t *testing.T) {
	result := countWordsInString("")

	if result != 0 {
		t.Errorf("Wrong number of words")
	}
}
