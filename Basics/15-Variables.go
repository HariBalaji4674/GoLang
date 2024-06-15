package main

import (
	"fmt"
	"time"
)

func main4() {
	fmt.Println("Peddireddy")

	var a = "Initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)

	var e,g int = 10,20
	fmt.Println(e,g)
	g,e = e,g
	fmt.Println(e,g)

	var num,str  = 20,"Simple Language"
	fmt.Println(num,str)

	// Go supports Constants values for strings numbers boolean etc ..
	const number = 25
	// number = 40
	fmt.Println(number)

	// constaant cannot modify once it is declared
	// constant will be used when there is unchanged vlaues to be declared

	// Loops for while nested for loop

	i := 1
	for i <= 3{
		fmt.Println(i)
		i = i+1
	}
	for j := 0; j<5;j++ {
		fmt.Println(j)
	}

	for i := range 4 {
		fmt.Println("Range",i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 10 {
		if n%2 == 0 {
			// continue
			fmt.Println(n)
		}
		// fmt.Println(n)
	}

	num1 := 20
	if num1 >= 10 {
		fmt.Println(num1)
	} else {
		fmt.Println(num1+1)
	}

	min := 10
	medi := 15
	max := 20

	fmt.Println(min)
	fmt.Println(medi)
	fmt.Println(max)

	if num := 9; num < 0 {
		fmt.Println("Either 8 or 7 are even")
	} else if num < 10 {
		fmt.Println(num, "Has 1 digit")
	} else {
		fmt.Println(num,"Has multiple digits")
	}

	// && || operators are used when two or more digits to be evaluated
	//  switch case statement
	m := 2

	fmt.Print("Write  ", m , "  as")
	switch m {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}


	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("Its WeekEnd Totally")
	default:
		fmt.Println("Its weekday totally")
	}

	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("Its Before Noon")
	default:
		fmt.Println("Its After Noon")
	}

	whoAmi := func(i interface{}){
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a Bool")
		case int:
		  fmt.Println("I am an Int")
		default:
		  fmt.Printf("Dont Know type %T \n", t)
	}
}
whoAmi(true)
whoAmi(1)
whoAmi("hey")
}