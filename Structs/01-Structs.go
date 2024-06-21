package Structs

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstname string
	lastname  string
	birthdate string
	createdAt time.Time
}

// methods in go called adding function to the struct

func (u user) method1(village string) {
	fmt.Println("This is a method 1")
	fmt.Println("Method 1: ", u.birthdate)
	fmt.Println("Method 1: ", u.firstname)
	fmt.Println("Method 1: ", u.lastname)
	fmt.Println("Method 1: ", u.createdAt)

	fmt.Println("Calling Village : ", village)

}

func (u *user) clearMethod() {
	fmt.Println("Clearing All These : ", u.birthdate, u.lastname, u.firstname)
	u.firstname = " "
	u.lastname = " "
	u.birthdate = " "

}
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func outputUserDetails(firstName, lastName, birthDate string) {
	fmt.Println(firstName, lastName, birthDate)
}

func outputStruct(u user) {
	fmt.Println(u.createdAt, u.birthdate, u.firstname, u.lastname)
}

// Passing Pointers to the function as struct

func pointStruct(u *user) {
	fmt.Println(u.birthdate, u.lastname, u.firstname, u.createdAt)
}

// Constructer Functions

func newUser(firstName, lastName, birthDate string) (*user, error) {

	if firstName == "" || lastName == "" || birthDate == "" {

		return nil, errors.New("This is errore message required ")

	}
	return &user{
		firstname: firstName,
		lastname:  lastName,
		birthdate: birthDate,
		createdAt: time.Now(),
	}, nil
}

/*

Only Capital Letters start only will work outside the package
Structs Embedding : builds new struct on top of existing Structs
Structrs inside the structs is called embedded structs

Loops
Structs
Methods
Packages
Variables


Creating Custom Constructors


*/
