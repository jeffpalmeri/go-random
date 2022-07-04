package binarysearchtrees

import "fmt"

type TestForTrav struct {
	tree *BST
	inOrder []int
	preOrder []int
	PostOrder []int
}

func makeTrees() []TestForTrav {
	t1 := BST{100, nil, nil}
	t1.InsertRecursive(&BST{-20, nil, nil})
	t1.InsertRecursive(&BST{-50, nil, nil})
	t1.InsertRecursive(&BST{-15, nil, nil})
	t1.InsertRecursive(&BST{-60, nil, nil})
	t1.InsertRecursive(&BST{50, nil, nil})
	t1.InsertRecursive(&BST{60, nil, nil})
	t1.InsertRecursive(&BST{55, nil, nil})
	t1.InsertRecursive(&BST{85, nil, nil})
	t1.InsertRecursive(&BST{15, nil, nil})
	t1.InsertRecursive(&BST{5, nil, nil})
	t1.InsertRecursive(&BST{-10, nil, nil})
	v1 := []int{}
	v1In := InOrderTraversal(&t1, v1)
	v1Pre := PreOrderTraversal(&t1, v1)
	v1Post := PostOrderTraversal(&t1, v1)

	t2 := BST{100, nil, nil}
	t2.InsertRecursive(&BST{502, nil, nil})
	t2.InsertRecursive(&BST{55000, nil, nil})
	t2.InsertRecursive(&BST{1001, nil, nil})
	t2.InsertRecursive(&BST{4500, nil, nil})
	t2.InsertRecursive(&BST{204, nil, nil})
	t2.InsertRecursive(&BST{205, nil, nil})
	t2.InsertRecursive(&BST{207, nil, nil})
	t2.InsertRecursive(&BST{208, nil, nil})
	t2.InsertRecursive(&BST{206, nil, nil})
	t2.InsertRecursive(&BST{203, nil, nil})
	t2.InsertRecursive(&BST{5, nil, nil})
	t2.InsertRecursive(&BST{15, nil, nil})
	t2.InsertRecursive(&BST{22, nil, nil})
	t2.InsertRecursive(&BST{57, nil, nil})
	t2.InsertRecursive(&BST{60, nil, nil})
	t2.InsertRecursive(&BST{5, nil, nil})
	t2.InsertRecursive(&BST{2, nil, nil})
	t2.InsertRecursive(&BST{3, nil, nil})
	t2.InsertRecursive(&BST{1, nil, nil})
	t2.InsertRecursive(&BST{1, nil, nil})
	t2.InsertRecursive(&BST{1, nil, nil})
	t2.InsertRecursive(&BST{1, nil, nil})
	t2.InsertRecursive(&BST{1, nil, nil})
	t2.InsertRecursive(&BST{-51, nil, nil})
	t2.InsertRecursive(&BST{-403, nil, nil})

	// "inOrderArray": [-403, -51, 1, 1, 1, 1, 1, 2, 3, 5, 5, 15, 22, 57, 60, 100, 203, 204, 205, 206, 207, 208, 502, 1001, 4500, 55000],
  // "postOrderArray": [-403, -51, 1, 1, 1, 1, 1, 3, 2, 5, 60, 57, 22, 15, 5, 203, 206, 208, 207, 205, 204, 4500, 1001, 55000, 502, 100],
  // "preOrderArray": [100, 5, 2, 1, -51, -403, 1, 1, 1, 1, 3, 15, 5, 22, 57, 60, 502, 204, 203, 205, 207, 206, 208, 55000, 1001, 4500]

	v2 := []int{}
	v2In := InOrderTraversal(&t2, v2)
	v2Pre := PreOrderTraversal(&t2, v2)
	v2Post := PostOrderTraversal(&t2, v2)
	fmt.Println("v2: ", v2)

	/*
      {"id": "10", "left": "5", "right": "15", "value": 10},
      {"id": "15", "left": null, "right": "22", "value": 15},
      {"id": "22", "left": null, "right": null, "value": 22},
      {"id": "5", "left": "2", "right": "5-2", "value": 5},
      {"id": "5-2", "left": null, "right": null, "value": 5},
      {"id": "2", "left": "1", "right": null, "value": 2},
      {"id": "1", "left": "-5", "right": null, "value": 1},
      {"id": "-5", "left": "-15", "right": "-5-2", "value": -5},
      {"id": "-5-2", "left": null, "right": "-2", "value": -5},
      {"id": "-2", "left": null, "right": "-1", "value": -2},
      {"id": "-1", "left": null, "right": null, "value": -1},
      {"id": "-15", "left": "-22", "right": null, "value": -15},
      {"id": "-22", "left": null, "right": null, "value": -22}
	*/

	// "inOrderArray": [-22, -15, -5, -5, -2, -1, 1, 2, 5, 5, 10, 15, 22],
  // "postOrderArray": [-22, -15, -1, -2, -5, -5, 1, 2, 5, 5, 22, 15, 10],
  // "preOrderArray": [10, 5, 2, 1, -5, -15, -22, -5, -2, -1, 5, 15, 22]

	t3 := BST{10, nil, nil}
	t3.InsertRecursive(&BST{15, nil, nil})
	t3.InsertRecursive(&BST{22, nil, nil})
	t3.InsertRecursive(&BST{5, nil, nil})
	t3.InsertRecursive(&BST{5, nil, nil})
	t3.InsertRecursive(&BST{2, nil, nil})
	t3.InsertRecursive(&BST{1, nil, nil})
	t3.InsertRecursive(&BST{-5, nil, nil})
	t3.InsertRecursive(&BST{-5, nil, nil})
	t3.InsertRecursive(&BST{-2, nil, nil})
	t3.InsertRecursive(&BST{-1, nil, nil})
	t3.InsertRecursive(&BST{-15, nil, nil})
	t3.InsertRecursive(&BST{-22, nil, nil})
	v3:= []int{}
	v3In := InOrderTraversal(&t3, v3)
	v3Pre := PreOrderTraversal(&t3, v3)
	v3Post := PostOrderTraversal(&t3, v3)
	fmt.Println("v3: ", v3)

	returnValue := []TestForTrav{
		{&t1, v1In, v1Pre, v1Post},
		{&t2, v2In, v2Pre, v2Post},
		{&t3, v3In, v3Pre, v3Post},
	}

	return returnValue
	
}

func makeTree() *BST {
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
	return &tree
}

func Pop(list []BST) (*BST, []BST) {
	last := list[len(list) - 1]
	list = list[:len(list) - 1]

	return &last, list
}

// values := make([]int, 0)
func (tree *BST) InOrderRecursive(values []int) []int {
	curr := tree

	if curr.left != nil {
		values = curr.left.InOrderRecursive(values)
	}
	values = append(values, curr.value)
	if curr.right != nil {
		values = curr.right.InOrderRecursive(values)
	}
	return values
}

// package main

// import "math"

// type BST struct {
// 	Value int

// 	Left  *BST
// 	Right *BST
// }

// func (tree *BST) FindClosestValue(target int) int {
//     closest := math.MaxInt32
//     curr := tree
//     for curr != nil {
//         if curr.Value == target {
//             return target
//         }
//         if target < curr.Value {
//             if curr.Value - target < closest {
//                 closest = curr.Value
//             }
//             curr = curr.Left
//             continue
//         }
//         if target > curr.Value {
//             if target - curr.Value < closest {
//                 closest = curr.Value 
//             }
//             curr = curr.Right
//             continue
//         }
//     }
    
// 	// if closest == math.MaxInt32 {
//  //        return -1
//  //    }
//     return closest
// }