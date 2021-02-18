package main

import (
	"fmt"
)

func minKBitFlips(A []int, K int) int {
	n := len(A)
	res := 0
	dp := make([]int, n+1)
	cnt := 0
	for i := 0; i < n; i++ {
		cnt += dp[i]
		if (A[i]+cnt)%2 == 0 {
			if i+K > n {
				return -1
			}
			dp[i+1]++
			dp[i+K]--
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(minKBitFlips([]int{0, 1, 0}, 1))
	fmt.Println(minKBitFlips([]int{1, 1, 0}, 2))
	fmt.Println(minKBitFlips([]int{0, 0, 0, 1, 0, 1, 1, 0}, 3))
}
