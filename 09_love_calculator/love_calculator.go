package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "FIRST_NAME", "SECOND_NAME")
		return
	}

	var buffer bytes.Buffer

	buffer.WriteString(strings.ToLower(os.Args[1]))
	buffer.WriteString(strings.ToLower(os.Args[2]))

	var sum = 0

	for char := range buffer.String() {
		sum += int(math.Pow(float64(char), float64(4)))
	}

	fmt.Printf("%d%%\n", sum%101)
}
