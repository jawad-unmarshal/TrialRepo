package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length  float64
	breadth float64
}

func (rectangle rect) area() float64 {
	return (rectangle.length * rectangle.breadth)

}
func (rectangle rect) perimeter() float64 {
	return 2. * (rectangle.breadth + rectangle.length)
}

type Sqr struct {
	length float64
}

func (rectangle Sqr) area() float64 {
	return (rectangle.length * rectangle.length)

}
func (rectangle Sqr) perimeter() float64 {
	return 4. * (rectangle.length)
}

type Cir struct {
	radius float64
}

func (rectangle Cir) area() float64 {
	return math.Pi * rectangle.radius * rectangle.radius

}
func (Circle Cir) diameter() float64 {
	return Circle.radius*2
}
func (rectangle Cir) perimeter() float64 {
	return math.Pi*rectangle.diameter()
}

func main() {
	var s shape
	s = Cir{radius: 45}
	fmt.Println(s.area())
	fmt.Println("Diameter: ",s.(Cir).diameter())
}
