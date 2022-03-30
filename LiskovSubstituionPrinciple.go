package main

import "fmt"

// LSP
// suppose that you are trying to deal with geometric shapes
// there is no right answear to this problem
// but the problem is that UseIt is designed to
// calculate only rectangle but not the square
// you might have getter and setter for height
type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int // use lower case here
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// suppose that you wanna add more
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// doing the same thing for the square......<- bad
func (s *Square) SetWidth(width int) {
	s.width = width
}

// it is like a driver of the Sized interface
func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected area", expectedArea)
	fmt.Print("Actual area ", actualArea)
}

/*
func main() {
	// address of Rectangle object
	rc := &Rectangle{3, 2}
	UseIt(rc)

}
*/
