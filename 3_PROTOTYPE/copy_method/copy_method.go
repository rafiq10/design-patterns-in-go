package copymethod

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// still quite a lot of job to copy
func (p *Person) DeepCopy() *Person {
	// IMPORTANT: we want to copy the pointer, so "*"p"
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}
