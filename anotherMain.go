package main

import (
	"fmt"
	"os"
)

func main() {

	fileread()
}

func fileread() {
	file, err := os.Open("main.txt")
	if err != nil {
		fmt.Println("Returns err : ", err)
		return
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			fmt.Print(string(buffer[:n]))
		}
		if err != nil {
			break
		}
	}
}
