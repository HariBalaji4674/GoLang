package main

import "fmt"

func main() {
	printNtimes(1, 10)
}

func printNtimes(i, n int) {

	if i > n {
		return
	}
	fmt.Println("Hari", i)
	printNtimes(i+1, n)
}
