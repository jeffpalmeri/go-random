package dynamicprogramming

import (
	"fmt"
	"math"
)

// func LvDistance(a, b string) int {
func LevenshteinDistance(a, b string) int {
	// Determine short and long string
	short, long := a, b
	if len(a) > len(b) {
		short, long = b, a
	}
	fmt.Println("Short: ", short)
	fmt.Println("Long: ", long)
	// Create even and odd slices, set to the length of the shortest string + 1
	even := make([]int, len(short) + 1)
	odd := make([]int, len(short) + 1)

	// Initialize previous and current slices
	var currentEdits, previousEdits []int
	// currentEdits := make([]int, len(short) + 1)
	// previousEdits := make([]int, len(short) + 1)

	// Set even slice to [0, 1, 2, ... n]
	for i := range even {
		even[i] = i
		odd[i] = math.MinInt32
	}
	fmt.Println("Even at this point: ", even)

	// Start looping
	for i := 1; i < len(long) + 1; i++ {
		// Determine whether we're at and even or odd index
		if i % 2 == 0 {
			currentEdits = even
			previousEdits = odd
			} else {
				currentEdits = odd
				previousEdits = even
			}
		currentEdits[0] = i
		for j := 1; j < len(short) + 1; j++ {
			// If the letters compared are the same
			if short[j - 1] == long[i - 1] {
				fmt.Println("Letters are even here: ", i, j)
				currentEdits[j] = previousEdits[j - 1]
			} else {
				fmt.Println("Letters are not event here")
				currentEdits[j] = min(currentEdits[j - 1], previousEdits[j - 1], previousEdits[j]) + 1
			}
		}
	}

	fmt.Println("Even: ", even)
	fmt.Println("Odd: ", odd)
	if len(short) % 2 == 0 {
		fmt.Println("Answer: ", even[len(even) - 1])
		return even[len(even) - 1]
	} else {
		fmt.Println("Answer; ", odd[len(odd) - 1])
		return odd[len(odd) - 1]
	}
}

func min(numbs ...int) int {
	// fmt.Println("Args: ", numbs)
	current := numbs[0]
	for _, val := range numbs {
		if val < current {
			current = val
		}
	}
	return current
}

// func LevenshteinDistance(a, b string) int {
// 	grid := make([][]string, len(a))
// 	for i := range grid {
// 		grid[i] = make([]string, len(b))
// 		for j := range grid[0] {
// 			grid[i][j] = "hi"
// 		}
// 	}


// 	m := make(map[string]int)

// 	m["one"] = 9;
// 	m["two"] = 10;

// 	fmt.Print(grid)
// 	fmt.Println("\ng---------------")
// 	fmt.Print(m)

// 	return 0
// }