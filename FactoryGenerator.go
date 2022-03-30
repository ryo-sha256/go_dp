package main

// 0, we want to create a factory depending upon
// the setting that we want employeess to be subsequently manufacutured
// if we can do that in one go / one statement, it would be nice

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// 1, functional approach
// a function that returns a function
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {

	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// 3, making a factory "struct"
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

// 4, this returns the pointer that is pointing at struct
func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

// 5, newer version
func NewEmployeeFactory2(position string,
	annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

/*
func main() {

	developerFactory := NewEmployeeFactory("developer", 80000)
	managerFactory := NewEmployeeFactory("manager", 80000)
	// 2, now you can pass these funciotns to be consumed by other functions
	// by consume, it means "invoking"
	developer := developerFactory("Adam")
	manager := managerFactory("Jane")


	// 6, new position : boss factory
	bossFactory := NewEmployeeFactory2("CEO", 100000)
	boss := bossFactory.Create("Sam")
} */
