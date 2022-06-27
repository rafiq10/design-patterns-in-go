package main

import (
	br "bridge/bridge"
)

func main() {
	raster := br.RasterRenderer{}
	vector := br.VectorRenderer{}
	circle := br.NewCircle(&raster, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()

	circle2 := br.NewCircle(&vector, 5)
	circle2.Draw()
	circle2.Resize(2)
	circle2.Draw()
}
