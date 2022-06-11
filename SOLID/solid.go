package main

import (
	srp "design-patterns-in-go/SOLID/single_responsability"
	"fmt"
)

func main() {
	j := &srp.Journal{}
	j.AddEntry("aaa")
	j.AddEntry("bbb")
	j.AddEntry("ccc")
	fmt.Println(j.String())
	srp.Save(j, "JournalEntries.txt")
}
