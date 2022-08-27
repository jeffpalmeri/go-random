package binarysearchtrees

import "fmt"

type BST struct {
	value int
	left *BST
	right *BST
}
// func PostOrderTraversal(root *BST, values []int) []int{
// 	fmt.Println("PostOrderTraversal")
// 	if root == nil {
// 		return values
// 	}
// 	fmt.Println("Appending next line")
// 	PostOrderTraversal(root.left, values)
// 	PostOrderTraversal(root.right, values)
// 	values = append(values, root.value)
// 	return values
// }


func TryingIterativeTraversal(root *BST, values []int) []int {
  // commenttttt
	stack := make([]*BST, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		fmt.Println("Lnegth: ", len(stack))
		curr := Pop2(stack)
		fmt.Println("Curr: ", curr)
		if curr.left != nil {
			stack = append(stack, curr.left)
			curr = curr.left
			// continue
		}
		values = append(values, curr.value)
		if curr.right != nil {
			stack = append(stack, curr.right)
			curr = curr.right
			// continue
			fmt.Println(curr)
		}
	}

	return values
}

func InorderIterative(root *BST) {
	if root == nil {
		return
	}
 
	temp := root
	stack := make([]*BST, 0)
	for len(stack) > 0 || temp != nil {
		if temp != nil {
			stack = append(stack, temp)
			temp = temp.left
		} else {
			obj := Pop2(stack)
			// temp = obj.(*Node)
			temp = obj
			fmt.Printf("%d ", temp.value)
			temp = temp.right
		}
	}
}

func prependInt2(data []int, val int) []int {
  // data := []int{1, 2, 3, 4}
  fmt.Println(data)
  data = append([]int{val}, data...)
  fmt.Println(data)
	return data
}
func Prepend(data []*BST, val *BST) []*BST {
  // data := []int{1, 2, 3, 4}
  fmt.Println(data)
  data = append([]*BST{val}, data...)
  fmt.Println(data)
	return data
}

func Pop2(list []*BST) *BST {
// func Pop2(list []int) int {
	length := len(list)
	first := list[0]
	list = list[1:length]

	fmt.Println("Last: ", list)
	return first
}

type MyIntSlice []int

func (list MyIntSlice) Pop3() (int, MyIntSlice) {
	fmt.Println("l address: ", list)
	// list := *l
	fmt.Println("List before: ", list)
	length := len(list)
	last := list[length - 1]
	list = list[:length - 1]

	fmt.Println("List after: ", list)
	// fmt.Println("Last: ", last)
	return last, list
}

func (root *BST) InsertNode(node *BST) *BST{
	curr := root
	for curr != nil {
		if node.value < root.value {
			if root.left == nil {
				root.left = node
			}	else {
				root = root.left
			}
		}
		if node.value >= root.value {
			if root.right == nil {
				root.right = node
			} else {
				root = root.right
			}
		}
	}
	return root
}

func (root *BST) InsertRecursive(node *BST) {
	// if node == nil {
	// 	return root
	// }
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
	// return root
}


// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// func BranchSums(root *BinaryTree) []int {
//     sums := []int{}
//     currentSum := 0;
//     return recurse(root, sums, currentSum)
// }

// func recurse(node *BinaryTree, sums []int, currentSum int) []int {
//    fmt.Println("Recursing") 
    
//     if node.Left != nil {
//         sums = recurse(node.Left, sums, currentSum + node.Value)
//     } else {
//         sums = append(sums, currentSum + node.Value)
//         return sums
//     }
//     if node.Right != nil {
//         sums = recurse(node.Right, sums, currentSum + node.Value)
//     } else {
//         sums = append(sums, currentSum + node.Value)
//         return sums
//     }
//     return sums
// } 
func BranchSums(root *BinaryTree) []int {
	sums := []int{}
	currentSum := 0;
	return recurse(root, sums, currentSum)
}

func recurse(node *BinaryTree, sums []int, currentSum int) []int {
 fmt.Println("Recursing") 
 if node == nil {
		 return sums
 } 
	
	if node.Left == nil && node.Right == nil {
			sums = append(sums, currentSum + node.Value)
	} else {
			sums = recurse(node.Left, sums, currentSum + node.Value)
			sums = recurse(node.Right, sums, currentSum + node.Value)
	}
	return sums
} 
