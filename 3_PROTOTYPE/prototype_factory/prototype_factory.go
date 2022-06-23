package prototypefactory

// we want to make a factory based on mainOffice and auxOffice prototypes
// a convenience approach
import (
	"bytes"
	"encoding/gob"
	"fmt"
)

var mainOffice = Employee{
	"", Address{
		Suite: 0, StreetAddress: "123 East Dr", City: "London",
	},
}
var auxOffice = Employee{
	"", Address{
		Suite: 0, StreetAddress: "66 West Dr", City: "London",
	},
}

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek into structure
	fmt.Println(b.Bytes())

	d := gob.NewDecoder(&b)
	result := Employee{}
	d.Decode(&result)
	return &result
}

func NewEmployee(proto *Employee, name string, suite int) *Employee {
	e := proto.DeepCopy()
	e.Name = name
	e.Office.Suite = suite
	return e
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return NewEmployee(&mainOffice, name, suite)
}
func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return NewEmployee(&auxOffice, name, suite)
}
