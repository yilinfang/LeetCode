package main

import "fmt"

func subarraysWithKDistinct(A []int, K int) int {
	n := len(A)
	num1 := make([]int, n+1)
	num2 := make([]int, n+1)
	res := 0
	total1, total2, left1, left2 := 0, 0, 0, 0
	for _, v := range A {
		if num1[v] == 0 {
			total1++
		}
		num1[v]++
		if num2[v] == 0 {
			total2++
		}
		num2[v]++
		for total1 > K {
			num1[A[left1]]--
			if num1[A[left1]] == 0 {
				total1--
			}
			left1++
		}
		for total2 > K-1 {
			num2[A[left2]]--
			if num2[A[left2]] == 0 {
				total2--
			}
			left2++
		}
		res += left2 - left1
	}
	return res
}

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3))
}
