package main

import (
	"fmt"
)

func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}
	return res
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}
