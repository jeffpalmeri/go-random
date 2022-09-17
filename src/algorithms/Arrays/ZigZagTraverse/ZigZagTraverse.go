package zigzagtraverse

import "fmt"

func ZigZagTraverse(grid [][]int) []int {
	fmt.Println("*********************************", grid)
	ans := []int{}
	i, j := 0, 0

	currState := down
	prevState := none
	for i != len(grid)-1 && j != len(grid[0])-1 {
		curr := grid[i][j]
		ans = append(ans, curr)

		i, j, currState, prevState = handleState(i, j, currState, prevState, grid)

		i = len(grid) - 1
		j = len(grid[0]) - 1
	}

	return ans
}

type state int

// down, left-down-diagonal, right-up-diagonal, right
const (
	down state = iota
	leftDownDiagonal
	rightUpDiagonal
	right
	none
)

func handleState(i, j int, state state, prevState state, grid [][]int) (int, int, state, state) {
	prevState = state
	nextState := state
	// iGood = iInbounds(i+1, grid)
	// jGood = jInbound(j+1, grid)
	maxI, maxJ := len(grid)-1, len(grid[0])-1
	switch state {
	case down:
		fmt.Println("In down state")
		if prevState == none || prevState == leftDownDiagonal {
			state = rightUpDiagonal
			i -= i
		} else if prevState == rightUpDiagonal {
			state = leftDownDiagonal
			i -= 1
		} else {
			fmt.Println("In an unaccounted for case")
		}
	case leftDownDiagonal:
		fmt.Println("In leftDownDiagaonState")
		if i != 0 && j != maxJ {
			i++
			j--
		} else if i == maxI && j != maxJ {

		}
	case rightUpDiagonal:
		fmt.Println("In rightUpDiagonal state")
		if i > 0 && j < maxJ {
			i++
			j++
		} else if i == 0 && j == maxJ {
			nextState = down
			i--
		} else if i == 0 && j < maxJ {
			nextState = right
			j++
		} else if i != 0 && j < maxJ {
			i--
			j++
		} else if i != 0 && j == maxJ {
			i++
			nextState = down
		}

	case right:
		fmt.Println("In right State")
	}

	return i, j, nextState, prevState
}

func jInbounds(j int, grid [][]int) bool {
	return j <= len(grid[0])-1
}

func iInbounds(i int, grid [][]int) bool {
	return i <= len(grid)-1
}
