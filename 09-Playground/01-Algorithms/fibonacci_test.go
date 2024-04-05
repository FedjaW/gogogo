package main

import (
	"testing"
)

func Test_Fibonacci_10_ResultsIn_55(t *testing.T) {
	// arrange
	n := 10

	// act
	result := Fibonacci(n)

	// assert
	if result != 55 {
		t.Error("n=10 should result in 55")
	}
}

func Test_Fibonacci_20_ResultsIn_6765(t *testing.T) {
	// arrange
	n := 20

	// act
	result := Fibonacci(n)

	// assert
	if result != 6765 {
		t.Error("n=20 should result in 6765")
	}
}

func Test_Fibonacci_0_ResultsIn_0(t *testing.T) {
	// arrange
	n := 0

	// act
	result := Fibonacci(n)

	// assert
	if result != 0 {
		t.Error("n=0 should result in 0")
	}
}

func Test_Fibonacci_1_ResultsIn_1(t *testing.T) {
	// arrange
	n := 1

	// act
	result := Fibonacci(n)

	// assert
	if result != 1 {
		t.Error("n=1 should result in 1")
	}
}

func Test_Fibonacci_2_ResultsIn_1(t *testing.T) {
	// arrange
	n := 2

	// act
	result := Fibonacci(n)

	// assert
	if result != 1 {
		t.Error("n=2 should result in 1")
	}
}