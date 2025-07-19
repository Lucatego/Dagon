package structures

import (
	"dagon/src/syntax/valid_types"
	"fmt"
)

type simple_tree[T valid_types.Types] struct {
	head *tree_node[T]
	size int64
}

func CreateTree[T valid_types.Types]() *simple_tree[T] {
	return &simple_tree[T]{
		head: nil,
		size: 0,
	}
}

func (t *simple_tree[T]) Add(data T) {
	// First case
	if t.head == nil {
		t.head = CreateNode(data)
		return
	}
	// Other cases
	err := t.head.InsertData(data)
	if err != nil {
		panic(fmt.Sprintf("Error: Tree unable to insert new data [Unexpected Error: %s.]", err.Error()))
	}
}
