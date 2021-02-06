package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	if people == nil {
		return nil
	}
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	})
	return people
}
