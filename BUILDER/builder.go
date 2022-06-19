package main

import (
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
}
