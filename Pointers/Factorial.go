package main

import "fmt"

func main() {
	fmt.Println("Peddireddy Hari Vardhan reddy")
	fmt.Println("Factorial of number is called : ", fact(10))
}

var factVal uint64 = 1
var n int

func fact(n int) uint64 {
	if n < 0 {
		fmt.Println("Factorial Of Negative Number doesnot exist : ")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}
	}
	return factVal
}
