package main

// factory is the design pattern to create an object in one go
// while builder is to more slowly/gradually create an object
// like when struct has too many fields
type PersonD struct {
	Name string
	Age  int
	// default property
	// you don't want to keep typing this constant variable
	EyeCount int
}

// SO, factory function is nothing more than
// a freestanding function which returns an instance of this struct that you want to create

func NewPersonD(name string, age int) *PersonD {

	// some verification
	if age < 16 {
		// ....
	}

	// notice the default value here
	return &PersonD{name, age, 2}
	// this will return the pointer that is pointing to the PersonD object
}

/*
func main() {
	p := NewPersonD("John", 30)

}
*/
