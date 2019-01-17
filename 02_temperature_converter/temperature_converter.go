package main

import (
	"fmt"
)

func main() {
	var inputScale = getScaleChoice("Input scale")
	var outputScale = getScaleChoice("Output scale")
	var value = getValue()
	fmt.Println("Result:", convert(inputScale, outputScale, value))
}

func getScaleChoice(questionText string) int {
	fmt.Println("1. Celsius\n2. Fahrenheit\n3. Kelvin")
	var i int
	for {
		fmt.Printf("%s: ", questionText)
		_, err := fmt.Scanf("%d", &i)
		if err == nil && i >= 1 && i <= 3 {
			return i
		}
	}
}

func getValue() int {
	var i int
	for {
		fmt.Print("Value: ")
		_, err := fmt.Scanf("%d", &i)
		if err == nil {
			return i
		}
	}
}

func convert(inputScale int, outputScale int, value int) float32 {
	switch inputScale {
	case 1:
		switch outputScale {
		case 1:
			return float32(value)
		case 2:
			return float32(value*9/5 + 32)
		case 3:
			return float32(value) + 273.15
		}
	case 2:
		switch outputScale {
		case 1:
			return float32((value - 32) * 5 / 9)
		case 2:
			return float32(value)
		case 3:
			return (float32(value) + 459.67) * 5 / 9
		}
	case 3:
		switch outputScale {
		case 1:
			return float32(value) - 273.15
		case 2:
			return float32(value)*9/5 - 459.67
		case 3:
			return float32(value)
		}
	}
	return 0.0
}
