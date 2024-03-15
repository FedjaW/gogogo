package data

import "fmt"

var Countries [10]string
var WellKnownPorts = map[string]int { "http": 80 }
// var codes map[int]string

func init() {
	fmt.Println("Init func from collections")
	Countries[0] = "Asemissen"
	Countries[9] = "Leo"
}