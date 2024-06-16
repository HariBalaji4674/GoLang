package main

import "fmt"

func main(){

	var u user = user{"peddireddy","hari vardhan reddy",15,67.89,true}
	fmt.Println(u)
	u.age = 25
	agedet(&u)
}

type user struct{
	firstname string
	secondname string
	age int
	marks float64
	status bool
}

func agedet(u *user) int{
	if u.age > 18{
		fmt.Println("eligible to vote")
	}else{
		value := 18 - u.age 
		fmt.Printf(" Retry After %v Years \n",value)
	}
	return u.age
}
