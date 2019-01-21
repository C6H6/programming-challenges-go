package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var persons = []string{"I", "You", "He", "She", "It", "We", "They"}
	var negations = []string{"don't", "doesn't"}

	var verbs = map[string]map[string][]string{
		"play": {"nouns": {"football", "guitar", "games"}, "variations": {"play", "plays"}},
		"like": {"nouns": {"ice-cream", "running", "motorcycles"}, "variations": {"like", "likes"}},
		"bake": {"nouns": {"cake", "pie", "cookies"}, "variations": {"bake", "bakes"}},
		"read": {"nouns": {"book", "newspaper", "letter"}, "variations": {"read", "reads"}},
	}

	var person = rand.Intn(7)
	var negation = rand.Intn(2) == 0
	var thirdPerson = isThirdPerson(person)

	var variation = 1
	if negation || (!negation && !thirdPerson) {
		variation = 0
	}

	var sentence = make([]string, 0, 4)
	sentence = append(sentence, persons[person])
	var verb = getRandomKey(verbs)

	if negation {
		if thirdPerson {
			variation = 1 - variation
		}
		sentence = append(sentence, negations[int(variation)])
		sentence = append(sentence, verbs[verb]["variations"][0])
	} else {
		sentence = append(sentence, verbs[verb]["variations"][variation])
	}

	sentence = append(sentence, verbs[verb]["nouns"][len(verbs[verb]["nouns"])-1])

	fmt.Println(strings.Join(sentence, " "))
}

func isThirdPerson(person int) bool {
	for i := 2; i <= 4; i++ {
		if person == i {
			return true
		}
	}
	return false
}

func getRandomKey(verbs map[string]map[string][]string) string {
	i := rand.Intn(len(verbs))
	var k string
	for k = range verbs {
		if i == 0 {
			break
		}
		i--
	}
	return k
}
