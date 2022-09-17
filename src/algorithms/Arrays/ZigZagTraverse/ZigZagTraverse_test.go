package zigzagtraverse

import (
	"fmt"
	"testing"
)

// go test -v -run TestZigZagTraverse ./... -count=1
func TestZigZagTraverse(t *testing.T) {
	grid := [][]int{
		{1, 3, 4, 10},
		{2, 5, 9, 11},
		{6, 8, 12, 15},
		{7, 13, 14, 16},
	}
	fmt.Println("Grid: ", grid)
	ans := ZigZagTraverse(grid)
	fmt.Println(ans)
}
