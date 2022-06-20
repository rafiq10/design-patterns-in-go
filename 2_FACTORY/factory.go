package main

import (
	ff "factory/factory_function"
	"fmt"
)

func main() {
	// factory function

	p := ff.NewPerson("Rafal", 35)
	// we can always change the default value afterwards
	p.EyeCount = 1
	fmt.Println(p)
}
