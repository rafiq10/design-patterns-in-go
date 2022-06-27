package bridge

import "fmt"

// circle, square
// raster, vector

// RasterCircle, RasterSquare, VectorCircle, VecorSquare

type Renderer interface {
	RenderCirlce(radius float32)
}

type VectorRenderer struct {
	//
}

func (v *VectorRenderer) RenderCirlce(radius float32) {
	fmt.Println("Drawing a circle of radius ", radius)
}

type RasterRenderer struct {
	DPI int
}

func (v *RasterRenderer) RenderCirlce(radius float32) {
	fmt.Println("Drawing pixels for a circle of radius ", radius)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() {
	// here is where we save double code for Raseter and Vector as we use the interface 'renderer'
	c.renderer.RenderCirlce(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}
