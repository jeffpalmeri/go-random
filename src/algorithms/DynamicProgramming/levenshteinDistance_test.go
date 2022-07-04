package dynamicprogramming

import (
	"fmt"
	"testing"
)

// go test -v ./src/algorithms/DynamicProgramming
// go test -v -run TestLvnDistance ./...
func TestLvDistance(t *testing.T) {
	fmt.Println(LevenshteinDistance("word", "one"))
}