package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	// Build the hash table
	for i, num := range nums {
		numMap[num] = i
	}

	// Find the complement
	for i, num := range nums {
		complement := target - num
		if val, ok := numMap[complement]; ok && val != i {
			return []int{i, val}
		}
	}

	return nil // No solution found
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if result != nil {
		fmt.Println("Indices of the two numbers that add up to", target, "are:", result)
	} else {
		fmt.Println("No two numbers found in the array that add up to", target)
	}
}
