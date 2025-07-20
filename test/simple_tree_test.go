package src_test

import (
	"dagon/src/structures"
	"fmt"
	"math/rand"
)

func SimpleTreeTest() {
	// 1
	tree := structures.CreateTree[int64]()
	// 2
	var max int
	fmt.Print("Enter the number of nodes to create: ")
	fmt.Scanf("%d", &max)
	// 3
	for i := 0; i < max; i++ {
		tree.Add(rand.Int63())
	}
}
