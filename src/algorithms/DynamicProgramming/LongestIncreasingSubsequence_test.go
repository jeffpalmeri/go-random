package dynamicprogramming

import (
	"fmt"
	"testing"
)

// go test -v -run TestLongestIncreasingSubsequence ./... -count=1
func TestLongestIncreasingSubsequence(t *testing.T) {
	var sample = []int{5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35}
	fmt.Println(sample)
	fmt.Println("Another line")

	LongestIncreasingSubsequence(sample)
}

// go test -v -run TestLargestDivisibleSubset ./... -count=1
func TestLargestDivisibleSubset(t *testing.T) {
	// var sample = []int{1,2,4,8,9,2,7,9,9,9}
	var sample = []int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}
	fmt.Println(sample)
	fmt.Println("Largest Divisible Subset")

	LargestDivisibleSubset(sample)
}
