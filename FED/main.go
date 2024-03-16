package main

import (
	"fmt"

	"leckomio.dev/go/fed/data"
)

func main() {
	// default initialization
	fed := data.Instructor{Id: 1, Name: "FED"}
	fed.PrintName()

	// factory design pattern of go
	hans := data.NewInstructor("hans")
	hans.PrintName()

	goCourse := data.Course{Id: 1, Name: "GoCourse", Duration: 12.3}
	fmt.Println(goCourse)

	goWorkshop := data.NewWorkshop("goWorkshop", fed)

	// Signable is an interface
	// Course and Workshop implement this interface
	// So we can emulate polymorphism
	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = goWorkshop

	for _, course := range courses {
		fmt.Println(course)
	}
}
