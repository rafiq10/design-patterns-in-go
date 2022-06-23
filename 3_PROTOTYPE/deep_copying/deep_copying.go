package deepcopying

type Person struct {
	Name    string
	Address *Address
}

type Address struct {
	StreetAddress, City, Country string
}
