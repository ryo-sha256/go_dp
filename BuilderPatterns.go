package main

import (
	"fmt"
	"strings"
)

// some objects are simple and can be created in a single constructor call(=factory pattern)
// other objects require a lot of ceremony to create
// having a factory function with 10 arguments is not productive

// when piece-wise object construction is complicated,
// provide an API for doing it succinctly
const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

// string method
type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

// utility function
func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName,
		HtmlElement{
			rootName,
			"",
			[]HtmlElement{},
		},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

// private method?
func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)

	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.name))

	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	// recursive call here
	for _, el := range e.elements {
		sb.WriteString(el.string((indent + 1)))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n",
		i, e.name))
	return sb.String()
}

func (b *HtmlBuilder) AddChildFluent(
	childName, childText string) *HtmlBuilder {

	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

// suppose you are writing a web server
/*
func main() {
	hello := "hello"
	//string builder
	sb := strings.Builder{}

	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")

	fmt.Println(sb.String())

	words := []string{"hello", "world"}
	sb.Reset()
	// <ul><li>..</li>...........
	sb.WriteString("<ul>")

	for _, v := range words {
		// so many steps to go through => builder pattern
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("<li>")
	}

	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	// after
	b := NewHtmlBuilder("ul")

	b.AddChildFluent("li", "hello").AddChildFluent("li", "hello")
	fmt.Println(b.String())
}
*/
