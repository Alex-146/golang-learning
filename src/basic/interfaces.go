package main

import (
	"fmt"
	"math"
)

type IShape interface {
	area() float64
	toString() string
}

type Circle struct {
	radius float64
}

func (c Circle) toString() string {
	return fmt.Sprintf("Circle with radius %.2f", c.radius)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rect struct {
	width float64
	height float64
}

func (r Rect) toString() string {
	return fmt.Sprintf("Rect with sides %.2f;%.2f", r.width, r.height)
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func main() {
	c := Circle {5}
	r := Rect {3, 4}
	shapes := []IShape{c, r}
	for _, shape := range shapes {
		output := fmt.Sprintf("%v has area: %.2f", shape.toString(), shape.area())
		fmt.Println(output)
	}
}