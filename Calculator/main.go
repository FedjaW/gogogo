package main

import "fmt"

func main() {

	var operation string
	var number1, number2 int

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("=================")
	fmt.Println("Which operation you want to perform, add, sub")
	fmt.Scanf("%s", &operation)
	fmt.Println("Number 1")
	fmt.Scanf("%d", &number1)
	fmt.Println("Number 2")
	fmt.Scanf("%d", &number2)

	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "sub":
		fmt.Println(number1 - number2)
	}

	fmt.Printf("N1 %v, N2 %v\n", number1, number2)
}
