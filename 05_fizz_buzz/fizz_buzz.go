package main

import (
	"flag"
	"fmt"
)

func main() {
	var numberPtr = flag.Int("number", 100, "Max value")
	flag.Parse()

	for i := 1; i <= *numberPtr; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		fmt.Println(i)
	}
}
