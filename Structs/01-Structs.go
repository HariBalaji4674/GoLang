package main

import "fmt"

func main() {

	UserDetails := User{}
	fmt.Println(UserDetails)

	StrInitial(UserDetails)

	NewUser(UserDetails)

}

type User struct {
	Name   string
	Age    int
	Marks  float64
	Status bool
}

func NewUser(u User) {
	fmt.Println(u.Age, u.Marks, u.Name, u.Status)
}

func StrInitial(u User) {
	var username string
	fmt.Println("Please Enter UserName : ")
	fmt.Scan(&username)
	fmt.Println(username)

	u.Name = username

}

// Structs Declaration
// Structs Calling or Instantiating
// Struct declaring
