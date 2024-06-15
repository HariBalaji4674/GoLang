package main

import "fmt"

func main() {
	fmt.Println("Calling Several Methods")
	fmt.Println(add1(10,20,30,40))

}
func add1(num1,num2,num3,num4 int) int{
	if num1 > num2{
		return num3
	}else {
		return num4
	}
}

func sub(num1,num2,num3 int) int {
	if num1>num2{
		return num1-num2
	}else{
		return num3
	}
}