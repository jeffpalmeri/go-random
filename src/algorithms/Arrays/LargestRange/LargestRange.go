package largestrange

import "fmt"

func LargestRange(arr []int) []int {
	// On first loop, put all values into a map
	vals := map[int]bool{}
	for _, val := range arr {
		if _, ok := vals[val]; !ok {
			vals[val] = true
		}
	}
	fmt.Println("Map: ", vals)

	tempRange := 1
	maxRange := 1
	rangeSlice := []int{1, 1}
	rangeAdder := 1

	for i := 0; i < len(arr); i++ {
		if arr[i] > rangeSlice[0] && arr[i] < rangeSlice[1] {
			continue
		}
		fmt.Println("*****", i)
		// if the next value needed in the range is in the map, then:
		//  - increment tempRange
		//  - check if tempRange is great than maxRange
		// for i+rangeAdder < len(arr) {
		for i < len(arr) {
			fmt.Println("i: ", i, tempRange, maxRange, arr[i]+rangeAdder)
			// fmt.Println(rangeAdder == tempRange, tempRange, rangeAdder)
			_, ok := vals[arr[i]+rangeAdder]
			if ok {
				tempRange += 1
				rangeAdder += 1
				if tempRange > maxRange {
					fmt.Println("new max: ", i, "Setting rangeSlice[0] to: ", i)
					maxRange = tempRange
					// rangeSlice[0], rangeSlice[1] = i, i+rangeAdder
					rangeSlice[0] = arr[i]
					rangeSlice[1] = arr[i] + rangeAdder - 1
				}
			} else {
				tempRange = 1
				rangeAdder = 1
				break
			}
		}

	}
	fmt.Println("Range: ", rangeSlice, maxRange)
	return rangeSlice
}
