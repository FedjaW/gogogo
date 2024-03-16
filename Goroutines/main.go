package main

import (
	"fmt"
	"time"
)

func printMessage(text string) {
	for i:=0; i<5; i++ {
		fmt.Println(text)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() { // main goroutine
	go printMessage("go is great")
	go printMessage("programming is fun")
	go printMessage("i love go so much more than js")
}

// the app ends when the main go routine ends
// the main goroutine ends when the "}" of the main function closes
// the main function will open three go routines and end immediately
// so there is no time for the other threads to print something to the screen
