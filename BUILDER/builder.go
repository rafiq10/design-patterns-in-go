package main

import (
	bf "builder/builder_facets"
	bp "builder/builder_parameter"
	cb "builder/creational_builder"
	"fmt"
)

func main() {
	b := cb.NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())

	// chaining the calls
	b2 := cb.NewHtmlBuilder("ul").
		AddChildFluent("li", "hello").
		AddChildFluent("li", "world")
	fmt.Println(b2.String())

	// Builder Facets
	pb := bf.NewPersonBuilder()
	pb.
		Lives().
		At("Gran Via").
		In("Madrid").
		WithPostCode("28001").
		Works().
		At("Telefonica").
		AsA("Programmer").
		Earning(50000)
	person := pb.Build()
	fmt.Println(person)

	// builder parameter
	// from the client's perspective they have to call SendEmail
	// and the user sees that there is a function to be provided to
	bp.SendEmail(func(b *bp.EmailBuilder) {
		// you don't have the access to the email object itself
		// you just have an access to the builder
		// the client of the API has the option to initialize the email object through the action function
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Hello").
			Body("Nice to meet you")
	})
}
