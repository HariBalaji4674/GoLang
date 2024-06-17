package main

import "fmt"

func main() {
	fmt.Println("This is structs class")
	var user User = User{"Peddireddy",18,true,"Hari.abalaji4674"}

	fmt.Println(user)

	fmt.Printf("%+v\n",user)

	user1  := &user
	fmt.Println(*user1)
}

type User struct {
	Name string
	Age int
	Status bool
	Email string
}