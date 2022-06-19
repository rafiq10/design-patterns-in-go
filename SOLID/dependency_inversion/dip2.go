package dependencyinversion

import "fmt"

type ResearchBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// high-level
type Research2 struct {
	Browser ResearchBrowser
}

// Investigate depends now on an interface and not a concrete type
func (r *Research2) Investigate() {
	for _, v := range r.Browser.FindAllChildrenOf("Janusz") {
		fmt.Printf("John has a child: %s\n", v.Name)
	}
}

// Finding children is put in a low-level module now
// If we change for example []Info in Relationships to a db, Research type is not affected
func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.Relations {
		if v.Relationship == Parent &&
			v.From.Name == name {
			result = append(result, r.Relations[i].To)
		}
	}
	return result
}
