package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	c, r := len(matrix), len(matrix[0]) // c为转置后的矩阵行数，r为转置后的矩阵列数
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
		for j := 0; j < c; j++ {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}
