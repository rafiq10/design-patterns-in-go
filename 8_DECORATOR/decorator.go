package main

import (
	dd "decorator/decorator"
)

func main() {

	// d := dd.Dragon{}
	// this would be an ambiguous selector (both bird and lizard have Age field):
	// d.Age = 10

	// This would be inconsistent:
	// d.Bird.Age = 10
	// d.Lizard.Age = 5

	d := dd.NewDragon()
	d.SetAge(11)
	d.Fly()
	d.Crawl()
}
