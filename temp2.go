package main

import "math"

func LevenshteinDistance2(a, b string) int {
	short, long := a, b
	if len(a) > len(b) {
		short, long = b, a
	}
	
	even := make([]int, len(short) + 1)
	odd := make([]int, len(short) + 1)
	var currentEdits, previousEdits []int

	for i := range even {
		even[i] = i
		odd[i] = math.MinInt32
	}
	for i := 1; i < len(long) + 1; i++ {
		if i % 2 == 1 {
			currentEdits = odd
			previousEdits = even
			} else {
				currentEdits = even
				previousEdits = odd
			}
		currentEdits[0] = i
		for j := 1; j < len(short) + 1; j++ {
			if short[j - 1] == long[i - 1] {
				currentEdits[j] = previousEdits[j - 1]
			} else {
				currentEdits[j] = min(currentEdits[j - 1], previousEdits[j - 1], previousEdits[j]) + 1
			}
		}
	}

	// if len(short) % 2 == 0 { // *************** wrong
	if len(long) % 2 == 0 { // *************** correct
		return even[len(even) - 1]
	} else {
		return odd[len(odd) - 1]
	}
}

func min(numbs ...int) int {
	current := numbs[0]
	for _, val := range numbs {
		if val < current {
			current = val
		}
	}
	return current
}
