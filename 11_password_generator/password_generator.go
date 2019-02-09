package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var lengthPtr = flag.Int("length", 8, "Password length")
	flag.Parse()

	result := getPassword(*lengthPtr)

	fmt.Println(string(result))
}

func getPassword(length int) []byte {
	var charsets = []string{"abcdefghijklmnopqrstuvwxyz", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "0123456789", "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"}
	var groups = getGroups(length, len(charsets))

	var result = make([]byte, length, length)

	var pos = 0
	for index, size := range groups {
		for i := 0; i < size; i++ {
			result[pos] = getRandomByteFromString(charsets[index])
			pos++
		}
	}

	return shuffle(result)
}

func getGroups(length int, groups int) []int {
	var result = make([]int, groups, groups)
	for i := 0; i < len(result); i++ {
		result[i] = (length + groups - i - 1) / (groups - i)
		length -= result[i]
	}

	return result
}

func getRandomByteFromString(str string) byte {
	return str[rand.Intn(len(str))]
}

func shuffle(arr []byte) []byte {
	for i := range arr {
		n := rand.Intn(i + 1)
		arr[i], arr[n] = arr[n], arr[i]
	}

	return arr
}
