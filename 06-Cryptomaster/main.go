package main

import (
	"fmt"
	"sync"

	"leckomio.dev/go/crypto/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)

			// go supports closures
			// hence we can access local variables that are part of the outer function
			wg.Done()
		}(currency) // immediately execute lambda func
	}
	wg.Wait()
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
