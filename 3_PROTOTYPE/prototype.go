package main

import (
	"fmt"
	cm "prototype/copy_method"
	cts "prototype/copy_through_serialization"
	dc "prototype/deep_copying"
	pf "prototype/prototype_factory"
)

func main() {
	john := dc.Person{Name: "John", Address: &dc.Address{StreetAddress: "123 London Rd", City: "London", Country: "UK"}}
	// this is copying the pointer
	jane := john
	jane.Name = "Jane"

	// this changes he object under the pointer to the  Address, so it changes also John's address
	jane.Address.StreetAddress = "321 Baker St"
	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

	// deep copy example
	// it requires a lot of code to just make a copy of an object
	john = dc.Person{Name: "John", Address: &dc.Address{StreetAddress: "123 London Rd", City: "London", Country: "UK"}}
	jane = john
	jane.Name = "Jane"
	jane.Address = &dc.Address{
		StreetAddress: john.Address.StreetAddress,
		City:          john.Address.City,
		Country:       john.Address.Country,
	}
	jane.Address.StreetAddress = "321 Baker St"
	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

	// DeepCopy() method
	rafal := &cm.Person{
		Name: "Rafa", Address: &cm.Address{StreetAddress: "123 London Rd", City: "London", Country: "UK"},
		Friends: []string{"Chris", "Matt"},
	}
	marta := rafal.DeepCopy()
	marta.Name = "Marta"
	marta.Address.StreetAddress = "321 Baker St"
	marta.Friends = append(marta.Friends, "Klaudia")
	fmt.Println(rafal, rafal.Address)
	fmt.Println(marta, marta.Address)

	// DeepCopy() deserialization
	robert := &cts.Person{
		Name: "Robert", Address: &cts.Address{StreetAddress: "123 London Rd", City: "London", Country: "UK"},
		Friends: []string{"Chris", "Matt"},
	}
	zuza := robert.DeepCopy()
	zuza.Name = "Zuza"
	zuza.Address.StreetAddress = "321 Baker St"
	zuza.Friends = append(zuza.Friends, "Klaudia")
	fmt.Println(robert, robert.Address)
	fmt.Println(zuza, zuza.Address)

	// prototype factory
	andrea := pf.NewMainOfficeEmployee("Andrea", 100)
	paula := pf.NewAuxOfficeEmployee("Paula", 133)
	fmt.Println(andrea, andrea.Office)
	fmt.Println(paula, paula.Office)
}
