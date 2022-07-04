package binarysearchtrees

func PostOrderTraversal(root *BST, values []int) []int{
	if root.left != nil {
		values = PostOrderTraversal(root.left, values)
	}
	if root.right != nil {
		values = PostOrderTraversal(root.right, values)
	}
	values = append(values, root.value)
	return values
}

func PreOrderTraversal(root *BST, values []int) []int {
	values = append(values, root.value)
	if root.left != nil {
		values = PreOrderTraversal(root.left, values)
	}
	if root.right != nil {
		values = PreOrderTraversal(root.right, values)
	}
	return values
}

func InOrderTraversal(root *BST, values []int) []int {
	if root.left != nil {
		values = InOrderTraversal(root.left, values)
	}
	values = append(values, root.value)
	if root.right != nil {
		values = InOrderTraversal(root.right, values)
	}
	return values
}

/////////////////////////

/*
Traverse down left and add them to the stack as long as there is a left.
Once there's not, pop from the stack, THEN append to vales, and change curr
to curr.right and continue.
*/
type Stack []*BST
func (list Stack) Pop4() (*BST, Stack) {
	length := len(list)
	last := list[length - 1]
	list = list[:length - 1]

	return last, list
}

func InOrderTraversalIterative(root *BST) []int {
	stack := make(Stack, 0)
	curr := root
	values := make([]int, 0)
	for len(stack) > 0 || curr != nil {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		} else {
			curr, stack = stack.Pop4()
			values = append(values, curr.value)
			curr = curr.right
		}
	}
	return values
}