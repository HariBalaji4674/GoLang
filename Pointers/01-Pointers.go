package main

import "fmt"

func main(){
	age := 32
	fmt.Println(age)

	// var newAge *int 
	newAge := &age
	// newAge := &age
	fmt.Println(*newAge)


	// Using pointer and using them in functions

	aged := getAdultYears(*newAge)
	fmt.Println(aged)


}

func getAdultYears(age int) int{
	return age-18
}