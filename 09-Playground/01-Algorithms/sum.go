package main

func Sum(numbers []float32) float32 {
	var sum float32 = 0
	for _, number := range numbers {
		sum = sum + number
	}

	return sum
}