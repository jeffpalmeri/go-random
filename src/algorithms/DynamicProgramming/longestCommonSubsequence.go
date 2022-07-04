package dynamicprogramming

import "fmt"

type Element struct {
	letter byte
	length int
	prevI int
	prevJ int
}

func LongestCommonSubsequence(str1 string, str2 string) []byte {
	fmt.Println("Hello from dynamic programming!!!")

	grid := make([][]Element, len(str1) + 1)

	for i := range grid {
		grid[i] = make([]Element, len(str2) + 1)
		for j := range grid[0] {
			grid[i][j] = Element{
				letter: 0,
				length: 0,
				prevI: -1,
				prevJ: -1,
			}
		}
	}

	fmt.Println("test")
	for i := 1; i < len(str1) + 1; i++ {
		for j := 1; j < len(str2) + 1; j++ {
			if(str1[i - 1] == str2[j - 1]) {
				grid[i][j].letter = str1[i - 1]
				grid[i][j].length = grid[i - 1][j - 1].length + 1
				grid[i][j].prevI = i - 1
				grid[i][j].prevJ = j - 1
			} else {
				if grid[i - 1][j].length > grid[i][j - 1].length {
					grid[i][j] = Element{
						letter: 0,
						length: grid[i - 1][j].length,
						prevI: i - 1,
						prevJ: j,
					}
				} else {
					grid[i][j] = Element{
						letter: 0,
						length: grid[i][j - 1].length,
						prevI: i,
						prevJ: j - 1,
					}
				}
			}
		}
	}

	// fmt.Print(grid)
	fmt.Println("---------")
	fmt.Println(buildSequence((grid)))
	
	return buildSequence(grid)
}

func buildSequence(grid [][]Element) []byte {
	sequence := make([]byte, 0)
	i := len(grid) - 1
	j := len(grid[0]) - 1

	for i != 0 && j != 0 {
		current := grid[i][j]
		if(current.letter != 0) {
			sequence = append(sequence, grid[i][j].letter)
		}
		i = current.prevI
		j = current.prevJ
	}

	return reverse(sequence)
}

func reverse(arr []byte) []byte {
	i := 0
	j := len(arr) - 1

	for i <= j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}
	return arr
}