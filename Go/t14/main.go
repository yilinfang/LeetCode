package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	length := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < length {
			length = len(strs[i])
		}
	}
	for i := 0; i < length; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if c != strs[j][i] {
				return res
			}
		}
		res += string(c)
	}
	return res
}

func main() {
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
