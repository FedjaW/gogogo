package main

import (
	"slices"
	"testing"
)

func Test_BubbleSort_WithNegativeNumbers(t *testing.T) {
	// arrange
	a := []float32{2.3, 0.1, -3.2, 99, -10}

	// act
	result := BubbleSort(a)

	// assert
	expectedResult := []float32{-10, -3.2, 0.1, 2.3, 99}
	if !slices.Equal(result, expectedResult) {
		t.Error("Array not sorted correct ")
	}
}

func Test_BubbleSort_PositiveNumbersOnly(t *testing.T) {
	// arrange
	a := []float32{2.3, 0.1, 3.2, 99, 10}

	// act
	result := BubbleSort(a)

	// assert
	expectedResult := []float32{0.1, 2.3, 3.2, 10, 99}
	if !slices.Equal(result, expectedResult) {
		t.Error("Array not sorted correct ")
	}
}