package liskovsubstitution

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) GetWidth() int {
	return r.Width
}

func (r *Rectangle) SetWidth(width int) {
	r.Width = width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) SetHeight(height int) {
	r.Height = height
}

type Square struct {
	Rectangle
}

func (r *Square) GetWidth() int {
	return r.Width
}

func (r *Square) SetWidth(width int) {
	r.Width = width
	r.Height = width
}

func (r *Square) GetHeight() int {
	return r.Height
}

func (r *Square) SetHeight(height int) {
	r.Width = height
	r.Height = height
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.Width = size
	sq.Height = size
	return &sq
}

// This could solve the problem
type Square2 struct {
	Size int // both height and width
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{Height: s.Size, Width: s.Size}
}

// It should work even if we extend the original object (ie: from rectangle to square)
func UseIt(sized Sized) {
	width := sized.GetWidth()
	// sized.SetHeight(10) sets both height and with for square afer extending the rectangle object
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Printf("Expected area: %d, Got area: %d\n", expectedArea, actualArea)
}
