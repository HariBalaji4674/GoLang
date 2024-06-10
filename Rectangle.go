package main

import (
	"fmt"
	"math"
)

func main() {

	r := rect{width: 5,height: 7}

	measure(r)

	c := circle{radius: 5}
	measure(c)

}

type rect struct{
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64{
	return 2*r.width*r.height
}

func (c circle) area() float64{
	return 2*math.Pi*c.radius*c.radius
}

type geometry interface{
	area() float64
}

func measure(g geometry) {
fmt.Println(g)
fmt.Println(g.area())
}

