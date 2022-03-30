package main

// summary
// To make a buider fluent, return the receiver = allows chanigng

// 1, there are some situation where you need more than 1 builder
// and you need to separate them
type Person struct {
	// address
	StreetAddress, Postcode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

// 4. adding utility methods to the PersonBuilder struct
// which returns the pointer that points to PersonAddressBuilder
// NOTICE: this is where one struct has two methods to AddressBuilder and JobBuilder
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	// dereference
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	// dereference
	return &PersonJobBuilder{*b}
}

// 5, now we have a way of transitioning from PersonBuilder
// to PersonAddressBuilder or PersonJobBuilder
// now you can jump to PersonJobBuilder from the original PersonBuilder
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		&Person{}}
}

// 3, adding additional builders
type PersonAddressBuilder struct {
	PersonBuilder
}

// 6, adding the method to a particular builder now
func (b *PersonAddressBuilder) At(
	streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(
	city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostCode(
	postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

type PersonJobBuilder struct {
	PersonBuilder
}

// 6`-continued
func (pjb *PersonJobBuilder) At(
	companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(
	position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(
	annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

// 7, now need one method to effectively use the two different method
// actually creating an object
func (b *PersonBuilder) Build() *Person {
	return b.person
}

// 2, we can have additional builder,
// so instead of having everything in one builder
// better to have separated builder that is responsible for one thing
/*
func main() {

	pb := NewPersonBuilder()
	pb.
		Lives().
		At("fort street").
		In("Victoria").
		WithPostCode("23111").
		Works().
		At("DMG").
		AsA("Software Engineer").
		Earning(65000)

	// actually returns the value
	person := pb.Build()
	fmt.Println(person)

}
*/
