package dynamicprogramming

import (
	"fmt"
)

// type answer interface {
// 	value() int
// 	indicies() []int
// }

type answer struct {
	value int
	indicies []int
}

func print2dGrid(g [][]int) {
	fmt.Println("[")
	for i := 0; i < len(g); i++ {
		fmt.Println(" ", g[i])
	}
	fmt.Println("]")
}

func KnapsackProblem(items [][]int, limit int) answer {

	grid := make([][]int, len(items))
	for i := 0; i < len(items); i++ {
		grid[i] = make([]int, limit + 1)
		for j := 0; j <= limit; j++ {
			grid[i][j] = 0
		}
	}
	
	// Fill first row of grid
	for j := 1; j <= limit; j++ {
		if items[0][1] <= j {
			grid[0][j] = items[0][0]
		} else {
			grid[0][j] = 0
		}
	}

	// Loop through rest of grid
	for i := 1; i < len(items); i++ {
		item := items[i]
		value := item[0]
		weight := item[1]

		grid[i][0] = 0
		for j := 1; j <= limit; j++ {
			if weight <= j {
				grid[i][j] = max(value + grid[i - 1][j - weight], grid[i - 1][j])
			} else {
				grid[i][j] = grid[i - 1][j]
			}
		}
	}

	print2dGrid(grid)
	fmt.Println(buildIndicies(grid, items, limit))

	trueAnswer := answer{value: grid[len(grid) - 1][len(grid[0]) - 1], indicies: buildIndicies(grid, items, limit)}
	fmt.Println("True Answer: ", trueAnswer)

	return trueAnswer
}

func buildIndicies(grid [][]int, items [][]int, limit int) []int {
	i := len(grid) - 1
	j := limit
	idx := []int{}

	for i > 0 {
		if i == -1 {
			fmt.Println("i is negative")
		}

		if grid[i][j] != grid[i - 1][j] {

			// t := items[i]
			// fmt.Println("t: ", t)

			// a = append(a, []int{t})
			// a = append(a, t)
			idx = append(idx, i)
			j = j - items[i][1]
		}
		i -= 1
	}

	if i == 0 && grid[i][j] != 0 {
		idx= append(idx, 0)
	}

	return idx
}

func max(numbs ...int) int {
	current := numbs[0]
	for _, val := range numbs {
		if val > current {
			current = val
		}
	}
	return current
}