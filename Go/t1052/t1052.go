package main

import "fmt"

func maxSatisfied(customers []int, grumpy []int, X int) int {
	res := 0
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			res += customers[i]
			customers[i] = 0
		}
	}
	cur, max := 0, 0
	for i, j := 0, 0; i < n; i++ {
		cur += customers[i]
		if i == j+X {
			cur -= customers[j]
			j++
		}
		if cur > max {
			max = cur
		}
	}
	return res + max
}

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}
