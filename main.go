package main

import (
	"fmt"
	"os"
	"time"

	// "example.com/practice/practice"
)

func main() {

	file, err := os.Open("main.txt")

	if err != nil {
		fmt.Println("Error list:", err)
		return
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			fmt.Println(string(buffer[:n]))
		}
		if err != nil {
			break
		}
	}iu90h
	done := make(chan bool)
	go add("peddireddy", done)
	<- done                                                          
}

func add(text string, goChan chan bool){
	time.Sleep(3*time.Second)
	fmt.Println("peddireddy", text)
	goChan <- true
}


	
	
