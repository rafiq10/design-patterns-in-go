package main

import (
	ad "adapter/adapter"
	"fmt"
)

func main() {
	rc := ad.NewRectangle(6, 4)
	a := ad.VectorToRaster(rc)
	fmt.Println(ad.DrawPoints(a))
}
