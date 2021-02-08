package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	// return wiggleMaxLengthDP(nums)
	return wiggleMaxLengthGreedy(nums)
}

func wiggleMaxLengthDP(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = 1
		dp[i][1] = 0
	}
	for i := 1; i < len(nums); i++ {
		l, d := 1, 0
		for j := 0; j < i; j++ {
			if dp[j][1] == 0 {
				if nums[i] > nums[j] {
					if l == 1 {
						l = 2
						d = -1
					}
				} else if nums[i] < nums[j] {
					if l == 1 {
						l = 2
						d = 1
					}
				}
			} else if dp[j][1] == 1 {
				if nums[i] > nums[j] {
					if dp[j][0]+1 > l {
						l = dp[j][0] + 1
						d = -1
					}
				}
			} else {
				if nums[i] < nums[j] {
					if dp[j][0]+1 > l {
						l = dp[j][0] + 1
						d = 1
					}
				}
			}
		}
		dp[i][0] = l
		dp[i][1] = d
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		if dp[i][0] > res {
			res = dp[i][0]
		}
	}
	return res
}

func wiggleMaxLengthGreedy(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	len1, len2 := 1, 1
	sig1, sig2 := 1, -1
	idx1, idx2 := 0, 0
	for i := 1; i < len(nums); i++ {
		if sig1 == 1 {
			if nums[i] > nums[idx1] {
				sig1 = -1
				idx1 = i
				len1++
			}
		} else {
			if nums[i] < nums[idx1] {
				sig1 = 1
				idx1 = i
				len1++
			}
		}
		if sig2 == 1 {
			if nums[i] > nums[idx2] {
				sig2 = -1
				idx2 = i
				len2++
			}
		} else {
			if nums[i] < nums[idx2] {
				sig2 = 1
				idx2 = i
				len2++
			}
		}
	}
	if len1 > len2 {
		return len1
	}
	return len2
}

func main() {
	fmt.Println(wiggleMaxLength(nil))
	fmt.Println(wiggleMaxLength([]int{}))
	fmt.Println(wiggleMaxLength([]int{1}))
	fmt.Println(wiggleMaxLength([]int{1, 1}))
	fmt.Println(wiggleMaxLength([]int{1, 2}))
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(wiggleMaxLength([]int{33, 53, 12, 64, 50, 41, 45, 21, 97, 35, 47, 92, 39, 0, 93, 55, 40, 46, 69, 42, 6, 95, 51, 68, 72, 9, 32, 84, 34, 64, 6, 2, 26, 98, 3, 43, 30, 60, 3, 68, 82, 9, 97, 19, 27, 98, 99, 4, 30, 96, 37, 9, 78, 43, 64, 4, 65, 30, 84, 90, 87, 64, 18, 50, 60, 1, 40, 32, 48, 50, 76, 100, 57, 29, 63, 53, 46, 57, 93, 98, 42, 80, 82, 9, 41, 55, 69, 84, 82, 79, 30, 79, 18, 97, 67, 23, 52, 38, 74, 15}))
}
