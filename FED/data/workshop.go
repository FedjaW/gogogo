package data

import "time"

type Workshop struct {
	Course // embedded
	Name string // notice the the Course has also a Name
	Date time.Time
}

// Factory design pattern from go (must start with New)
func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{Name: name} // can't access Course.Name in constructor/initializer (what is it called?) of Workshop
	w.Course.Name = instructor.Name
	return w
}

// to demonstrate interfaces
func (worshop Workshop) SignUp() bool {
	return true
}