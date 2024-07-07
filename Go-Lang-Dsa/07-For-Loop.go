package main

import (
	"fmt"
)

func main() {
	for range "peddireddy" {
		fmt.Println("peddireddy")
	}

	i := 1

	for {
		fmt.Println("Hari Vardhan Reddy", i)
		if i == 10 {
			break
		}
		i++
	}
	k := 1
	for ; k <= 10; k++ {
		fmt.Printf("%v ", k)
	}

	for k := 1; ; k++ {
		fmt.Print(k)
		if k == 10 {
			break
		}
	}

	marks := []int{10, 20, 30, 40, 50, 60, 70}

	fmt.Println(marks)

	for i := range marks {
		fmt.Println(marks[i])
	}
	s := 1
	for k <= 11 {
		fmt.Println(s)
		
	}

	array := []int{10,20,30,40,50}
	fmt.Println(array)

}
