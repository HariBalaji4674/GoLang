package main

import "fmt"

func main() {
	revString("peddireddy")
}

func revString(name string) {

	fmt.Println(name)
	name2 := ""
	fmt.Println(name2)
	for i := len(name) - 1; i >= 0; i-- {

		name2 = name2 + string(name[i])
		fmt.Print(string(name[i]))

	}

}

/*
Reverse a String:
- Give a string
- create a temp string
- traverse through the string using loop
-
*/
