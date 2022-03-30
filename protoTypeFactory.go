package main

// summary: you can have a preconfigured object

// related to prototype
// you can have preconfigured object

type Employee2 struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee2 {
	switch role {
	case Developer:
		return &Employee2{"", "developer", 60000}
	case Manager:
		return &Employee2{"", "manager", 80000}

	default:
		panic("Unsupported Role")
	}
}

func main() {

	m := NewEmployee(Manager)
	// now you have the access to the declared object
	m.Name = "Sam"

}
