package main

import "fmt"

func main() {
	fmt.Println(" This is methods Class ")
	var Details AddUser = AddUser{"Peddireddy", 24, 2024, true}
	fmt.Println(Details)

	// Detailed()
	fmt.Println(Details.Detailed())

}

// Adding functions to structs is called the Methods
// Receiver arguments should mention in before function name 
// 

type AddUser struct {
	Name   string
	Age    int
	Year   int
	Status bool
}

func (u AddUser) Detailed() (int, string, bool, int) {
	return u.Age, u.Name, u.Status, u.Year
}

func newAddUser() AddUser {
	
	return AddUser{
		

	}  
}




