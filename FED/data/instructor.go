package data

type Instructor struct {
	Id      int
	Name string

	// "private" variable because the first letter is not capitalized
	penisSize float32
}

func (instructor Instructor) PrintName() {
	print(instructor.Name)
}
