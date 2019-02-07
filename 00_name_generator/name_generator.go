package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println(getName())
}

func getName() string {
	var firstNameSyllables = []string{"mon", "fay", "mag", "shi", "zag", "blarg", "rash", "izen"}
	var lastNameSyllables = []string{"malo", "zak", "abo", "wonk", "fok"}

	var buffer bytes.Buffer

	buffer.WriteString(getSubName(firstNameSyllables))
	buffer.WriteString(" ")
	buffer.WriteString(getSubName(lastNameSyllables))

	return buffer.String()
}

func getSubName(nameSet []string) string {
	rand.Seed(time.Now().UnixNano())
	var syllables = rand.Intn(2) + 2

	var buffer bytes.Buffer
	for i := 0; i < syllables; i++ {
		buffer.WriteString(nameSet[rand.Intn(len(nameSet))])
	}

	return strings.Title(buffer.String())
}
