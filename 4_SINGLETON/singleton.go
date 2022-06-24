package main

import (
	"fmt"
	s "singleton/singleton"
	"strconv"
)

func main() {
	sg := s.GetSingletonDB()
	population := sg.GetPopulation("Seoul")
	fmt.Println("Seoul population: " + strconv.Itoa(population))
}
