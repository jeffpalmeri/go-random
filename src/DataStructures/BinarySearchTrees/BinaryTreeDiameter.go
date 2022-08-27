package binarysearchtrees

import (
	"fmt"
	"math"
)

/*
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}
*/
type TreeInfo struct {
	height int
	width  int
}

func BinaryTreeDiameter(tree *BinaryTree) []int {
	// Write your code here.
	fmt.Println("Hello I'm writing this in nvim!!!!")
	node := tree
	values := []int{}
	for node != nil {

	}
	return s(node, values)
}

func s(node *BinaryTree, values []int) []int {
	if node.Left != nil {
		values = s(node.Left, values)
	}
	values = append(values, node.Value)
	if node.Right != nil {
		values = s(node.Right, values)
	}
	return values
}

func doWork(node *BinaryTree) (int, int) {
	if node == nil {
		return 0, math.MinInt
	}
	leftMaxBranchSum, leftMaxSoFar := doWork(node.Left)
	rightMaxBranchSum, rightMaxSoFar := doWork(node.Right)
	maxBranchSumNoCurrent := max(leftMaxBranchSum, rightMaxBranchSum)

	// value := node.Value
	maxBranchSum := max(maxBranchSumNoCurrent+node.Value, node.Value)
	maxSumWithTriangle := max(maxBranchSum, leftMaxBranchSum+node.Value, +rightMaxBranchSum)
	newMaxValue := max(maxSumWithTriangle, leftMaxSoFar, rightMaxSoFar)

	return maxBranchSum, newMaxValue

}

func max(first int, numbers ...int) int {
	current := first
	for _, val := range numbers {
		if val > current {
			current = val
		}
	}
	return current
}
