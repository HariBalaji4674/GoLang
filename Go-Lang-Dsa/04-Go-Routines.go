package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main(){
	go responseUrl("https://www.golangprograms.com")
	go responseUrl("https://coderwall.com")
	go responseUrl("https://stackoverflow.com")

	time.Sleep(10*time.Second)
}

func responseUrl(url string){

	fmt.Println("Step 1 : ", url)
	response,err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Step 2 : ", url)
	defer response.Body.Close()

	fmt.Println("Step 3 : ", url)
	body,err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 4 : ", len(body))

}