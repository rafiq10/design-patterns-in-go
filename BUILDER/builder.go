package main

import (
	bf "builder/builder_facets"
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
}
