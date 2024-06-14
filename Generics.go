package main

import "fmt"

func main() {
	fmt.Println("Peddireddy ")
	result := add(20.3,30.5)
	fmt.Println(result)
    fmt.Scan()
}

// func add(a int , b int) int {
// 	return a+b
// }

func add[T num] (a T, b T) T{
	return a+b
}

type num interface{
	int | uint | uint16 | uint64 | float32 | float64
}