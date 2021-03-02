package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	for _, v := range arr {
		if k >= v {
			k++
		}
	}
	return k
}

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
}
