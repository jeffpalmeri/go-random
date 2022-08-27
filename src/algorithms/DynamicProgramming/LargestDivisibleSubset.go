package dynamicprogramming

import (
	"fmt"
	"sort"
)

// [1,2,4,8]

// go test -v -run .TestLongestIncreasingSubsequence /... -count=1
func LargestDivisibleSubset(input []int) []int {
	sort.Ints(input)
	lengths := make([]int, len(input))
	indicies := make([]int, len(input))
	lengths[0] = 1
	indicies[0] = -1
	for i := 0; i < len(lengths); i++ {
		lengths[i] = 1
		indicies[i] = -1
	}
	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			if input[j] % input[i] == 0 ||  input[i] % input[j] == 0 {
				if lengths[j] + 1 > lengths[i] {
					lengths[i] = lengths[j] + 1
					indicies[i] = j // ??
				}
			}
		}
	}
	fmt.Println("Lengths array: ", lengths)
	fmt.Println("Indicies array: ", indicies)
	a := LDSbuildArr(lengths, indicies, input)
	fmt.Println("Answer: ", a)
	// return []int{1,2,3}
	return a
}

func LDSbuildArr(lengths []int, indicies []int, originalArr []int) []int {
	arr := []int{}
	maxIndex := 0
	maxValue := 0
	// maxIndex = 10 for example
	for i, val := range lengths {
		if val > maxValue {
			maxValue = val
			maxIndex = i
		}
	}

	for maxIndex != -1 {
		arr = append(arr, originalArr[maxIndex])
		maxIndex = indicies[maxIndex] 
	}

	return LDSreverse(arr)
}

func LDSreverse(a []int) []int {
	last := len(a) - 1
	first := 0
	for last > first {
		a[first], a[last] = a[last], a[first]
		last--
		first++
	}

	return a
}


// https://cs.opensource.google/go/x/tools/+/refs/tags/gopls/v0.9.1:gopls/doc/vim.md#neovim-install