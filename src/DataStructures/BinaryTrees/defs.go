package binarytrees

import (
	"fmt"
	"io"
)

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func MakeTree() *BinaryTree {
	tree := BinaryTree{100, nil, nil}
	tree.InsertRecursive(&BinaryTree{-20, nil, nil})
	tree.InsertRecursive(&BinaryTree{-50, nil, nil})
	tree.InsertRecursive(&BinaryTree{-15, nil, nil})
	tree.InsertRecursive(&BinaryTree{-60, nil, nil})
	tree.InsertRecursive(&BinaryTree{50, nil, nil})
	tree.InsertRecursive(&BinaryTree{60, nil, nil})
	tree.InsertRecursive(&BinaryTree{55, nil, nil})
	tree.InsertRecursive(&BinaryTree{85, nil, nil})
	tree.InsertRecursive(&BinaryTree{15, nil, nil})
	tree.InsertRecursive(&BinaryTree{5, nil, nil})
	tree.InsertRecursive(&BinaryTree{-10, nil, nil})
	return &tree
}

func (root *BinaryTree) InsertBF(node *BinaryTree) {
	if root.left == nil && root.right == nil {
		root.left = node
	}
	if root.left == nil {
		root.left = node
	}
	if root.right == nil {
		root.right = node
	}
}

func (root *BinaryTree) InsertRecursive(node *BinaryTree) {
	if node.value < root.value {
		if root.left == nil {
			root.left = node
		} else {
			root.left.InsertRecursive(node)
		}
	}
	if node.value >= root.value {
		if root.right == nil {
			root.right = node
		} else {
			root.right.InsertRecursive((node))
		}
	}
}

func PrintBinaryTree(w io.Writer, node *BinaryTree, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.value)
	PrintBinaryTree(w, node.left, ns+2, 'L')
	PrintBinaryTree(w, node.right, ns+2, 'R')
}

func AnotherTrav(node *BinaryTree, values []int) []int {
	if node.left != nil {
		values = AnotherTrav(node.left, values)
	}
	values = append(values, node.value)
	if node.right != nil {
		values = AnotherTrav(node.right, values)
	}
	return values
}

func BreadthFirstSearch(node *BinaryTree, values []int) []int {
	fmt.Println("This should happen")
	queue := make([]BinaryTree, 0)
	current := *node

	queue = append(queue, *node)

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		values = append(values, current.value)
		if current.left != nil {
			queue = append(queue, *current.left)
		}
		if current.right != nil {
			queue = append(queue, *current.right)
		}

	}
	return values
}

func GetParents(node *BinaryTree, m map[int]*BinaryTree, parent *BinaryTree) map[int]*BinaryTree {
	if node != nil {
		m[node.value] = parent
		GetParents(node.left, m, node)
		GetParents(node.right, m, node)
	}
	return m
}

func BreadthFirstInsertion(node *BinaryTree, newNode *BinaryTree) {
	queue := []*BinaryTree{node}
	var curr *BinaryTree
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]

		if curr.left == nil {
			curr.left = newNode
			break
		}
		if curr.right == nil {
			curr.right = newNode
			break
		}
		queue = append(queue, curr.left)
		queue = append(queue, curr.right)
	}
}
