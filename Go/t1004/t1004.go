package main

import (
	"fmt"
)

func longestOnes(A []int, K int) (ans int) {
	var left, lsum, rsum int
	for right, v := range A {
		rsum += 1 - v
		for lsum < rsum-K {
			lsum += 1 - A[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
