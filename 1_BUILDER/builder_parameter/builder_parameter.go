package builderparameter

import (
	"strings"
)

// hide the email struct from API user
type email struct {
	from, to, subject, body string
}

// use the builder to properly construct the "email" struct
type EmailBuilder struct {
	email email
}

// Use the fluid builder with some validation
func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain '@': " + from)
	}
	b.email.from = from
	return b
}

// Use the fluid builder with some validation
func (b *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("email should contain '@': " + to)
	}
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

func sendMailImpl(email *email) {

}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	// the API client has to provide the action function that initializes email object through the builder
	action(&builder)
	// the sendMailImpl function uses the email object hidden behind the builder
	sendMailImpl(&builder.email)
}
