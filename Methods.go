package main

import "fmt"

func main(){
	fmt.Println(rectangle.height)

}

type rect struct{
	height int
	width int
	length int
}

rectangle := rect {
	height : 20,
	width : 40,
	length : 50,
}

func (rectangle rect) area() int {
	return rectangle.width * rectangle.height*rectangle.length
}