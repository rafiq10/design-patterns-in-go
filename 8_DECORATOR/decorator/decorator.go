package decorator

import "fmt"

type Aged interface {
	// getter and setter
	Age() int
	SetAge(age int)
}
type Bird struct {
	age int
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying")
	}
}
func (b *Bird) Age() int {
	return b.age
}
func (b *Bird) SetAge(age int) {
	b.age = age
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int {
	return l.age
}
func (l *Lizard) SetAge(age int) {
	l.age = age
}

func (l *Lizard) Crawl() {
	if l.age >= 5 {
		fmt.Println("Crawling")
	}
}

// If we want Dragon to inritate from both bird and lizard we could do the following
// However both bird and lizard have the Age field
// type Dragon struct {
// 	Bird
// 	Lizard
// }

// this dragon type is meant to extend the behaviors of bird and lizard
// effectively it simply provides a better access to the fields in a consistent way
type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) Age() int {
	return d.bird.age
}
func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}
func (d *Dragon) Fly() {
	d.bird.Fly()
}
func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

// factory function to create a new dragon and initialize it
func NewDragon() *Dragon {
	return &Dragon{bird: Bird{}, lizard: Lizard{}}
}
