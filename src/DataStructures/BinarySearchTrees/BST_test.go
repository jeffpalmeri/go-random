package binarysearchtrees

import (
	"reflect"
	"testing"
)

// Func to make a tree and return the tree and an array of values after traversal
// Use the tree to run my traversal on
// Compare the answer from my traversal to the correct array
func TestTraversals(t *testing.T) {
	compareValues := makeTrees()

	for _, val := range compareValues {
		// Each val in the loop is a struct that has a tree and the values from a traversal.
		// So I can run my new traversal on that tree and then compare the values

		tree := val.tree
		correctAnswerIn := val.inOrder
		correctAnswerPre := val.preOrder
		correctAnswerPost := val.PostOrder

		myAnswerIn := InOrderTraversal(tree, []int{})
		myAnswerPre := PreOrderTraversal(tree, []int{})
		myAnswerPost := PostOrderTraversal(tree, []int{})

		// Test for inorder comparison
		if !reflect.DeepEqual(correctAnswerIn, myAnswerIn) {
			t.Errorf("Inorder slices are not equal")
		} else {
			t.Logf("Inorder slices are equal")
		
		}
		// Test for preorder comparison
		if !reflect.DeepEqual(correctAnswerPre, myAnswerPre) {
			t.Errorf("Preorder slices are not equal")
		} else {
			t.Logf("Preorder slices are equal")
		}

		// Test for postorder comparison
		if !reflect.DeepEqual(correctAnswerPost, myAnswerPost) {
			t.Errorf("Postorder slices are not equal")
		} else {
			t.Logf("Postorder slices are equal")
		}
	}
}