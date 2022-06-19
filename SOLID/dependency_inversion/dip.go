package dependencyinversion

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	Name string
	//
}

type Info struct {
	From         *Person
	Relationship Relationship
	To           *Person
}

// low-level; ie storage mechanism
type Relationships struct {
	Relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.Relations = append(r.Relations, Info{parent, Parent, child})
	//r.Relations = append(r.Relations, Info{child, Child, parent})
}

// high-level
type Research struct {
	//this breaks DIP as high-level depends on low-level object
	Relationships Relationships
}

func (r *Research) Investigate() {
	relations := r.Relationships.Relations
	for _, rel := range relations {
		if rel.From.Name == "Janusz" &&
			rel.Relationship == Parent {
			fmt.Printf("John has a child: %s\n", rel.To.Name)
		}
	}
}
