package adapter

import (
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

// this is the interface we are given
func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{
		[]Line{
			{0, 0, width, 0},
			{0, 0, 0, height},
			{width, 0, width, height},
			{0, height, width, height}}}
}

// the interface we have
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

// as DrawPoints uses RasterImage interface, we need an adapter that adapts
// VectorImage struct to RasterImage interface
func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}
	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}
	return b.String()
}

// SOLUTION:
// an adapter that converts vector to rasterImage
type VectorToRasterAdapter struct {
	points []Point
}

func (v VectorToRasterAdapter) GetPoints() []Point {
	return v.points
}
func (v *VectorToRasterAdapter) AddLine(line Line) {
	// find edges of the line
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := top - bottom
	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}
	fmt.Println("generated", len(v.points), "points")
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := VectorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.AddLine(line)
	}
	return &adapter
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
