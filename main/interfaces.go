package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rec struct {
	width, height int
}

func (r rec) area() float64 {
	return float64(r.width * r.height)
}

func (r rec) perimeter() float64 {
	return float64(2*r.width + 2*r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Sqrt(c.radius)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rec{10, 20}
	c := circle{10}

	measure(r)
	measure(c)
}
