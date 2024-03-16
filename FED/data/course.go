package data

import "fmt"

type DurationInHours float32

type Course struct {
	Id       int
	Name     string
	Duration DurationInHours
}

// fmt (here in main) is looking for a method with name "String"
// then it will use this when you print the course struct
func (c Course) String() string {
	return fmt.Sprintf("name: %v, duration %v", c.Name, c.Duration)
}
