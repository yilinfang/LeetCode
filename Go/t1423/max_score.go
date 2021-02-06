package main

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	w := n - k
	s := 0
	for _, p := range cardPoints[:w] {
		s += p
	}
	min := s
	for i := w; i < n; i++ {
		s += cardPoints[i] - cardPoints[i-w]
		if s < min {
			min = s
		}
	}
	total := 0
	for _, p := range cardPoints {
		total += p
	}
	return total - min
}
