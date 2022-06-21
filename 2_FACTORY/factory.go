package main

import (
	ff "factory/factory_function"
	fgen "factory/factory_generator"
	intf "factory/interface_factory"
	pf "factory/prototype_factory"
	"fmt"
)

func main() {
	// 3.1 factory function

	p := ff.NewPerson("Rafal", 35)
	// we can always change the default value afterwards
	p.EyeCount = 1
	fmt.Println(p)

	// 3.2 interface factory
	p2 := intf.NewPerson("Rafal", 71)
	p2.SayHello()

	// 3.3 factory generator
	// functional approach
	devFactory := fgen.NewEmployeeFactory("developer", 60000)
	mgerFactory := fgen.NewEmployeeFactory("manager", 80000)

	developer := devFactory("Rafal")
	manager := mgerFactory("Olek")
	fmt.Println(developer, manager)

	// structural approach
	// in this approach we can change factory fields of 'position' and 'annual income'
	// whereas in the functional approach we cannot alter them afterwards
	bossFactory := fgen.NewEmployeeFactory2("Boss", 100000)
	boss := bossFactory.Create("Andrzej")
	fmt.Println(boss)

	bossFactory.AnnualIncome = 120000
	boss = bossFactory.Create("Andrzej")
	fmt.Println(boss)

	// 3.4 Prototype Factory
	d := pf.NewEmployee(pf.Developer)
	d.Name = "Rafal"
	fmt.Println(d)
}
