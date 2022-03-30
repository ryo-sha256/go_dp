package main

import "fmt"

// in my personal opinion,
// interface factory is the simplest and the most
// intuitive factory pattern in go

type PersonE interface {
	SayHello()
}

// 0, this time, the struct will be private
// meaning that this struct will be hidden from the consumers
type personE struct {
	name string
	age  int
}

type tiredPersonE struct {
	name string
	age  int
}

// 2, you can have another struct that has the same method
// but it generates a different output
func (p *personE) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old\n", p.name, p.age)
}

func (p *tiredPersonE) SayHello() {
	fmt.Printf("Sorry I'm tired to talk to you")
}

// notice here, it returns the interface that is declared at the top of the file
func NewPerson(name string, age int) PersonE {
	// 3, you can do kinda like a method overloading
	// by putting the logic inside of the interface builder function
	if age > 100 {
		// 4, this way, we can have different types of object in background
		return &tiredPersonE{name, age}
	}
	return &personE{name, age}
}

// 1, this time, it returns the interface

/*
func main() {

	// 4, now the type of method is different
	// based on the invocation of the object here
	p := NewPerson("James", 34)
	p.SayHello()
}
*/
