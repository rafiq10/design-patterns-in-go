package main

import (
	cc "composite/composite"
	nn "composite/neuron"
	"fmt"
)

func main() {
	drawing := cc.GraphicObject{Name: "Drawing", Color: "", Children: nil}
	drawing.Children = append(drawing.Children, *cc.NewCircle("Red"))
	drawing.Children = append(drawing.Children, *cc.NewSquare("Blue"))

	group := cc.GraphicObject{Name: "Group", Color: "", Children: nil}
	group.Children = append(group.Children, *cc.NewCircle("Green"))
	group.Children = append(group.Children, *cc.NewSquare("Green"))
	drawing.Children = append(drawing.Children, group)
	fmt.Println(drawing.String())

	// NEURONS
	neuron1, neuron2 := &nn.Neuron{}, &nn.Neuron{}
	layer1, layer2 := nn.NewNeuronLayer(2), nn.NewNeuronLayer(3)
	nn.Connect(neuron1, neuron2)
	nn.Connect(neuron1, layer1)
	nn.Connect(layer2, neuron1)
	nn.Connect(layer1, layer2)

}
