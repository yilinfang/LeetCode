package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x, y := (i*n+j)/c, (i*n+j)%c
			res[x][y] = nums[i][j]
		}
	}
	return res
}

func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
}
