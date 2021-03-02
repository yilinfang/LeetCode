package main

import "math"

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	s := 0
	minS := math.MaxInt64
	minI := 0
	for i := 0; i < l; i++ {
		s += gas[i] - cost[i]
		if s < minS {
			minS = s
			minI = i
		}
	}
	if s < 0 {
		return -1
	}
	return (minI + 1) % l
}
