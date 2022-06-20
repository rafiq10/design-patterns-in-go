package factoryfunction

type Person struct {
	Name     string
	Age      int
	EyeCount int // normally people have 2 eyes
}

// a simple factory function that simply sets the EyeCount to a default value
// it can also do some validation
func NewPerson(name string, age int) *Person {
	if age < 16 {
		// ...
	}
	return &Person{Name: name, Age: age, EyeCount: 2}
}
