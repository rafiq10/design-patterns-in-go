package main

import (
	ad "adapter/adapter"
	"fmt"
)

func main() {
	rc := ad.NewRectangle(6, 4)
	a := ad.VectorToRaster(rc)
	b := ad.VectorToRasterCached(rc) // second adapter - it doesn't add the points as they are already in cache
	fmt.Println(ad.DrawPoints(a))
	fmt.Println(ad.DrawPoints(b))
}
