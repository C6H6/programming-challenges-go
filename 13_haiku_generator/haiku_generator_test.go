package main

import (
	"github.com/ernestas-poskus/syllables"
	"strings"
	"testing"
)

func TestGetHaikuReturnsHaikuWithProperLengths(t *testing.T) {
	result := getHaiku([]byte("summer"), []byte("hot"), []byte("pencil"), []byte("cup"), []byte("red"), []byte("runs"))

	array := strings.Split(result, "\n")

	if len(array) != 3 {
		t.Errorf("Wrong number of rows: %d", len(array))
	}

	// syllables.CountSyllables works with accuracy +/- 1 syllable
	if
	syllables.CountSyllables([]byte(array[0])) < 4 ||
		syllables.CountSyllables([]byte(array[0])) > 6 ||
		syllables.CountSyllables([]byte(array[1])) < 6 ||
		syllables.CountSyllables([]byte(array[1])) > 8 ||
		syllables.CountSyllables([]byte(array[2])) < 4 ||
		syllables.CountSyllables([]byte(array[2])) > 6 {
		t.Error("Wrong number of syllables in rows")
	}
}
