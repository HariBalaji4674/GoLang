package main

import "fmt"

func main() {

	age := 32
	fmt.Println(age)

	var newAge *int
	newAge = &age      // referencing the Address of Pointer to newAge
	fmt.Println("Address of Variable : ", newAge) // reference/Address Value Printing in screen

	fmt.Println("Deferencing Value : ", *newAge) // Dereferencing the Operator 
	// fmt.Println("After Changing value : ", getAdultAge(age))
	getAdultage1(newAge)
	fmt.Println("Modified Value : ", age)

}

// func getAdultAge(age int) int {
// 	return age - 18
// }

func getAdultage1(age *int) {
	*age = *age + 18
}
