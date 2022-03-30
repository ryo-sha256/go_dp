package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var counter = 0

type Project struct {
	files []string
}
type Persistance struct {
	lineSeparator string
}

// after making Save function below
func (f *Project) String() string {
	return strings.Join(f.files, "\n")
}

// Single Responsibility Principle
// @return current counter
func (f *Project) AddFile(fileName string) int {
	counter++
	addition := fmt.Sprintf("%d: %s",
		counter, fileName)
	f.files = append(f.files, addition)
	return counter
}

// one function is responsible for one thing
func (f *Project) DeleteFile(fileName string) {
	//

}

// separation of concerns
// you don't want to put everything into one sink
// IOW, one function does one thing
// you might want to save the file?
func (f *Project) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(f.String()), 0644)
}

func (f *Project) Load(filename string) {

}

func (f *Project) LoadFromWeb(url *url.URL) {

}

var lineSeparator = "\n"

func SaveToFile(f *Project, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(f.files, lineSeparator)), 0644)
}

var LineSeparator = "\n"

func (p *Persistance) SaveToFile(f *Project, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(f.files, LineSeparator)), 0644)
}

/*
func main() {

	fmt.Println("this is test run")
	p := Project{}

	p.AddFile("first file")
	p.AddFile("second file")

	fmt.Println(p.String())

	// in order to persist to a file
	// recommended to use a separate function

	SaveToFile(&p, "filename.txt")
	// or
	persist := Persistance{"\r\n"}
	persist.SaveToFile(&p, "filename2.txt")
}
*/
