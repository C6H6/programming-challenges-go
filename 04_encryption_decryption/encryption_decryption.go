package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "TEXT_TO_ENCRYPT/DECRYPT")
		return
	}

	text := os.Args[1]

	fmt.Println(crypt(text))
}

func crypt(text string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(text); i++ {
		var newChar = string((text[i] + 64) % 128)
		buffer.WriteString(newChar)
	}

	return buffer.String()
}
