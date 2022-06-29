package main

import (
	dd "decorator/decorator"
	ddd "decorator/decorator_2"
	"fmt"
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

	// decorator2
	circle := ddd.Circle{Radius: 2}
	fmt.Println(circle.Render())
	square := ddd.Square{Side: 5}
	fmt.Println(square.Render())

	redCircle := ddd.ColoredShape{Shape: &circle, Color: "Red"}
	fmt.Println(redCircle.Render())
	greenSquare := ddd.ColoredShape{Shape: &square, Color: "Greeen"}
	fmt.Println(greenSquare.Render())

	// the problem is that we lose circle.Resize() method in the redCircle struct
	// we cannot call redCircle.Resize()
	// we cannot add Resize() to the Shape inerface as Square doesn't have it
	// it is necessary to create a new decorator

	// Decorators can be composed
	rhsCircle := ddd.TransparentShape{Shape: &redCircle, Transparency: 0.5}
	fmt.Println(rhsCircle.Render())
}
