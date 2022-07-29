package main

import (
	pp "design-patterns/proxy/protection-proxy"
)

func main(){
	car := pp.NewCarProxy(&pp.Driver{Age: 22})
	car.Drive()
	car2 := pp.NewCarProxy(&pp.Driver{Age: 12})
	car2.Drive()
}