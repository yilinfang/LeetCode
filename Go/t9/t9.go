package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	a, b := x, 0
	for x > 0 {
		b = 10*b + x%10
		x /= 10
	}
	return a == b
}

func main() {
	fmt.Println(isPalindrome(-1))
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(11))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(1222))
}
