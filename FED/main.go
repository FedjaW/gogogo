package main

import (
	"fmt"

	"leckomio.dev/go/fed/data"
)

func main() {
	fed := data.Instructor{Id: 1, Name: "FED"}
	fed.PrintName()

	hans := data.NewInstructor("hans")
	hans.PrintName()

	goCourse := data.Course{Id: 1, Name: "GoCourse", Duration: 12.3}
	fmt.Println(goCourse)
}
