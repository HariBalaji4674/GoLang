package main

import "fmt"

func main() {
	age := 10
	fmt.Println("Age : ", age)

	var newAge *int
	newAge = &age
	fmt.Println("Address/Reference : ", newAge)
	fmt.Println("Dereference : ", *newAge)

	getAge(newAge)
	// fmt.Println(getAge(age))
	fmt.Println(age)
}

func getAge(age *int) {
 *age = *age + 18
}
