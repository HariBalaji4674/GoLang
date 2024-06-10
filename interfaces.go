// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main(){

// 	fmt.Println("Peddireddy Hari Vardhan Reddy")
// 	var re rect
// 	re.length = 20
// 	re.width = 30
// 	fmt.Println(re.area())

// 	r := rectan{length:4, breadth: 6}
// 	c := cicle{radius: 10}

// 	measure(r)
// 	measure(c)

// }

// type rect struct {
// 	width float64
// 	length float64
// }

// func (r rect) area() float64 {
// 	return 2*r.length + 2*r.width
// }

// // Interface is an Named collection of method signatures

// type geometry interface {
// 	area()
// 	rectangle()
// }

// type rectan struct{
// 	length,breadth float64
// 	// breadth float64
// }

// type cicle struct{
// 	radius float64
// }

// func (r rectan) rectangle() float64{
// 	return 2*r.length*r.breadth
// }

// func (c cicle) area() {
// 	return math.Pi*c.radius*c.radius
// }

// func measure(g geometry){
// 	fmt.Println(g)
// 	fmt.Println(g.area())
// 	fmt.Println(g.rectangle())
// }