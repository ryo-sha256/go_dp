package main

import "strings"

// passing the builder function as PARAMETER
// called "builderParameter

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

// 0, Builder parameter
// *it's called fluent interface
func (b *EmailBuilder) From(from string) *EmailBuilder {
	// you could add some error handlings
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

// 0' have some additional
func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

// 1, somewhere behind the scene, you want to have a function called
// sendEmailImpl but you don't want clients to work on the email object
// ultimately you just want them to work on ONLY builders
func sendMailImpl(email *email) {

}

// 2, this is going to be a function that applies to a builder
// it TAKES a builder and do something,, calls on builder
type build func(*EmailBuilder)

// notice this is a public method
func SendEmail(action build) {
	builder := EmailBuilder{}
	// action is type function
	// you actually call the parameter (look at the main function)
	action(&builder)
	// calling
	sendMailImpl(&builder.email)
}

/*
func main() {

	// 3, look at the parameter
	// you pass the builder function that accepts those parameters
	// to BUILD the object
	// this is why it's called builder parameters
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("to@gmail.com").
			Subject("temp title").
			Body("this is the message")
	})

}
*/
