package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ernestas-poskus/syllables"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Inspired by https://www.poem-generator.org.uk/haiku/
func main() {
	timeNoun := getWord("A time or year, time of day (e.g. summer, night)")
	timeAdj := getWord("An adjective to describe that time")
	noun1 := getWord("First singular noun")
	noun2 := getWord("Second singular noun")
	nounAdj := getWord("Noun adjective")
	verb := getWord("Verb (e.g. sings, runs)")

	haiku := getHaiku(timeNoun, timeAdj, noun1, noun2, nounAdj, verb)

	fmt.Printf("\n%s\n", haiku)
}

func getWord(text string) []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text + ": ")
	result, _ := reader.ReadString('\n')
	return []byte(result[:len(result)-1])
}

func getHaiku(timeNoun []byte, timeAdj []byte, noun1 []byte, noun2 []byte, nounAdj []byte, verb []byte) string {
	var result bytes.Buffer

	timeSyllables := syllables.CountSyllables(timeNoun)
	timeAdjSyllables := syllables.CountSyllables(timeAdj)
	nounAdjSyllables := syllables.CountSyllables(nounAdj)
	verbSyllables := syllables.CountSyllables(verb)

	// FIRST LINE
	firstLineSyllables := timeSyllables + timeAdjSyllables

	if firstLineSyllables == 5 {
		result.Write(timeAdj)
		result.WriteByte(' ')
		result.Write(timeNoun)
	} else if firstLineSyllables > 5 {
		if timeSyllables <= timeAdjSyllables && timeSyllables < 5 {
			missingSyllables := 5 - timeSyllables
			result.Write(getAdjectives(missingSyllables))
			result.WriteByte(' ')
			result.Write(timeNoun)
		} else if timeAdjSyllables < timeSyllables && timeAdjSyllables < 5 {
			missingSyllables := 5 - timeAdjSyllables
			result.Write(timeAdj)
			result.WriteByte(' ')
			result.Write(getTime(missingSyllables))
		} else {
			missingSyllables := rand.Intn(3) + 1
			result.Write(getAdjectives(missingSyllables))
			result.WriteByte(' ')
			result.Write(getAdjectives(5 - missingSyllables))
		}
	} else {
		if firstLineSyllables == 4 {
			result.Write(timeAdj)
			result.WriteByte(' ')
			result.Write(timeAdj)
			result.Write([]byte("time"))
		}  else {
			result.Write(getAdjectives(5 - timeSyllables))
			result.WriteByte(' ')
			result.Write(timeNoun)
		}
	}
	// END OF FIRST LINE

	result.WriteByte('\n')

	// SECOND LINE
	secondLineNoun := &noun1
	thirdLineNoun := &noun2
	if len(noun2) <= len(noun1) {
		secondLineNoun = &noun2
		thirdLineNoun = &noun1
	}

	secondLineNounSyllables := syllables.CountSyllables(*secondLineNoun)
	missingSyllables := 7 - 1 - nounAdjSyllables - secondLineNounSyllables - verbSyllables

	if missingSyllables == 0 {
		result.Write(getArticle(nounAdj))
		result.WriteByte(' ')
		result.Write(nounAdj)
		result.WriteByte(' ')
		result.Write(*secondLineNoun)
		result.WriteByte(' ')
		result.Write(verb)
	} else if missingSyllables > 0 {
		result.Write(getArticle(nounAdj))
		result.WriteByte(' ')
		result.Write(nounAdj)
		result.WriteByte(',')
		result.WriteByte(' ')
		result.Write(getAdjectives(missingSyllables))
		result.WriteByte(' ')
		result.Write(*secondLineNoun)
		result.WriteByte(' ')
		result.Write(verb)
	} else {
		missingSyllables := 7 - 1 - secondLineNounSyllables - verbSyllables
		result.Write(getArticle(*secondLineNoun))
		result.WriteByte(' ')
		result.Write(getAdjectives(missingSyllables))
		result.WriteByte(' ')
		result.Write(verb)
	}
	// END OF SECOND LINE

	result.WriteByte('\n')

	// THIRD LINE

	missingSyllables = 5 - syllables.CountSyllables(*thirdLineNoun)
	result.Write(getTerm(missingSyllables))
	result.WriteByte(' ')
	result.Write(*thirdLineNoun)
	// END OF THIRD LINE

	return result.String()
}

func getAdjectives(syllables int) []byte {
	set := map[int][]byte{
		0: []byte(""),
		1: []byte("good"),
		2: []byte("perfect"),
		3: []byte("interesting"),
		4: []byte("beautiful"),
	}

	return set[syllables]
}

func getTime(syllables int) []byte {
	set := map[int][]byte{
		1: []byte("night"),
		2: []byte("summer"),
		3: []byte("century"),
		4: []byte("nanosecond"),
	}

	return set[syllables]
}

func getTerm(syllables int) []byte {
	set := map[int][]byte{
		1: []byte("with"),
		2: []byte("next to"),
		3: []byte("watching the"),
		4: []byte("enjoying the"),
	}

	return set[syllables]
}

func getArticle(word []byte) []byte {
	x := word[0]
	for _, c := range []byte("aeiou") {
		if x == c {
			return []byte("an")
		}
	}
	return []byte("a")
}
