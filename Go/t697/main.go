package main

import "fmt"

func findShortestSubArray(nums []int) int {
	d := make(map[int]int)
	fp := make(map[int]int)
	lp := make(map[int]int)
	for i, v := range nums {
		value, ok := d[v]
		if ok {
			d[v] = value + 1
			lp[v] = i
		} else {
			d[v] = 1
			fp[v] = i
			lp[v] = i
		}

	}
	check := make([]int, 0)
	max := 0
	for k, v := range d {
		if v > max {
			max = v
			check = check[0:0]
			check = append(check, k)
		} else if v == max {
			check = append(check, k)
		}
	}
	res := len(nums)
	for _, v := range check {
		r := lp[v] - fp[v] + 1
		if r < res {
			res = r
		}
	}
	return res
}

func main() {
	fmt.Println(findShortestSubArray([]int{1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}
