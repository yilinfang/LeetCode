package main

import "fmt"

// NumArray is a exported type
type NumArray struct {
	sums []int
}

// Constructor is a exported function
func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums: sums}
}

// SumRange is a exported funciton
func (na *NumArray) SumRange(i int, j int) int {
	return na.sums[j+1] - na.sums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	na := Constructor(nums)
	fmt.Println(na.SumRange(0, 2))
	fmt.Println(na.SumRange(2, 5))
	fmt.Println(na.SumRange(0, 5))
}
