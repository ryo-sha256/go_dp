package main

type PersonC struct {
	name, position string
}

type personMod func(*PersonC)

// 0, make the struct that can contains all the actions for the personMod
// personMod is used for BuilderParameter pattern
type PersonCBuilder struct {
	actions []personMod
}

// fluent parameter
// 1, make the first method that is going to be in actions[]
func (b *PersonCBuilder) Called(name string) *PersonCBuilder {
	b.actions = append(b.actions, func(p *PersonC) {
		p.name = name
	})
	return b
}

// 2, builder function that actually allocates a new object
// and traverse the actions list and call each of the assigned methods
func (b *PersonCBuilder) Build() *PersonC {
	p := PersonC{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

// 3, just responsible for appending the position data to actions array[]
func (b *PersonCBuilder) WorksAsA(position string) *PersonCBuilder {
	// you are appending function,
	// but not "Calling" them
	b.actions = append(b.actions, func(p *PersonC) {
		p.position = position
	})
	return b
}

/*
func main() {

	// 3, here I can make a builder
	b := PersonCBuilder{}

	// Build() should be placed at the end
	p := b.Called("Ryo").
		WorksAsA("DMG").
		Build()
	fmt.Println(*p)
}
*/
