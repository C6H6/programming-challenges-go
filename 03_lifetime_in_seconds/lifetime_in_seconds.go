package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var birth = getDate()
	fmt.Printf("%d seconds", int64(getResult(birth)))
}

func getDate() time.Time {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Date of birth (DD-MM-YYYY): ")
		text, readErr := reader.ReadString('\n')
		date, timeErr := time.Parse("02-01-2006\n", text)

		if readErr == nil && timeErr == nil {
			return date
		}
	}
}

func getResult(birth time.Time) float64 {
	var now = time.Now()
	var diff = now.Sub(birth)

	return diff.Seconds()
}
