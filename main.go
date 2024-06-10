package main

import "fmt"

func main1() {

	var associate employee_details

	associate.employeeName = "Peddireddy Hari Vardhan Reddy"
	associate.address = "chiyapaduVillage"
	associate.grade = 23
	associate.rollno = 9493834674
	associate.salary = 1000

	fmt.Println(associate)

	fmt.Println(associate.salary)
	fmt.Println(associate.employeeName)
	fmt.Println(associate.rollno)
	fmt.Println(associate.grade)

	fmt.Println("Peddireddy")
	var own messageTosend
	own.message = "Please Use My Declaration"
	own.number = 9347998315
	fmt.Println(own.message)
	fmt.Println(own.number)

	var twoown details
	twoown.become = true
	twoown.idea = "Killing Dream"
	twoown.name = "SImplicity"
	twoown.village = "Chiyyapadu"
	twoown.pincode = 516355
	twoown.marks = 675.90

	fmt.Println(twoown)

	test(messageTosend{number: 9493834674, message: "Thanks For coming"})
}

func test(m messageTosend) {
	fmt.Println(m.message, m.number)
}

type messageTosend struct {
	number  int
	message string
}
type details struct {
	idea    string
	name    string
	village string
	marks   float64
	pincode int
	become  bool
}

type employee_details struct {
	employeeName string
	salary       int
	rollno       int
	address      string
	grade        float64
}
