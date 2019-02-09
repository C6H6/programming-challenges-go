package main

import (
	"testing"
)

func TestConvertProperlyCelsius2Fahrenheit(t *testing.T) {
	result := convert(1, 2, 100)

	if result != 212 {
		t.Errorf("Wrong result %f for value %d", result, 100)
	}
}

func TestConvertProperlyFahrenheit2Celsius(t *testing.T) {
	result := convert(2, 1, 100)

	if result != 37 {
		t.Errorf("Wrong result %f for value %d", result, 100)
	}
}

func TestConvertProperlyFahrenheit2Kelvin(t *testing.T) {
	result := convert(2, 3, 100)

	if result != 310.9278 {
		t.Errorf("Wrong result %f for value %d", result, 100)
	}
}

func TestConvertProperlyKelvin2Fahrenheit(t *testing.T) {
	result := convert(3, 2, 100)

	if result != -279.67 {
		t.Errorf("Wrong result %f for value %d", result, 100)
	}
}

func TestConvertProperlyKelvin2Celsius(t *testing.T) {
	result := convert(3, 1, 100)

	if result != -173.149994 {
		t.Errorf("Wrong result %f for value %d", result, 100)
	}
}

func TestConvertProperlyCelsius2Kelvin(t *testing.T) {
	result := convert(1, 3, 100)

	if result != 373.15 {
		t.Errorf("Wrong result %f for value %d", result, 100)
	}
}
