package main

import (
	"fmt"
)

func main() {
	fmt.Println("Control Structures Methods")
	fmt.Println(add3(67, 23))
	loop(2000, 100)
}

func add3(num1, num2 int) int {
	return calcu(num1, num2)
}

func calcu(num1, num2 int) int {
	return num1 + num2
}

func loop(salary, age int) {
	for i := 0; i <= 10; i++ {
		fmt.Println("Please add some values : ")
		fmt.Scan(&salary, &age)
		fmt.Println(ready(i, i+1))
	}
}

func ready(num1, num2 int) int {
	return num1 + num2
}

/*
Conditions [ if else else-if nested-if]
Loops [for]

*/
