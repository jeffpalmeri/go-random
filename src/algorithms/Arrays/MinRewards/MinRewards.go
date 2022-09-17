package minrewards

import (
	"fmt"
	"math"
)

// Get a list of the local mins
func MinRewards(arr []int) int {
	ans := make([]int, len(arr))
	smallest := math.MaxInt
	smallestIdx := -1
	for i, val := range arr {
		if val < smallest {
			smallest = val
			smallestIdx = i
		}
	}
	fmt.Println(smallest, smallestIdx)
	ans[smallestIdx] = 1
	left, right := smallestIdx-1, smallestIdx+1
	for left >= 0 {
		fmt.Println("Left loop: ", left)
		if arr[left] < arr[left+1] {
			ans[left] = ans[left+1] - 1
		} else {
			ans[left] = ans[left+1] + 1
		}
		left--
	}
	for right < len(arr) {
		if arr[right] > arr[right-1] {
			ans[right] = ans[right-1] + 1
		} else {
			ans[right] = ans[right-1] - 1
		}
		right++
	}

	fmt.Println(ans)
	return sum(ans)
}

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
