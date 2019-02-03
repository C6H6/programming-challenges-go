package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "START_NUMBER")
		return
	}

	param := os.Args[1]
	number, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		fmt.Println("Invalid number")
		return
	}
	fmt.Printf("Start: %d\n", number)

	steps := 0

	for number != 1 {
		steps++
		number = getNewNumber(number)
		fmt.Println(number)
	}

	fmt.Printf("Steps: %d\n", steps)
}

func getNewNumber(number uint64) uint64 {
	if number%2 == 0 {
		return number / 2
	}

	return 3*number + 1
}
