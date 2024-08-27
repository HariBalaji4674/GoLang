package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	emptyFile, err := os.Create("empty.file")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
	myFunc()
	add(10, 20)
	sum := add4(20,30)
	fmt.Println(sum)


}

func myFunc() {
	fmt.Println("This is main Function")
}

func add(x, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

func add4(a, b int) int {
	total := 0
	total = a + b
	return total
}
