package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var numberPtr = flag.Int("number", 1, "Number of draws")
	flag.Parse()

	for i := 0; i < *numberPtr; i++ {
		fmt.Println(getResult())
	}
}

func getResult() string {
	set := []string{"head", "tail"}
	return set[rand.Intn(2)]
}
