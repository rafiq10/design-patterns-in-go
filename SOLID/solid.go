package solid

import (
	"fmt"
	dip "solid/dependency_inversion"
	lsp "solid/liskov_substitution"
	ocp "solid/open_closed"
	srp "solid/single_responsability"
)

func main() {

	// 	// srp
	j := &srp.Journal{}
	j.AddEntry("aaa")
	j.AddEntry("bbb")
	j.AddEntry("ccc")
	fmt.Println(j.String())
	srp.Save(j, "JournalEntries.txt")

	// 	// ocp
	apple := &ocp.Product{Name: "Apple", Color: ocp.Green, Size: ocp.Small}
	tree := &ocp.Product{Name: "Tree", Color: ocp.Green, Size: ocp.Large}
	house := &ocp.Product{Name: "House", Color: ocp.Blue, Size: ocp.Medium}

	products := []ocp.Product{*apple, *tree, *house}
	f := ocp.Filter{}

	// Bad way
	fmt.Println("Green products (old):")
	for _, v := range f.FilterByColor(products, 1) {
		fmt.Printf(" - %s is green \n", v.Name)
	}

	// Better way
	fmt.Println("Green products (new):")
	greenSpec := ocp.ColorSpecification{Color: ocp.Green}
	bf := ocp.BetterFilter{}
	for _, v := range bf.Filter(products, &greenSpec) {
		fmt.Printf(" - %s is green \n", v.Name)
	}

	largeSpec := ocp.SizeSpecification{Size: ocp.Large}
	lgSpec := ocp.AndSpecification{FirstSpec: greenSpec, SecondSpec: largeSpec}

	fmt.Println("Large green and products:")
	for _, v := range bf.Filter(products, &lgSpec) {
		fmt.Printf(" - %s is large and green \n", v.Name)
	}

	// lsp

	rc := &lsp.Rectangle{Width: 2, Height: 3}
	lsp.UseIt(rc)

	sq := lsp.NewSquare(5)
	lsp.UseIt(sq)

	sq2 := lsp.Square2{Size: 5}
	rc2 := sq2.Rectangle()
	lsp.UseIt(&rc2)

	// dip
	parent := dip.Person{Name: "Janusz"}
	child1 := dip.Person{Name: "Rafal"}
	child2 := dip.Person{Name: "Marta"}
	rels := dip.Relationships{}
	rels.AddParentAndChild(&parent, &child1)
	rels.AddParentAndChild(&parent, &child2)

	// r depends on rels
	r := dip.Research{Relationships: rels}
	r.Investigate()

	// r2 depends on ResearchBrowser inerface - a better approach
	r2 := &dip.Research2{Browser: &rels}
	r2.Investigate()
}
