package main

import "fmt"

func main3() {
	// fmt.Println(rectangle.height)
	fmt.Println("Peddireddy")

}

type rect1 struct {
	height int
	width  int
	length int
}

func (rectangle rect1) area() int {
	return rectangle.width * rectangle.height * rectangle.length
}

// Structs are the collection types
