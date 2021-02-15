package main

import "fmt"

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	v1 := m[s[0]]
	res := v1
	for i := 1; i < len(s); i++ {
		v2 := m[s[i]]
		if v1 < v2 {
			res += v2 - 2*v1
		} else {
			res += v2
		}
		v1 = v2
	}
	return res
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
