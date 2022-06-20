package interfacefactory

import "fmt"

type Person interface {
	SayHello()
}

// the type is not exposed
type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, and I am %d years old!\n", p.name, p.age)
}

// we can have 2 different structs under he same interface
func (p *tiredPerson) SayHello() {
	fmt.Printf("Hi, I am too tired\n")
}

// depending on age we can return one or another type
func NewPerson(name string, age int) Person {
	if age > 70 {
		return &tiredPerson{name: name, age: age}
	}
	return &person{name: name, age: age}
}
