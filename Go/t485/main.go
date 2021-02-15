package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	n1, n2, res := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if n1 == 1 {
				n1 = 0
				if n2 > res {
					res = n2
				}
				n2 = 0
			}
		} else {
			if n1 == 0 {
				n1 = 1
				n2 = 1
			} else {
				n2++
			}
		}
	}
	if n2 > res {
		res = n2
	}
	return res
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
