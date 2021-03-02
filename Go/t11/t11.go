package main

import "fmt"

func maxArea(height []int) int {
	res := 0
	i, j := 0, len(height)-1
	for i < j {
		cap := 0
		if height[i] <= height[j] {
			cap = (j - i) * height[i]
			i++
		} else {
			cap = (j - i) * height[j]
			j--
		}
		if cap > res {
			res = cap
		}
	}
	return res
}

func main() {
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
	fmt.Println(maxArea([]int{1, 2, 1}))
}
