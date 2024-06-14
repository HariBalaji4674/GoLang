package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (re rect) area() float64 {
	return re.height * re.width
}

func (ce circle) area() float64 {
	return ce.radius * ce.radius
}

func (re rect) perim() float64 {
	return 2 * re.height * re.width
}
func (ce circle) perim() float64 {
	return 2 * math.Pi * ce.radius * ce.radius
}

func printShapesInfo(s geometry) {
	fmt.Println(s.area())
	fmt.Println(s.perim())
}
func main2() {
	geometries := []geometry{
		rect{width: 20, height: 20},
		circle{radius: 5},
		// rect{width: 15, height: 15},
		// circle{radius: 4},
	}
	for _, v := range geometries {
		printShapesInfo(v)
		fmt.Println("-----")
	}

}
