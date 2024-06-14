package main

import "fmt"

func main10() {
	fmt.Println("Peddireddy Hari Vardhan Reddy")

	emp := HR{
		Author:     Author{Name: "Alice", Department: "Engineering", Year: 2024},
		EmployeeID: "12345",
	}

	fmt.Println("Details Of Author ")
	fmt.Println(emp.Name, emp.Department, emp.Year)
	fmt.Println("EmployeeID : ", emp.EmployeeID)
}

type Author struct {
	Name       string
	Department string
	Year       int
}

type HR struct {
	Author
	EmployeeID string
}

// Anonomous Struct
// Spoiler Alert
// embedded structs : Shower them into single struct
// Embedded struct is a Struct with in a struct
// Methods in GO Programming Language
// behavious and Finding on their own
//
