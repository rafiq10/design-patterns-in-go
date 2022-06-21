package builderfacets

// We want to have separate builders for address and job of a Person
type Person struct {
	// address
	StreetAddress, PostCode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

// PersonAddressBuilder and PersonJobBuilder aggregate to PersonBuilder
type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

// Lives() returns a pointer to PersonAddressBuilder and gives access to its methods (At, In, WithPostCode)
// So we can chain the building of the PersonBuilder object
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	//copies the pointer of PersonBuilder
	return &PersonAddressBuilder{*b}
}

// Lives() returns a pointer to PersonAddrPersonJobBuilderessBuilder and gives access to its methods (At, AsA, Earning)
// So we can chain the building of the PersonBuilder object
func (b *PersonBuilder) Works() *PersonJobBuilder {
	//copies the pointer of PersonBuilder
	return &PersonJobBuilder{*b}
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}
func (b *PersonAddressBuilder) WithPostCode(postCode string) *PersonAddressBuilder {
	b.person.PostCode = postCode
	return b
}
func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}
func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}
func (b *PersonJobBuilder) Earning(income int) *PersonJobBuilder {
	b.person.AnnualIncome = income
	return b
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}
