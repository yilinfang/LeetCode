package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	idx := 0
	for ; idx < len(s); idx++ {
		if s[idx] != ' ' {
			break
		}
	}
	if idx == len(s) || s[idx] != '-' && s[idx] != '+' && (s[idx] < '0' || s[idx] > '9') {
		return 0
	}
	sig := true
	if s[idx] == '-' || s[idx] == '+' {
		sig = s[idx] != '-'
		idx++
	}
	res := 0
	for ; idx < len(s) && '0' <= s[idx] && s[idx] <= '9'; idx++ {
		if sig {
			res = 10*res + int(s[idx]-'0')
			if res > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			res = 10*res - int(s[idx]-'0')
			if res < math.MinInt32 {
				return math.MinInt32
			}
		}
	}
	return res
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("+"))
	fmt.Println(myAtoi("  0000000000012345678"))
	fmt.Println(myAtoi("9223372036854775808"))
}
