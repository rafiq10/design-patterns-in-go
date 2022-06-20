package main

import (
	ff "factory/factory_function"
	intf "factory/interface_factory"
	"fmt"
)

func main() {
	// factory function

	p := ff.NewPerson("Rafal", 35)
	// we can always change the default value afterwards
	p.EyeCount = 1
	fmt.Println(p)

	// interface factory

	p2 := intf.NewPerson("Rafal", 71)
	p2.SayHello()
}
