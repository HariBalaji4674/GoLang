package main

import (
	"fmt"
	"strconv"
)

func palindrome() {
	fmt.Println("This is Palindrome Function")
	fmt.Println(palindrom(121))
	fmt.Println(palindrom(13231))
	fmt.Println(palindrom(133221))
}

func palindrom(num int) bool {
	s := strconv.Itoa(num)
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--

	}
	return true

}
