package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var numberPtr = flag.Int("number", 1, "Number of draws")
	flag.Parse()

	var set = []string{"head", "tail"}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < *numberPtr; i++ {
		fmt.Println(set[rand.Intn(2)])
	}
}
