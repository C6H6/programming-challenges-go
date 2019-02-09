package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var numberPtr = flag.Int("number", 100, "Max value")
	flag.Parse()

	for i := 1; i <= *numberPtr; i++ {
		fmt.Println(getValue(i))
	}
}

func getValue(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return "Fizz Buzz"
	}
	if i%5 == 0 {
		return "Buzz"
	}
	if i%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(i)
}
