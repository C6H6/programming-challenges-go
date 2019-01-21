package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func main() {
	var phrase = readPhrase()
	var masked []byte

	for i := 0; i < len(phrase); i++ {
		masked = append(masked, '*')
	}
	printMasked(masked)

	var tries = 8

	for tries > 0 {
		var shot = readChar()
		var success = false
		masked, success = process(phrase, masked, shot)

		if !success {
			tries--
		}

		printMasked(masked)

		fmt.Printf("Tries left: %d\n", tries)

		if string(masked) == phrase {
			break
		}
	}

	if tries > 0 {
		fmt.Println("You win!")
	} else {
		fmt.Printf("You lose")
	}

	fmt.Printf("Phrase: %s\n", phrase)
}

func readPhrase() string {
	for {
		fmt.Print("Enter secret phrase: ")
		password, err := terminal.ReadPassword(0)

		if err == nil {
			fmt.Println()
			return string(password)
		}
	}
}

func readChar() string {
	fmt.Print("Guess: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return string([]byte(input)[0])
}

func process(phrase string, masked []byte, shot string) ([]byte, bool) {
	var success = false

	for pos, char := range phrase {
		if string(char) == shot {
			success = true
			masked[pos] = byte(char)
		}
	}

	return masked, success
}

func printMasked(slice []byte) {
	fmt.Println(string(slice))
}
