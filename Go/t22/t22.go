package main

import "fmt"

func generateParenthesis(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	dp[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, v1 := range dp[j] {
				for _, v2 := range dp[i-1-j] {
					dp[i] = append(dp[i], "("+v1+")"+v2)
				}
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
	fmt.Println(generateParenthesis(5))
	fmt.Println(generateParenthesis(6))
	fmt.Println(generateParenthesis(7))
	fmt.Println(generateParenthesis(8))
}
