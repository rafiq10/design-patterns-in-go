package prototypefactory

// on demand objects created depending on an enum
// we could use interfaces instead of struct
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{Name: "", Position: "developer", AnnualIncome: 60000}
	case Manager:
		return &Employee{Name: "", Position: "manager", AnnualIncome: 80000}
	default:
		panic("usupported role")
	}
}
