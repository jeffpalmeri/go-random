package binarysearchtrees

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// func (tree *BST) print() {

// }

func print(w io.Writer, node *BST, ns int, ch rune) {
	if node == nil {
			return
	}

	for i := 0; i < ns; i++ {
			fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.value)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}
func TestInsert(t *testing.T) {
	root := BST{5, nil, nil}
	// root.InsertRecursive(&BST{4, nil, nil})
	root.InsertRecursive(&BST{6, nil, nil})
	fmt.Println(root)
}

func TestInsertRecursive(t *testing.T) {
	root := BST{5, nil, nil}
	root.InsertRecursive(&BST{4, nil, nil})
	root.InsertRecursive(&BST{6, nil, nil})
	root.InsertRecursive(&BST{3, nil, nil})
	root.InsertRecursive(&BST{7, nil, nil})
	root.InsertRecursive(&BST{2, nil, nil})
	root.InsertRecursive(&BST{8, nil, nil})
	fmt.Println(root)
	print(os.Stdout, &root, 0, 'M')
}
func TestPostOrderTraversal(t *testing.T) {

}

func TestRandtest(t *testing.T) {
	tree := BST{100, nil, nil}
	// tree.InsertRecursive(&tree)
	tree.InsertRecursive(&BST{-20, nil, nil})
	tree.InsertRecursive(&BST{-50, nil, nil})
	tree.InsertRecursive(&BST{-15, nil, nil})
	tree.InsertRecursive(&BST{-60, nil, nil})
	tree.InsertRecursive(&BST{50, nil, nil})
	tree.InsertRecursive(&BST{60, nil, nil})
	tree.InsertRecursive(&BST{55, nil, nil})
	tree.InsertRecursive(&BST{85, nil, nil})
	tree.InsertRecursive(&BST{15, nil, nil})
	tree.InsertRecursive(&BST{5, nil, nil})
	tree.InsertRecursive(&BST{-10, nil, nil})
	print(os.Stdout, &tree, 0, 'M')	
}

func TestTraversal(t *testing.T) {
	tree := BST{100, nil, nil}
	tree.InsertRecursive(&BST{-20, nil, nil})
	tree.InsertRecursive(&BST{-50, nil, nil})
	tree.InsertRecursive(&BST{-15, nil, nil})
	tree.InsertRecursive(&BST{-60, nil, nil})
	tree.InsertRecursive(&BST{50, nil, nil})
	tree.InsertRecursive(&BST{60, nil, nil})
	tree.InsertRecursive(&BST{55, nil, nil})
	tree.InsertRecursive(&BST{85, nil, nil})
	tree.InsertRecursive(&BST{15, nil, nil})
	tree.InsertRecursive(&BST{5, nil, nil})
	tree.InsertRecursive(&BST{-10, nil, nil})
	// print(os.Stdout, &tree, 0, 'M')

	values := make([]int, 0)
	ans := PostOrderTraversal(&tree, values)
	ans2 := PreOrderTraversal(&tree, values)
	ans3 := InOrderTraversal(&tree, values)

	fmt.Println("PostOrderTraversal: ", ans)
	fmt.Println("PreOrderTraversal: ", ans2)
	fmt.Println("InOrderTraversal", ans3)
}

func TestPop3(t *testing.T) {
	l := MyIntSlice{1,2,3,4,5}
	last, l := l.Pop3()

	fmt.Println("Result: ", last, l)
}

func TestOne(t *testing.T) {
	tree := BST{100, nil, nil}
	tree.InsertRecursive(&BST{-20, nil, nil})
	tree.InsertRecursive(&BST{-50, nil, nil})
	tree.InsertRecursive(&BST{-15, nil, nil})
	tree.InsertRecursive(&BST{-60, nil, nil})
	tree.InsertRecursive(&BST{50, nil, nil})
	tree.InsertRecursive(&BST{60, nil, nil})
	tree.InsertRecursive(&BST{55, nil, nil})
	tree.InsertRecursive(&BST{85, nil, nil})
	tree.InsertRecursive(&BST{15, nil, nil})
	tree.InsertRecursive(&BST{5, nil, nil})
	tree.InsertRecursive(&BST{-10, nil, nil})
	// values := make([]int, 0)

	// f := TryingIterativeTraversal(&tree, values)
	f := InOrderTraversalIterative(&tree)
	fmt.Println("f: ", f)
}

func TestPrepend(t *testing.T) {
	data := []int{1, 2, 3, 4}
	d := prependInt2(data, 99)
	fmt.Println(d)
}

// *********************
// traversals2.go
func Test2(t *testing.T) {
	tree := makeTree()
	values := make([]int, 0)
	values = tree.InOrderRecursive(values)
	fmt.Println("Values: ", values)

	// values2 := make([]int, 0)
	// values2 = InOrderTraversal(tree, values2)
	// fmt.Println("Values2: ", values2)
}
