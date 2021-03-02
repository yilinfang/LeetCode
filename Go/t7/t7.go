package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	sig := 1
	if x < 0 {
		x = -x
		sig = -1
	}
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}
	res *= sig
	if math.MinInt32 <= res && res <= math.MaxInt32 {
		return res
	}
	return 0
}

func main() {
	fmt.Println(reverse(120))
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(0))
	fmt.Println(reverse(1534236469))
}
