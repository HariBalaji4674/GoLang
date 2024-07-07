package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)

	fmt.Println("Staring go routines ")

	go responce("https://www.golangprograms.com")
	go responce("https://coderwall.com")
	go responce("https://stackoverflow.com")

	wg.Wait()
	fmt.Println("Terminating the services")
}

func responce(url string) {

	defer wg.Done() // Used to decrement the counter

	fmt.Println("Step 1: ", url)
	responce, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 2 : ", url)
	defer responce.Body.Close()

	fmt.Println("Step 3 : ", url)
	body, err := io.ReadAll(responce.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 4 : ", len(body))
}
