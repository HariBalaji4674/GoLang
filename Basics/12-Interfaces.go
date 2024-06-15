package main

import "fmt"

// declare interface
type methods interface {
	method1() float64
	method2() float64
}

// Declare structs
type method1_vars struct {
	length, height float64
}
type method2_vars struct {
	breadth float64
}

// initialize Methods
func (mv1 method1_vars) method1() float64 {
	return mv1.height * mv1.length
}
func (mv2 method2_vars) method2() float64 {
	return mv2.breadth
}
func (mv1 method1_vars) method2() float64 {
	return 2 * mv1.height * mv1.length
}
func (mv2 method2_vars) method1() float64 {
	return 2 * mv2.breadth
}
func applyInter(ms methods) {
	fmt.Println(ms.method1())
	fmt.Println(ms.method2())
}

func main7() {
	r := method1_vars{length: 20, height: 15}
	c := method2_vars{breadth: 12}
	applyInter(r)
	applyInter(c)
}

// Interface are used to decklare two method signatures
//
