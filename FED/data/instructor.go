package data

type Instructor struct {
	Id   int
	Name string

	// "private" variable because the first letter is not capitalized
	penisSize float32
}

func (instructor Instructor) PrintName() {
	print(instructor.Name)
}

// there is no pattern to define a constructor
// but there are factories

func NewInstructor(name string) Instructor {
	return Instructor{Name: name}
}
