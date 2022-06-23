package copythroughserialization

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// we use "encoding/gob" to deserialize bytes
// this library knows that when there is a pointer it has to get the value of the pointer
// It is valid for every object
func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(b.Bytes())
	d := gob.NewDecoder(&b)
	result := Person{}
	d.Decode(&result)
	return &result
}

type Address struct {
	StreetAddress, City, Country string
}
