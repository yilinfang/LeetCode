package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if nums == nil || len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		//* 重复元素不再计算
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				//* 重复元素不再计算
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum(nil))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
