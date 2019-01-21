package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 0-Rock; 1-Paper; 2-Scissors
	// [user][computer]
	var results = [][]string{
		{"Draw", "Loss", "Win"},
		{"Win", "Draw", "Loss"},
		{"Loss", "Win", "Draw"},
	}

	var options = []string{"Rock", "Paper", "Scissors"}

	var user = getChoice()
	rand.Seed(time.Now().UnixNano())
	var computer = rand.Intn(3)

	fmt.Printf("%s! (%s vs. %s)\n", results[user][computer], options[user], options[computer])
}

func getChoice() int {
	fmt.Println("1. Rock\n2. Paper\n3. Scissors")
	var i int
	for {
		fmt.Printf("Choice: ")
		_, err := fmt.Scanf("%d", &i)
		if err == nil && i >= 1 && i <= 3 {
			return i - 1
		}
	}
}
