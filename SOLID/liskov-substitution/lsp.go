package liskovsubstitution

import "fmt"

type Sized interface {
	GetWith() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func (r *Square) GetWidth() int {
	return r.width
}

func (r *Square) SetWidth(width int) {
	r.width = width
	r.height = width
}

func (r *Square) GetHeight() int {
	return r.height
}

func (r *Square) SetHeight(height int) {
	r.width = height
	r.height = height
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// This could solve the problem
type Square2 struct {
	size int // both height and width
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{height: s.size, width: s.size}
}

// It should work even if we extend the original object (ie: from rectangle to square)
func UseIt(sized Sized) {
	width := sized.GetWith()
	// sized.SetHeight(10) sets both height and with for square afer extending the rectangle object
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWith() * sized.GetHeight()
	fmt.Printf("Expected area: %d, Got area: %d", expectedArea, actualArea)
}
