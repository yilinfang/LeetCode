package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	for _, a := range A {
		l, r := 0, len(a)-1
		for l < r {
			if a[l] == a[r] {
				a[l] ^= 1
				a[r] ^= 1
			}
			l++
			r--
		}
		if l == r {
			a[l] ^= 1
		}
	}
	return A
}

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
}
