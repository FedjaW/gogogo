package main

import (
	"fmt"

	"leckomio.dev/go/crypto/api"
)

func main() {
	go getCurrencyData("BTC")
	go getCurrencyData("ETH")
	go getCurrencyData("BCH")
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	// The . operator works on pointers
	// it will go the memory to that the pointer is pointing and take the value
	// print(rate.Price, err)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The rate for %v is %.2f Euro \n", rate.Currency, rate.Price)
}
