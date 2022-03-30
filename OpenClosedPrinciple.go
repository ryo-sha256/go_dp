package main

// OCP
// open for extension, closed for modification

// imagine you are operating a physical store
// i.e.  filtered by price?

type Color int
type Size int

const (
	red Color = iota
	green
	blue
)

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

// filter type
type Filter struct {
	// filterbyColor?
}

// filterd by color
// @return the array of pointers that point products which matches the passed color
func (f *Filter) FilterByColor(
	products []Product, color Color) []*Product {
	// make a memory allocation of the array for those pointers
	result := make([]*Product, 0)

	// for loop
	// index, value
	for i, v := range products {
		if v.color == color {
			// for the append you are supposed to pass the reference
			result = append(result, &products[i])
		}
	}
	return result
}

// filter by size
// @return the array of pointers that point products whose size is passed size
func (f *Filter) FilterBySize(
	products []Product, size Size) []*Product {

	result := make([]*Product, 0)
	// for loop
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// filter by size and color
// @return the array of pointers that point products whose size and color match the passed variables
func (f *Filter) FilterBySizeAndColor(
	products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// specification interface
type Specification interface {
	IsSatisfied(p *Product) bool
}

/////////////////////////
type ColorSpecification struct {
	color Color
}

func (s ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == s.color
}

///////////////////////////
type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

////////////////////////

// example of "composite specification"
type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type BetterFilter struct{}

// open for extension, closed for modification
func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

/*
func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")

	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green", v.name)
	}

	fmt.Println("filter by size")
	for _, v := range f.FilterBySize(products, large) {
		fmt.Printf(" - %d is green", v.size)
	}

	fmt.Printf("filter by color and size")
	for _, v := range f.FilterBySizeAndColor(products, large, green) {
		fmt.Printf(" - %s is green and large", v.name)

	}

	// specification
	fmt.Printf("Green products (new) \n")
	greenSpec := ColorSpecification{green}
	// separation and call it
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	//now we can combine large and green for the specification
	greenAndLargeSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large green products:\n")
	for _, v := range bf.Filter(products, greenAndLargeSpec) {
		fmt.Printf("found this: large and green -> %s", v.name)
	}

}
*/
