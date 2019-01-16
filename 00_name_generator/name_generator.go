package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var firstNameSyllables = []string{"mon", "fay", "mag", "shi", "zag", "blarg", "rash", "izen"}
	var lastNameSyllables = []string{"malo", "zak", "abo", "wonk", "fok"}

	var firstName = getName(firstNameSyllables)
	var lastName = getName(lastNameSyllables)

	fmt.Printf("%s %s", firstName, lastName)
}

func getName(nameSet []string) string {
	rand.Seed(time.Now().UnixNano())
	var syllables = rand.Intn(2) + 2

	var buffer bytes.Buffer
	for i := 0; i < syllables; i++ {
		buffer.WriteString(nameSet[rand.Intn(len(nameSet))])
	}

	return strings.Title(buffer.String())
}
