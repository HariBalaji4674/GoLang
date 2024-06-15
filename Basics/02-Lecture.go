package main

import "fmt"

func main() {

	fmt.Println("peddireddy Hari vardhan reddy")
	name := "peddireddy"
	fmt.Scan(&name)
	fmt.Println("Thats cool ", name)

	var revenue float64
	var expenses float64
	var tax_rate float64

	fmt.Print("Please eneter revenue : ")
	fmt.Scan(&revenue)

	fmt.Print("Please eneter expenses : ")
	fmt.Scan(&expenses)

	fmt.Print("Please eneter tax_rate : ")
	fmt.Scan(&tax_rate)

	fmt.Println(revenue)
	fmt.Println(expenses)
	fmt.Println(tax_rate)

	var ebt = revenue - expenses
	profit := ebt * (1 - tax_rate/100)
	fmt.Println(ebt)
	fmt.Println(profit)

	//%v is represented as placeholder for value of an variable or constant
	fmt.Printf("%v \n", 286)

	fmt.Printf("%.2f\n", 2345.46787458)
	fmt.Println("%T", 456547)

	// when to print value later Sprintf
	var valueVar float64 = 467.09877
	formattedvalue := fmt.Sprintf("Entered Value %.1f ", valueVar)
	fmt.Println(formattedvalue)

	fmt.Println(`Peddireddy Hari Vardhan Reddy
	this is line number 2 
	this is line number 3
	this is line number 4
	`)

	fmt.Println(method1(30, 40))
	fmt.Println(method2("pedddireddy", "hari vardhan reddy"))
	fmt.Println(method3(456.89, 987.564))

	/*
		Backtick is used to write multiple lines in a single printf
		Functions : Sequence of steps to solve a probelm
		we can add variable and constants outside the functions as well

		Go is a static typed language we need to specify the return type particularly


	*/
}

func method1(number1, number2 int) int {
	return number1 + number2
}

func method2(name1, name2 string) string {
	return name1 + name2
}
func method3(marks1, marks2 float64) float64 {
	return marks1 + marks2
}
