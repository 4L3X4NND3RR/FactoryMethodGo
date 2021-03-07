package main

import (
	"fmt"

	"github.com/4L3X4NND3RR/FactoryPattern/FigureFactory"
)

func main() {
	square := FigureFactory.NewSquare(10)
	fmt.Printf("Square: %f \n", square.Area())
	rectangle := FigureFactory.NewRectangle(30, 20)
	fmt.Printf("Rectangle: %f \n", rectangle.Area())
	triangle := FigureFactory.NewTriangle(30, 45, 50)
	fmt.Printf("Triangle: %f", triangle.Area())
}
