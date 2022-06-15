package main

import (
	openclosed "design-patterns-in-go/SOLID/open_closed"
	srp "design-patterns-in-go/SOLID/single_responsability"
	"fmt"
)

func main() {

	// srp
	j := &srp.Journal{}
	j.AddEntry("aaa")
	j.AddEntry("bbb")
	j.AddEntry("ccc")
	fmt.Println(j.String())
	srp.Save(j, "JournalEntries.txt")

	// ocp

	apple := &openclosed.Product{Name: "Apple", Color: openclosed.Green, Size: openclosed.Small}
	tree := &openclosed.Product{Name: "Tree", Color: openclosed.Green, Size: openclosed.Large}
	house := &openclosed.Product{Name: "House", Color: openclosed.Blue, Size: openclosed.Medium}

	products := []openclosed.Product{*apple, *tree, *house}
	f := openclosed.Filter{}

	// Bad way
	fmt.Println("Green products (old):")
	for _, v := range f.FilterByColor(products, 1) {
		fmt.Printf(" - %s is green \n", v.Name)
	}

	// Better way
	fmt.Println("Green products (new):")
	greenSpec := openclosed.ColorSpecification{Color: openclosed.Green}
	bf := openclosed.BetterFilter{}
	for _, v := range bf.Filter(products, &greenSpec) {
		fmt.Printf(" - %s is green \n", v.Name)
	}

	largeSpec := openclosed.SizeSpecification{Size: openclosed.Large}
	lgSpec := openclosed.AndSpecification{FirstSpec: greenSpec, SecondSpec: largeSpec}

	fmt.Println("Large green and products:")
	for _, v := range bf.Filter(products, &lgSpec) {
		fmt.Printf(" - %s is large and green \n", v.Name)
	}
}
