package functionalbuilder

type Person struct {
	name, position string
}

type personMod func(*Person)

// in the personBuilder we want to be able to specify that a person has been given a name
type PersonBuilder struct {
	actions []personMod
}

// Fluent interface to make chaning the calls possible
func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

// Fluent interface to make chaning the calls possible
func (b *PersonBuilder) WorksAs(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := &Person{}
	for _, a := range b.actions {
		a(p)
	}
	return p
}
