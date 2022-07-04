package dynamicprogramming

import (
	"testing"
)

// go test -v ./src/algorithms/DynamicProgramming
func TestFunc(t *testing.T) {
	LongestCommonSubsequence("hi", "h")
	LongestCommonSubsequence("ZXVVYZW", "XKYKZPW")
}

// go test -v -run TestLevenshteinDistance ./...
func TestLevenshteinDistance(t *testing.T) {
	LevenshteinDistance("word", "one")
}

func TestKnapsackProblem(t *testing.T) {
	// var items = [][]int{
	// 	{1, 2},
	// 	{4, 3},
	// 	{5, 6},
	// 	{6, 7},
	// }
	// var items2 = [][]int{
	// 	{1, 2},
	// 	{4, 3},
	// 	{5, 6},
	// 	{6, 9},
	// }
	var items3 = [][]int{
		{400, 85},
    {465, 100},
    {255, 55},
    {350, 45},
    {650, 130},
    {1000, 190},
    {455, 100},
    {100, 25},
    {1200, 190},
    {320, 65},
    {750, 100},
    {50, 45},
    {550, 65},
    {100, 50},
    {600, 70},
    {240, 40},
	}

	// fmt.Println(KnapsackProblem(items, 10))
	// KnapsackProblem(items, 10)
	// KnapsackProblem(items2, 11)
	KnapsackProblem(items3, 200)
}