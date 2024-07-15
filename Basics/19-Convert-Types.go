package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Function Conversion")
	strConv := "100"
	intVer,err := strconv.Atoi(strConv)

	fmt.Println(intVer,err)
}