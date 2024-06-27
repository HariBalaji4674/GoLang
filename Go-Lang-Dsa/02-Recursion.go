package main

import "fmt"

func main() {
	printName(0)
	Printname(10)
	fmt.Println(naturalNumbers(6))
}

func printName(count int) {
	// count := 0
	if count == 10 {
		return
	}
	fmt.Println("Peddireddy")
	count = count + 1
	printName(count)
}

func Printname(n int) {
	if n > 0 {
		Printname(n - 1)
		fmt.Println("Peddireddy %d", n)
	}
}

func naturalNumbers(n int) int {

	if n == 0 {
		return 0
	}
	return n + naturalNumbers(n-1)
}

func isPalindrome(n int) {

}
