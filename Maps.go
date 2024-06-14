package main

import "fmt"

func main9(){
	m := make(map[string]int)
	fmt.Println(m)
	m["name"] = 123
	m["age"] = 345
	m["ret"] = 56
	fmt.Println(m)

	n := map[string]int{"name": 1234,"age":23,"area":24}
	fmt.Println(n)

	n1 := map[string]string{"name":"peddireddy","age":"24","simple":"pedd","duty":"rowdy"}
	fmt.Println(n1)

	// delete particular element in map
	delete(m,"name")
	fmt.Println(m)

	// clear method will clear elements in map
	clear(m)
	fmt.Println(m)

	j := map[string]int{"maths":76,"science":87,"social":45}
	fmt.Println(j)

	_,marks := j["social"]
	fmt.Println(marks)

	k := []int{20,30,40,50,60}
	fmt.Println(k)

	for k1,v1 := range k {
		fmt.Println(k1,v1)
	}
	
}

