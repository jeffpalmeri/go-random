package minrewards

import (
	"fmt"
	"testing"
)

// go test -v -run TestMinRewards ./... -count=1
func TestMinRewards(t *testing.T) {
	arr := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}
	ans := MinRewards(arr)
	fmt.Println("Ans: ", ans)
}
