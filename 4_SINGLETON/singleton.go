package main

import (
	"fmt"
	"singleton/singleton"
	s "singleton/singleton"
	"strconv"
)

func main() {
	sg := s.GetSingletonDB()
	population := sg.GetPopulation("Seoul")
	fmt.Println("Seoul population: " + strconv.Itoa(population))

	// problems with singleton
	// the problem is that if we want to test it we should be able to mock the database
	cities := []string{"Seoul", "Mexico City"}
	tp := singleton.GetTotalPopulation(cities)

	ok := tp == 34900000
	fmt.Println(ok)
}
