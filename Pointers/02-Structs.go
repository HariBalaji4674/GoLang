package main

import "fmt"

func main(){

	fmt.Println("Pelase enter the value: ")
	var age int
	fmt.Scan(&age)
	fmt.Println(age)

	ages := fmt.Sprint("This is a String")
	fmt.Println(ages)

	fmt.Scanf()

}

func getUserdata() {

	fmt.Println(" This is Secondary Method ")

}
