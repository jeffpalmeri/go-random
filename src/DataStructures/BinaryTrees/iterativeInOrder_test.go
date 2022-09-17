package binarytrees

import (
	"fmt"
	"io"
	"testing"
	// binarytreeutils "github.com/jeffpalmeri/go-files/src/DataStructures/BinaryTrees/utils"
)

func PrintBinaryTreeC(w io.Writer, node *BinaryTreeC, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.value)
	PrintBinaryTreeC(w, node.left, ns+2, 'L')
	PrintBinaryTreeC(w, node.right, ns+2, 'R')
}

type BinaryTreeC struct {
	value int

	left   *BinaryTreeC
	right  *BinaryTreeC
	parent *BinaryTreeC
}

func (tree *BinaryTreeC) IterativeInOrderTraversal(callback func(int)) {
	HRS := map[int]bool{}
	visited := map[int]bool{}
	arr := []int{}

	startingNode := getLeftMostNode(tree, &HRS)
	fmt.Println(startingNode, HRS)
	/*
		fmt.Println("HRS: ", HRS)
		arr = append(arr, startingNode.value)
		visited[startingNode.value] = true
	*/
	curr := startingNode
	/*
		for curr != nil {
			// if the parent has a right subtree, add parent to visited, and then get left most node of that right subtree to continue the traversal on
			// If the parent doesn't have a right subtree, then add it to visited and continue
			if curr.parent != nil && !HRS[curr.parent.value] {
				arr = append(arr, curr.parent.value)
				visited[curr.parent.value] = true
				curr = curr.parent
			} else if curr.parent != nil && HRS[curr.parent.value] {
				curr = getLeftMostNode(curr.right, &HRS)
				continue
			}
		}
	*/

	for curr != nil {
		if !visited[curr.value] {
			arr = append(arr, curr.value)
			visited[curr.value] = true
		}

		if HRS[curr.value] && !visited[curr.value] {
			// curr has a right subtree, so get it's left most value and continue
			curr = getLeftMostNode(curr.right, &HRS)
			continue
		} else {
			// curr does not have a right subtree, so need to continue moving back up the tree
			curr = curr.parent
		}
	}

}

func getLeftMostNode(node *BinaryTreeC, HRS *map[int]bool) *BinaryTreeC {
	curr := node
	for node != nil {
		if node.right != nil {
			(*HRS)[node.value] = true
		}
		curr = node.left
	}
	return curr
}

func MakeTreeC() *BinaryTreeC {
	one := BinaryTreeC{1, nil, nil, nil}
	two := BinaryTreeC{2, nil, nil, nil}
	three := BinaryTreeC{3, nil, nil, nil}
	four := BinaryTreeC{4, nil, nil, nil}
	six := BinaryTreeC{6, nil, nil, nil}
	seven := BinaryTreeC{7, nil, nil, nil}
	nine := BinaryTreeC{9, nil, nil, nil}

	one.left = &two
	one.right = &three
	two.left = &four
	four.right = &nine
	three.left = &six
	three.right = &seven

	nine.parent = &four
	four.parent = &two
	two.parent = &one
	three.parent = &one
	six.parent = &three
	seven.parent = &three

	return &one
}

// go test -v -run TestStuff ./... -count=1
func TestStuff(t *testing.T) {
	tree := MakeTreeC()
	// PrintBinaryTreeC(os.Stdout, tree, 0, 'R')
	c := func(int) {}
	tree.IterativeInOrderTraversal(c)
}
