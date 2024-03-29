package main

import (
	"fmt"

	"leckomio.dev/go/io/data"
)

func init() {
	fmt.Println("Init func 1")
}

func init() {
	fmt.Println("Init func 2")
}

func calcTax(price float32) (float32, float32) {
	return price * 0.9, price * 0.2
}

func calcTax2(price float32) (a float32, b float32) {
	a = price * 0.9
	b = price * 0.2
	return
}

var newPrice = 99

func birthday(age *int) {
	if *age > 110 {
		panic("To old to be true")
	}
	fmt.Printf("Thi is a pointer %v and this the value %v\n", age, *age)
	*age++
}

func main() {
	defer fmt.Println("Bye")
	defer fmt.Println("Good")

	message := "Hello from Go"
	var price int = 12
	fmt.Println(message, price, newPrice)
	sayHello()
	fmt.Println(data.MaxSpeed)
	fmt.Println(data.Countries)

	result, _ := calcTax(12.0)
	fmt.Println(result)

	result2, _ := calcTax2(12.0)
	fmt.Println(result2)

	age := 22
	birthday(&age)
	fmt.Println(age)
	fmt.Println(&age)
	fmt.Println(*&age)
}
