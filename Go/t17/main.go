package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := []string{}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	for _, s := range m[digits[0]] {
		res = append(res, s)
	}
	for i := 1; i < len(digits); i++ {
		l := len(res)
		for j := 0; j < l; j++ {
			str := res[0]
			res = res[1:]
			for _, s := range m[digits[i]] {
				res = append(res, str+s)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("79"))
	fmt.Println(letterCombinations("234"))
}
