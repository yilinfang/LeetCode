package main

func sortArrayByParityII(A []int) []int {
	r := make([]int, len(A))
	a := 0
	b := 1
	for _, i := range A {
		if i%2 == 0 {
			r[a] = i
			a += 2
		} else {
			r[b] = i
			b += 2
		}
	}
	return r
}
