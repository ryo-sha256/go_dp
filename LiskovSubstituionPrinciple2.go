package main

import "fmt"

type A struct {
}

// ADDING THIS ...
// make the interface that has the function that you want to use
// this is newly created interface
type tester interface {
	Test()
}

func (a A) Test() {
	fmt.Println("Printing A")
}

type B struct {
	A
}

// before
func ImpossibleLiskovSubstition(a A) {
	a.Test()
}

// after
func BetterImpossibleLiskovSubstition(a tester) {
	a.Test()
}

/*
func main() {
	// before
	a := B{}
	// this will cause the panic because you canot use
	// a type B as type A in argument to ImpossibleLiskovSubstition function
	// in Go, there's no inheritance
	ImpossibleLiskovSubstition(a)


	//after
	a := B{}
	BetterImpossibleLiskovSubstition(a)


}





*/
