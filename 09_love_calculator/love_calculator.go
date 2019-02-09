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

	fmt.Printf("%d%%\n", getResult(os.Args[1], os.Args[2]))
}

func getResult(name1 string, name2 string) int {
	var buffer bytes.Buffer

	buffer.WriteString(strings.ToLower(name1))
	buffer.WriteString(strings.ToLower(name2))

	var sum = 0

	for char := range buffer.String() {
		sum += int(math.Pow(float64(char), float64(4)))
	}

	return sum%101
}
