package main

import "fmt"

// * high/low level module
// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstraction
type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type PersonB struct {
	name string
}

// how do we model the relationship between different people
type Info struct {
	from         *PersonB
	relationship Relationship
	to           *PersonB
}

// low-level module
// the abstact that low-level and high level module will rely on
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*PersonB
}

// this is low-lelvel module
// how do we go about storing those information
type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*PersonB {
	//	panic("impltement me")
	// result would be the empty slice of person pointer first
	result := make([]*PersonB, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(
	parent, child *PersonB) {
	// it should be bi-directional
	r.relations = append(r.relations,
		Info{parent, Parent, child})
	r.relations = append(r.relations,
		Info{child, Child, parent})
}

// high level module
// here we can do Dependency Inversion Principle
// this does work, however it violates DIP
type Research struct {
	// >>> before
	// break DIP
	// relationships Relationships

	// after <<<<
	// now we have the abstraction that Reserach struct can depend on
	browser RelationshipBrowser
}

// and then you define the method that can be called from the Research struct
// e.g. you can find all the relationships from John
func (r *Research) Investigate() {
	// !!major problem
	// relationships uses the internal module here
	// if the r.relaptionships.relations is database storage
	// so code depeding on the low level module will break
	/*
		relations := r.relationships.relations // relations []Info
		for _, rel := range relations {
			if rel.from.name == "John" {
				if rel.relationship == Parent {
					fmt.Println("John has a child called ", rel.to.name)
				}
			}
		}
	*/

	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}

}

// adding child/parent or remove
/*
func main() {

	// >>>> before
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	//we can create relaptionship here
	relationships := Relationships{}
	// add rel of John and Chris
	relationships.AddParentAndChild(&parent, &child1)
	// add rel of John and Matt
	relationships.AddParentAndChild(&parent, &child2)



	// sine interface "RelationshipBrowser"
	// is { browser *Person[] }
	r := Research{&relationships}
	r.Investigate()
	// <<<<after

}
*/
