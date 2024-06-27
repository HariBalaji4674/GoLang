package main

import "fmt"

func main() {
	i := 0
	n := 10
	reverseNumbers(i, n)

}

func reverseNumbers(i, n int) {
	if i < 1 {
		return
	}
	fmt.Println(i)
	reverseNumbers(i-1, n)
}
