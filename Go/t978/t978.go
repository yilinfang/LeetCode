package main

import "fmt"

func maxTurbulenceSize(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return len(arr)
	}
	length, res := 1, 0
	sig := 0
	for i := 1; i < len(arr); i++ {
		if sig == 0 {
			if arr[i] > arr[i-1] {
				length = 2
				sig = -1
			} else if arr[i] < arr[i-1] {
				length = 2
				sig = 1
			} else {
				length = 1
				sig = 0
			}
		} else if sig == -1 {
			if arr[i] < arr[i-1] {
				length++
				sig = 1
			} else if arr[i] > arr[i-1] {
				length = 2
				sig = -1
			} else {
				length = 1
				sig = 0
			}
		} else {
			if arr[i] > arr[i-1] {
				sig = -1
				length++
			} else if arr[i] < arr[i-1] {
				length = 2
				sig = 1
			} else {
				length = 1
				sig = 0
			}
		}
		if length > res {
			res = length
		}
	}
	return res
}

func main() {
	fmt.Println(maxTurbulenceSize(nil))
	fmt.Println(maxTurbulenceSize([]int{}))
	fmt.Println(maxTurbulenceSize([]int{1}))
	fmt.Println(maxTurbulenceSize([]int{1, 0}))
	fmt.Println(maxTurbulenceSize([]int{1, 1}))
	fmt.Println(maxTurbulenceSize([]int{4, 8, 12, 16}))
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
}
