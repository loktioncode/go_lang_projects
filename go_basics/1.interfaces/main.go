package main

import (
	"fmt"
	"math"
)

//define interface
type Shape interface {
	area() float64
}

//structs
type Circle struct {
	x, y, rad float64
}

type Square struct {
	width, height float64
}

//value pointer
func (c Circle) area() float64 {
	return math.Pi * c.rad * c.rad
}

func (s Square) area() float64 {
	return s.width * s.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, rad: 5}
	square := Square{width: 5, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("SQUARE Area: %f\n", getArea(square))
}
