package heaps

import (
	"fmt"
	"testing"
)

// Do not edit the class below except for the buildHeap,
// siftDown, siftUp, peek, remove, and insert methods.
// Feel free to add new properties and methods to the class.
type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	// Do not edit the lines below.
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
	//     for i, _ := range array {
	//         h.siftDown(i, h.length() - 1)
	//         fmt.Println("*******: ", *h, len(array))
	//     }
	for i := len(array) - 1; i >= 0; i-- {
		h.siftUp()
	}
}

func (h *MinHeap) siftDown(currentIdx, endIdx int) {
	// fmt.Println("currentIdx at begining: ", currentIdx, endIdx)
	childOneIdx := currentIdx*2 + 1
	fmt.Println("child one idx: ", childOneIdx)
	// childTwoIdx := currentIdx*2 + 2
	// declare childTwoIdx but don't initialize
	// if the index of what would be child 2 in valid, then gets its value
	// If not, ignore it and only use child one
	for childOneIdx <= endIdx {
		//		fmt.Println("in while: ", childOneIdx, childTwoIdx, currentIdx, (*h)[currentIdx], *h)
		childTwoIdx := -1
		if currentIdx*2+2 <= endIdx {
			childTwoIdx = currentIdx*2 + 2
		}
		smallerChildIdx := min(childOneIdx, childTwoIdx, h)
		fmt.Println("Smaller index:", smallerChildIdx)
		//		fmt.Println("smaller child idx:", smallerChildIdx)
		currentVal := (*h)[currentIdx]
		childVal := (*h)[smallerChildIdx]
		//		fmt.Println("Currnet val, childVal ", currentVal, childVal)
		if currentVal > childVal {
			h.swap(currentIdx, smallerChildIdx)
			//			fmt.Println("heap after swap: ", *h)
			currentIdx = smallerChildIdx
			childOneIdx = currentIdx*2 + 1
			childTwoIdx = currentIdx*2 + 2
		} else {
			return
		}
	}
}

func (h *MinHeap) siftUp() {
	lastIdx := h.length() - 1
	for lastIdx > 0 {
		parentIdx := (lastIdx - 1) / 2
		if (*h)[lastIdx] < (*h)[parentIdx] {
			h.swap(lastIdx, parentIdx)
			lastIdx = parentIdx
		} else {
			return
		}
	}
}

func (h MinHeap) Peek() int {
	return h[0]
}

func (h *MinHeap) Remove() int {
	returnNode := h.Peek()
	h.swap(0, h.length()-1)
	// I was forgetting to remove the last node after swapping...
	(*h) = (*h)[:h.length()-2]
	h.siftDown(0, h.length()-1)
	return returnNode
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}

func (h MinHeap) swap(idxA, idxB int) {
	h[idxA], h[idxB] = h[idxB], h[idxA]
}

func (h MinHeap) length() int {
	return len(h)
}

func min(childOneIdx, childTwoIdx int, h *MinHeap) int {
	if childTwoIdx == -1 {
		return childOneIdx
	}
	left, right := (*h)[childOneIdx], (*h)[childTwoIdx]
	if left <= right {
		return childOneIdx
	}
	return childTwoIdx
}

// go test -v -run TestHeapConstruction ./... -count=1
func TestHeapConstruction(t *testing.T) {
	arr := []int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}
	heap := MinHeap([]int{})
	fmt.Println(heap)
	heap.Insert(7)
	fmt.Println(heap)
	heap.Insert(8)
	fmt.Println(heap)
	heap.Insert(5)
	fmt.Println(heap)
	heap.Insert(13)
	fmt.Println(heap)
	heap.Insert(2)
	fmt.Println(heap)
	fmt.Println("************************")

	h := MinHeap{}
	for _, val := range arr {
		h.Insert(val)
	}
	fmt.Println(h)
	fmt.Println("************************")
	h.Insert(76)
	peak := h.Peek()
	fmt.Println("Peek: ", peak)
	fmt.Println(h)
	h.Remove()
	fmt.Println(h)
}

// go test -v -run TestInsert ./... -count=1
func TestInsert(t *testing.T) {
	arr := []int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}
	h := MinHeap{}
	for _, val := range arr {
		h.Insert(val)
		fmt.Println(h, " inserted:", val, h.IsValidHeap())
	}
	h.Insert(76)
	fmt.Println(h)
}

func (h *MinHeap) IsValidHeap() bool {
	//	q := []int{h.Peek()}
	// Check each child, and see if parent they are greater or not
	parentIdx := 0

	for parentIdx <= h.length()/2-1 {
		child1Idx := parentIdx*2 + 1
		child2Idx := parentIdx*2 + 2
		// Check if child one is less than parent
		if (*h)[parentIdx] <= (*h)[child1Idx] {
			// Check if the index of child two is valid
			if child2Idx < h.length() {
				// If child two does exist, then check if it's less than parent
				if !((*h)[parentIdx] <= (*h)[child2Idx]) {
					return false
				}
			}
			parentIdx++
			continue
		} else {
			return false
		}
	}
	return true
}

// go test -v -run TestIsHeapValid ./... -count=1
func TestIsHeapValid(t *testing.T) {
	h1 := MinHeap([]int{-5, 2, 6, 8, 24, 7, 24})
	h2 := MinHeap([]int{-5, 24, 6, 8, 2, 7, 24})
	h3 := MinHeap([]int{})
	fmt.Println(h1.IsValidHeap())
	fmt.Println(h2.IsValidHeap())
	fmt.Println(h3.IsValidHeap())
}

// go test -v -run TestBrokenInsert ./... -count=1
func TestBrokenInsert(t *testing.T) {
	h := MinHeap([]int{-5, 7, 8, 24, 24, 12, 24, 391, 48, 56})
	fmt.Println(h, h.IsValidHeap())
	h.Insert(2)
	fmt.Println(h, h.IsValidHeap())
}

// go test -v -run TestArrayToHeap ./... -count=1
func TestArrayToHeap(t *testing.T) {
	arr := []int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}
	heap := MinHeap(arr)
	fmt.Println(heap, heap.IsValidHeap())
	/*
		first := (len(arr) - 2) / 2
		for curr := first + 1; curr >= 0; curr-- {
			heap.siftDown(curr, len(arr)-1)
		}
	*/
	heap.buildHeap(arr)
	fmt.Println(heap, heap.IsValidHeap())
}

func (h *MinHeap) buildHeap(arr []int) {

	first := (len(arr) - 2) / 2
	for curr := first + 1; curr >= 0; curr-- {
		h.siftDown(curr, len(arr)-1)
	}
}

// go test -v -run TestAlgoExpert ./... -count=1
func TestAlgoExpert(t *testing.T) {
	arr := []int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}
	h := MinHeap(arr)
	(&h).buildHeap(arr)
	fmt.Println("bulildHeap: ", h.IsValidHeap(), h)
	p := h.Peek()
	fmt.Println("peek: ", h.IsValidHeap(), p)
}

// go test -v -run TestSwapDown ./... -count=1
func TestSwapDown(t *testing.T) {
	arr := []int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}
	h := MinHeap(arr)
	(&h).buildHeap(arr)
	fmt.Println("bulildHeap: ", h.IsValidHeap(), h)
	r := h.Remove()
	fmt.Println("remove: ", r, h.IsValidHeap(), h)
}

// go test -v -run TestPeek ./... -count=1
func TestPeek(t *testing.T) {
	arr := []int{544, -578, 556, 713, -655, -359, -810, -731, 194, -531, -685, 689, -279, -738, 886, -54, -320, -500, 738, 445, -401, 993, -753, 329, -396, -924, -975, 376, 748, -356, 972, 459, 399, 669, -488, 568, -702, 551, 763, -90, -249, -45, 452, -917, 394, 195, -877, 153, 153, 788, 844, 867, 266, -739, 904, -154, -947, 464, 343, -312, 150, -656, 528, 61, 94, -581}
	fmt.Println("********")
	h := MinHeap(arr)
	(&h).buildHeap(arr)
	fmt.Println("bulildHeap: ", h.IsValidHeap(), h)
	fmt.Println(len(arr))
}

// O(log(n)) time | O(1) space
func (h *MinHeap) AlgoSiftDown(currentIndex, endIndex int) {
	childOneIdx := currentIndex*2 + 1
	for childOneIdx <= endIndex {
		childTwoIdx := -1
		if currentIndex*2+2 <= endIndex {
			childTwoIdx = currentIndex*2 + 2
		}
		indexToSwap := childOneIdx
		if childTwoIdx > -1 && (*h)[childTwoIdx] < (*h)[childOneIdx] {
			indexToSwap = childTwoIdx
		}
		if (*h)[indexToSwap] < (*h)[currentIndex] {
			h.swap(currentIndex, indexToSwap)
			currentIndex = indexToSwap
			childOneIdx = currentIndex*2 + 1
		} else {
			return
		}
	}
}
