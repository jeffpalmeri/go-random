package largestrange

import (
	"fmt"
	"testing"
)

// go test -v -run TestLargestRange ./... -count=1
func TestLargestRange(t *testing.T) {
	// arr := []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}
	arr := []int{4, 2, 1, 3}
	ans := LargestRange(arr)
	fmt.Println("Answer: ", ans)
}
