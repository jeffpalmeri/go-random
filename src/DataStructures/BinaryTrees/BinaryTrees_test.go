package binarytrees

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run TestBasicStuff ./... -count=1
func TestBasicStuff(t *testing.T) {
	tree := MakeTree()
	PrintBinaryTree(os.Stdout, tree, 0, 'H')

	m := make([]int, 0)
	// m = AnotherTrav(tree, m)
	// fmt.Println("Values: ", m)
	m = BreadthFirstSearch(tree, m)
	fmt.Println("Values: ", m)

	fmt.Println("*****************************")

	ma := make(map[int]*BinaryTree)
	ma = GetParents(tree, ma, nil)
	fmt.Println("Parents Map: ", ma)

	root := BinaryTree{5, nil, nil}
	BreadthFirstInsertion(&root, &BinaryTree{4, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{5, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{6, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{7, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{8, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{9, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{10, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{11, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{12, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{13, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{14, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{15, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{16, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{17, nil, nil})
	fmt.Println("Root now: ", root)
	PrintBinaryTree(os.Stdout, &root, 0, 'H')

	testSlice := []int{1, 2, 3}
	fmt.Println("testSlice before: ", testSlice, len(testSlice))
	slicePointerTest(&testSlice)
	fmt.Println("testSlice after: ", testSlice, len(testSlice))

}

func slicePointerTest(sl *[]int) {
	s := *sl
	(*sl) = append((*sl), 8)
	(*sl) = append((*sl), 9)
	(*sl) = append((*sl), 10)
	(*sl) = append((*sl), 11)
	(*sl) = append((*sl), 12)
	(*sl) = append((*sl), 13)
	s = append(s, 5)
	s = append(s, 5)
	s = append(s, 5)
	s = append(s, 5)
	s = append(s, 5)
}

func makeBinaryTree() *BinaryTree {
	root := BinaryTree{1, nil, nil}
	BreadthFirstInsertion(&root, &BinaryTree{4, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{5, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{6, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{7, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{8, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{9, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{10, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{11, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{12, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{13, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{14, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{15, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{16, nil, nil})
	BreadthFirstInsertion(&root, &BinaryTree{17, nil, nil})
	return &root
}

// go test -v -run TestKDistance ./... -count=1
func TestKDistance(t *testing.T) {
	ans := []int{}
	root := makeBinaryTree()
	PrintBinaryTree(os.Stdout, root, 0, 'R')
	parents := map[int]*BinaryTree{}
	parents = GetParents(root, parents, nil)
	fmt.Println("Parents: ", parents)

	startingNode := getStartingNode(7, parents)
	fmt.Println("startingNode: ", startingNode)

	type tuple struct {
		node     *BinaryTree
		distance int
	}

	queue := []tuple{}
	queue = append(queue, tuple{startingNode, 0})
	visited := map[int]bool{}
	var current tuple
	for len(queue) > 0 {
		// fmt.Println("queue at this point: ", queue)
		current, queue = queue[0], queue[1:]

		visited[current.node.value] = true
		if current.distance == 2 {
			ans = append(ans, current.node.value)
		}
		// Check if left, right, and parent are present before adding them to the queue
		// fmt.Println("I think cur node left is undefined here: ", current.node.left)
		if current.node.left != nil {
			if _, present := visited[current.node.left.value]; !present {
				queue = append(queue, tuple{current.node.left, current.distance + 1})
			}
		}
		if current.node.right != nil {
			if _, present := visited[current.node.right.value]; !present {
				queue = append(queue, tuple{current.node.right, current.distance + 1})
			}
		}
		// has the parent been visited yet?
		// the parent is the current nodes value in the parents map
		currentVal := current.node.value
		parentOfCurrentVal := parents[currentVal]
		// fmt.Println("parentOfCurrentVal: ", parentOfCurrentVal)
		if parentOfCurrentVal != nil {
      fmt.Println("Parent is not nil")
			_, parentVisitedYet := visited[parentOfCurrentVal.value]
      fmt.Println("Parent has not een visited yet")
			if parentVisitedYet == false {
        fmt.Println("this should happen next")
				queue = append(queue, tuple{parentOfCurrentVal, current.distance + 1})
			}
		}
	}

	fmt.Println("Answer after all of this: ", ans)

}

func getStartingNode(targetValue int, parentsMap map[int]*BinaryTree) *BinaryTree {
	targetParent := parentsMap[targetValue]

	if targetParent.left != nil && targetParent.left.value == targetValue {
		return targetParent.left
	}
	return targetParent.right
}
