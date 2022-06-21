package openclosed

type Color int

const (
	Red Color = iota
	Green
	Blue
)

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type Product struct {
	Name  string
	Color Color
	Size  Size
}

type Filter struct {
	//
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	Color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.Color == c.Color
}

type SizeSpecification struct {
	Size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size == s.Size
}

// With BetterFilter you are unlkiely to modify the Specification
// We don't want to add methods to the type once deployed
type BetterFilter struct{}

func (b *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.Color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// Extended Functionality
// This is adding additional functions to the type
func (f *Filter) FilterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.Size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

type AndSpecification struct {
	FirstSpec, SecondSpec Specification
}

func (s *AndSpecification) IsSatisfied(p *Product) bool {
	return s.FirstSpec.IsSatisfied(p) && s.SecondSpec.IsSatisfied(p)
}
