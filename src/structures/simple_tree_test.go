package structures

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSimpleTree(t *testing.T) {
	// 1
	tree := CreateTree[int64]()
	// 2
	var max int
	fmt.Print("Enter the number of nodes to create: ")
	fmt.Scanf("%d", &max)
	// 3
	for i := 0; i < max; i++ {
		tree.Add(rand.Int63())
	}
}
