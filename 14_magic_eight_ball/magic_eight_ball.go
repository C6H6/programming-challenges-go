package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ask()
	fmt.Println(getAnswer())
}

func ask() {
	fmt.Print("Write, say or think your question: ")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

func getAnswer() string {
	answers := []string{"It is certain.", "It is decidedly so.", "Without a doubt.", "Yes - definitely.",
		"You may rely on it.", "As I see it, yes.", "Most likely.", "Outlook good.", "Yes.", "Signs point to yes.",
		"Reply hazy, try again.", "Ask again later.", "Better not tell you now.", "Cannot predict now.",
		"Concentrate and ask again.", "Don't count on it.", "My reply is no.", "My sources say no.",
		"Outlook not so good.", "Very doubtful."}

	return fmt.Sprint(answers[rand.Intn(20)])
}
