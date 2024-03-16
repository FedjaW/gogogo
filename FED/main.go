package main

import "leckomio.dev/go/fed/data"

func main() {
	fed := data.Instructor{Id: 1, Name: "FED"}

	fed.PrintName()
}
