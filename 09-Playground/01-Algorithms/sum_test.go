package main

import "testing"

func Sum_PositiveNumbers(t *testing.T) {
	// arrange
	numbers := []float32{5, 3, 99, 3}

	// act
	result := Sum(numbers)

	// assert
	if result != 110 {
		t.Error("Sum should be 110")
	}
}

func Sum_PositivAndNegativeNumbers(t *testing.T) {
	// arrange
	numbers := []float32{-5, 3, -99, 3}

	// act
	result := Sum(numbers)

	// assert
	if result != -98 {
		t.Error("Sum should be -98")
	}
}

func Sum_FloatNumbers(t *testing.T) {
	// arrange
	numbers := []float32{3.2, 2.3, 19, 0.3, 0.7}

	// act
	result := Sum(numbers)

	// assert
	if result != 25.5 {
		t.Error("Sum should be 25.5")
	}
}