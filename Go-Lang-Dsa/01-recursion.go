package main

import "fmt"

func main() {
	fmt.Println("Peddireddy Hari Vardhan Reddy")

	b := add1(10)

	fmt.Println(b)

	repeat(10)
	repeat1(10)

	n := []int{10, 20, 30, 40, 50}
	a := 0
	mean(n, a)
	fmt.Println(natural(10))
}

func add1(a int) int {

	if a == 0 {
		return 0
	}
	return a + add1(a-1)

}
func repeat(n int) {
	if n > 0 {
		repeat(n - 1)
		fmt.Println(n)
	}
}

func repeat1(n int) {
	if n > 0 {
		fmt.Println(n)
		repeat1(n - 1)
	}
}

func mean(n []int, a int) {
	length := len(n)
	for i := range n {
		a = a + n[i]
	}
	fmt.Println(a)
	meanArr := a / length
	fmt.Println(meanArr)

}

func natural(n int) int {
	// sum := 0
	// if n == 0 {
	// 	return 1
	// }
	// return sum + natural(n-1)

	return (n * (n + 1))/2

}
