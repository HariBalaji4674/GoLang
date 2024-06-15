package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("This is is some data to enter into a file")

	for i := 0; i < 10; i++ {
		data := fmt.Sprintf("")
	}

	filename := "data.txt"

	err := os.WriteFile(filename, data, 0777)

	if err != nil {
		fmt.Println("Error writing to the file : ", err)
		return
	}
	fmt.Println("Data has enetred into the file correctly")
}

/*
	os.write(filename,"input",permission)
	input for a file should be []byte(input variable /data to be entered)
	os.readfile(filename,permission)
	when you that it will give some value but we are not usin it hence we can save that as _
	nil is a special value in go like nil is empty value
	panic funtion when yo are not able to continue the code we will use panic
*/
