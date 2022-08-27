package leetcodeweeklycontests

import (
    "fmt"
)

func largestLocal(grid [][]int) [][]int {
    res := make([][]int, len(grid) - 2)
    for i := range res {
        res[i] = make([]int, len(grid[0]) - 2)
    }
    fmt.Println("Res: ", res)
    
    for i := 0; i < len(grid[0]) - 2; i++ {
        for j := 0; j < len(grid[0]) - 2; j++ {
            for m := i; m < i + 3; m++ {
                for n := j; n < j + 3; n++ {
                    res[i][j] = max(res[i][j], grid[m][n])
                }
            }
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
