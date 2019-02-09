package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	user := getUserChoice()
	computer := getComputerChoice()

	fmt.Println(getResult(user, computer))
}

func getUserChoice() int {
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

func getComputerChoice() int {
	return rand.Intn(3)
}

func getResult(user int, computer int) string {
	// 0-Rock; 1-Paper; 2-Scissors
	// [user][computer]
	var results = [][]string{
		{"Draw", "Loss", "Win"},
		{"Win", "Draw", "Loss"},
		{"Loss", "Win", "Draw"},
	}

	var options = []string{"Rock", "Paper", "Scissors"}

	return fmt.Sprintf("%s! (%s vs. %s)", results[user][computer], options[user], options[computer])
}
