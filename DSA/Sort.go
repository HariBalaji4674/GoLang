package main

import "fmt"

func SortArray(){
	array := []int{4,5,3,2,9,8,5}
	fmt.Println(array)
	reverseSort(array)
}

func reverseSort(arr []int){
	n := len(arr)
	arr2 := []int{}
	fmt.Println(arr2)
	max := arr[0]
	for i := 0 ; i < n-1 ;i++{
		if arr[i] < max{
			arr2 = append(arr2, arr[i])
		}
		// if arr[i+1] < arr[i]{
		// 	arr2 = append(arr2, arr[i])
		// }
	}
	fmt.Println(arr2)
}
