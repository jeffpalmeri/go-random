package apartmenthunting

import (
	"fmt"
	"math"
)

type Block map[string]bool

func ApartmentHunting(blocks []Block, reqs []string) int {
	reqStatus := map[string]int{}
	for _, val := range reqs {
		reqStatus[val] = 0
	}

	left, right := 0, 0
	window := []int{0, math.MaxInt}

	updateReqStatus(blocks[0], &reqStatus, true, reqs)

	for left < len(blocks)-1 && right <= len(blocks)-1 {
		fmt.Println("l r", left, right)
		fmt.Println("Contains reqs: ", containsAllReqs(&reqStatus, reqs))
		fmt.Println(reqStatus)
		if containsAllReqs(&reqStatus, reqs) || right == len(blocks)-1 {
			// Do longest calculation first, then adjust things for the next interation
			if containsAllReqs(&reqStatus, reqs) && right-left < window[1]-window[0] {
				window[0], window[1] = left, right
			}
			fmt.Println("here? ", left)
			left += 1
			updateReqStatus(blocks[left-1], &reqStatus, false, reqs)
		} else {
			right += 1
			updateReqStatus(blocks[right], &reqStatus, true, reqs)
			fmt.Println("Updating reqs now")
		}
	}
	fmt.Println("window: ", window)
	return (window[0] + window[1]) / 2
}

func updateReqStatus(newBlock Block, reqStatus *map[string]int, add bool, reqs []string) *map[string]int {
	for _, val := range reqs {
		if add {
			if newBlock[val] {
				(*reqStatus)[val] = (*reqStatus)[val] + 1
			}
		} else {
			if newBlock[val] && (*reqStatus)[val] != 0 {
				(*reqStatus)[val] = (*reqStatus)[val] - 1
			}
		}
	}
	return reqStatus
}

func containsAllReqs(reqStatus *map[string]int, reqs []string) bool {
	for _, val := range reqs {
		if (*reqStatus)[val] > 0 {
			continue
		} else {
			return false
		}
	}
	return true
}
