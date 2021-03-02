package main

import (
	"fmt"
)

func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		a, b := nums[i], nums[i+1]
		if a > b {
			cnt++
			if cnt > 1 {
				return false
			}
			if i > 0 && b < nums[i-1] {
				nums[i+1] = a
			}
		}
	}
	return true
}

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
}
