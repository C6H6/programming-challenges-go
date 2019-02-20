package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "FILENAME")
		return
	}

	file := os.Args[1]
	count, err := count(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(count)
}

func count(name string) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		sum += countWordsInString(line)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sum, nil
}

func countWordsInString(str string) int {
	regex := regexp.MustCompile("\\w+")
	matches := regex.FindAllStringIndex(str, -1)
	return len(matches)
}
