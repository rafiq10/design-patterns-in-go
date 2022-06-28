package main

import (
	cc "composite/composite"
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
}
